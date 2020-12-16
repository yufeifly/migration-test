package commands

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli/command/migration"
	"github.com/yufeifly/validator/cli/command/redis"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		migration.NewMigrationCommand(),
		redis.NewRedisCommand(),
	)
}
