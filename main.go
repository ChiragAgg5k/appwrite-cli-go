package main

import (
	"fmt"
	"os"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
