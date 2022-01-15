package run

import (
	"fmt"
	"os"

	"github.com/kubetrail/qq/pkg/crypto"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

func D(cmd *cobra.Command, args []string) error {
	key, err := base58.Decode(os.Getenv("AES_KEY"))
	if err != nil {
		return fmt.Errorf("could not decode AES key: %w", err)
	}

	msg, err := base58.Decode(args[0])
	if err != nil {
		return fmt.Errorf("could not decode message: %w", err)
	}

	out, err := crypto.DecryptWithAesKey(msg, key)
	if err != nil {
		return fmt.Errorf("could not decrypt message: %w", err)
	}

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", out); err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	return nil
}
