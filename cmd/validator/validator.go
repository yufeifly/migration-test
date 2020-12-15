package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yufeifly/validator/cli/command"
	"github.com/yufeifly/validator/cli/command/commands"
	"os"
)

func newValidatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "validator COMMAND",
		Short:         "A self-sufficient tester and validator of my container migration project",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          noArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			command.ShowHelp(cmd, args)
			return nil
		},
	}
	cmd.SetOut(os.Stdout)
	commands.AddCommands(cmd)
	return cmd
}
func noArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf("validator: '%s' is not a validator command.\nSee 'validator --help'", args[0])
}

func main() {
	cmd := newValidatorCommand()
	if err := cmd.Execute(); err != nil {
		logrus.Printf("cmd Execute err: %v", err)
		os.Exit(1)
	}
}
