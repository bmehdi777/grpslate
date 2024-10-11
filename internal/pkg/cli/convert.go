package cli

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func newCmdConvertHuman() *cobra.Command {
	convertCmd := cobra.Command{
		Use: "human <numbers>",
		Long: "Convert machine permissions to human language",
		RunE: handlerConvertHuman,
		Args: cobra.MinimumNArgs(1),
	}

	return &convertCmd
}

func newCmdConvertMachine() *cobra.Command {
	convertCmd := cobra.Command{
		Use: "machine <strings>",
		Long: "Convert human permissions to machine language",
		Run: handlerConvertMachine,
		Args: cobra.MinimumNArgs(1),
	}

	return &convertCmd
}

func handlerConvertHuman(cmd *cobra.Command, args []string) error {
	rights := strings.Split(args[0], "")
	if len(rights) != 3 {
		return errors.New("Rights should be 3 numbers.")
	}
	if _, err := strconv.Atoi(args[0]); err != nil {
		return errors.New("Rights should be 3 numbers.")
	}


	return nil
}

func handlerConvertMachine(cmd *cobra.Command, args []string) {

}
