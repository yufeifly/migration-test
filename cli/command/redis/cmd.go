package redis

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
	"github.com/yufeifly/validator/cli/command"
)

func NewRedisCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redis",
		Short: "manage redis operations",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp,
	}
	cmd.AddCommand(
		newBreedCommand(),
	)
	return cmd
}
