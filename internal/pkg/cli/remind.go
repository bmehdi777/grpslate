package cli

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func newCmdRemind() *cobra.Command {
	remindCmd := cobra.Command{
		Use: "remind",
		Run: handlerRemind,
		Args: cobra.NoArgs,
	}

	return &remindCmd
}

func handlerRemind(cmd *cobra.Command, args []string) {

}
