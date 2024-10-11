package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func newCmdRemind() *cobra.Command {
	remindCmd := cobra.Command{
		Use:  "reminder",
		Run:  handlerRemind,
		Args: cobra.NoArgs,
	}

	return &remindCmd
}

func handlerRemind(cmd *cobra.Command, args []string) {
	userUnderlineColor := color.New(color.FgCyan).Add(color.Underline).SprintFunc()
	userColor := color.New(color.FgCyan).SprintFunc()
	groupUnderlineColor := color.New(color.FgRed).Add(color.Underline).SprintFunc()
	groupColor := color.New(color.FgRed).SprintFunc()
	otherUnderlineColor := color.New(color.FgGreen).Add(color.Underline).SprintFunc()
	otherColor := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("%s%s%s\n", userUnderlineColor("rwx"), groupUnderlineColor("rwx"), otherUnderlineColor("rwx"))

	fmt.Printf("%s%s%s\n", userColor("\\ /"), groupColor("\\ /"), otherColor("\\ /"))

	fmt.Printf(" %s  %s  %s  \n", userColor("v"), groupColor("v"), otherColor("v"))
	fmt.Printf(" %s  %s  %s  \n", userColor("|"), groupColor("|"), otherColor("Other's rights"))

	fmt.Printf(" %s  %s  \n", userColor("|"), groupColor("Group's rights"))
	fmt.Printf(" %s  \n", userColor("User's rights"))
}
