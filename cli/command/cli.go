package command

import (
	"github.com/spf13/cobra"
	"os"
)

// ShowHelp shows the command help.
func ShowHelp(cmd *cobra.Command, args []string) error {
	cmd.SetOut(os.Stderr)
	cmd.HelpFunc()(cmd, args)
	return nil
}
