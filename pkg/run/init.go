package run

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/monnand/dhkx"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/pbkdf2"
)

func Init(cmd *cobra.Command, _ []string) error {
	// Get a group. Use the default one would be enough.
	g, err := dhkx.GetGroup(0)
	if err != nil {
		return fmt.Errorf("could not get group: %w", err)
	}

	// Generate a private key from the group.
	// Use the default random number generator.
	priv, err := g.GeneratePrivateKey(nil)
	if err != nil {
		return fmt.Errorf("could not get private key: %w", err)
	}

	// Get the public key from the private key.
	pub := base58.Encode(priv.Bytes())
	var peerPub string
	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Send to peer: %s\nEnter peer key: ", pub); err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	if _, err := fmt.Fscanf(cmd.InOrStdin(), "%s", &peerPub); err != nil {
		return fmt.Errorf("could not read from input: %w", err)
	}

	b, err := base58.Decode(peerPub)
	if err != nil {
		return fmt.Errorf("could not decode peer key: %w", err)
	}

	k, err := g.ComputeKey(dhkx.NewPublicKey(b), priv)
	if err != nil {
		return fmt.Errorf("could not compute key from peer key: %w", err)
	}

	key := k.Bytes()
	salt := md5.Sum(key)

	key = pbkdf2.Key(key, salt[:8], 4096, 32, sha256.New)
	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "export AES_KEY=%s\n", base58.Encode(key)); err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	return nil
}
