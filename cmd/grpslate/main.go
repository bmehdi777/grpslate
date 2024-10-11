package main

import (
	"fmt"
	"os"

	"github.com/bmehdi777/grpslate/internal/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
