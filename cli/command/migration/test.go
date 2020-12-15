package migration

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
	"github.com/yufeifly/validator/test/multipleservices"
)

func newTestCommand() *cobra.Command {
	var opts multipleservices.TestOptions
	cmd := &cobra.Command{
		Use: "test",
		Short: "test migration",
		Args: cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := multipleservices.TestMultipleService(opts)
			return err
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.Platform, "platform", "p", "pc","platform of test, my own pc or server")
	flags.IntVarP(&opts.Number, "number", "n", 1,"number of services to migrate")
	return cmd
}
