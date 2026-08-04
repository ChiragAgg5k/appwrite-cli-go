package main

import (
	"fmt"
	"os"

	"github.com/ChiragAgg5k/appwrite-cli-go/internal/app"
	"github.com/ChiragAgg5k/appwrite-cli-go/internal/cmd"
)

func main() {
	root := cmd.NewRootCommand()
	// pflag drops the value in `--enabled false`; join it on first. See
	// RewriteBooleanValues.
	root.SetArgs(cmd.RewriteBooleanValues(root, os.Args[1:]))

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, cmd.FormatError(err))
		if app.Flags().Report {
			fmt.Fprintln(os.Stderr, cmd.ReportBlock(err))
		}
		os.Exit(1)
	}
}
