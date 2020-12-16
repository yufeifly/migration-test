package redis

import (
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli"
	"github.com/yufeifly/validator/redisbreed/breeder"
)

func newBreedCommand() *cobra.Command {
	var opts breeder.BreedOpts
	cmd := &cobra.Command{
		Use:   "breed",
		Short: "breed redis service",
		Args:  cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return breeder.BreedRedis(opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.RedisServer, "redis-server", "", "", "redis server address, format: 192.168.0.1:6666")
	flags.StringVarP(&opts.Range, "range", "", "", "key range to validate, [start: end) , format: key{1:10000}")
	return cmd
}
