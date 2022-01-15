package run

import (
	"fmt"
	"os"
	"strings"

	"github.com/kubetrail/qq/pkg/crypto"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

func E(cmd *cobra.Command, args []string) error {
	key, err := base58.Decode(os.Getenv("AES_KEY"))
	if err != nil {
		return fmt.Errorf("could not decode AES key: %w", err)
	}

	out, err := crypto.EncryptWithAesKey([]byte(strings.Join(args, " ")), key)
	if err != nil {
		return fmt.Errorf("could not encrypt message: %w", err)
	}

	if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%s\n", base58.Encode(out)); err != nil {
		return fmt.Errorf("could not write to output: %w", err)
	}

	return nil
}
