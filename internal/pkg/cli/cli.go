package cli

import "github.com/spf13/cobra"

func newCmdRoot() *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "grpslate",
		Short: "Unix group translator",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	rootCmd.AddCommand(newCmdConvertHuman())
	rootCmd.AddCommand(newCmdRemind())

	return &rootCmd

}

func Execute() error {
	return newCmdRoot().Execute()
}
