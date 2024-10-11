package cli

import "github.com/spf13/cobra"

func newCmdConvert() *cobra.Command {
	convertCmd := cobra.Command{
		Use: "convert <input>",
		Run: handlerConvert,
		Args: cobra.MinimumNArgs(1),
	}

	return &convertCmd
}

func handlerConvert(cmd *cobra.Command, args []string) {

}
