package commands

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli/command/migration"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		migration.NewMigrationCommand(),
	)
}
