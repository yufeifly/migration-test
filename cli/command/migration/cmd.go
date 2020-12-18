package migration

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
	"github.com/yufeifly/validator/cli/command"
)

func NewMigrationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migration",
		Short: "manage migration",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp,
	}
	cmd.AddCommand(
		newTestCommand(),
		newVerifyCommand(),
	)
	return cmd
}
