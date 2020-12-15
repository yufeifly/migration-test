package migration

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
)

type verifyOptions struct {
	mode string
}

func newVerifyCommand() *cobra.Command {
	var opts verifyOptions
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "verify the result of migration",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVerify(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.mode, "mode", "m", "pc","verify result of migration on server or my own pc")
	return cmd
}

func runVerify(opts verifyOptions) error {
	fmt.Printf("hello, %v\n", opts.mode)
	return nil
}
