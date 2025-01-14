package cmd

import (
	"fmt"
	"github.com/yasukotelin/kdcf/formatter"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "kdcf",
	Long: `kdcf is a CLI tool for formatting Kotlin data classes.

e.g.
$ kdcf 'User(name="user name", age=20)'
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Println()
			s := formatter.FormatKotlinData(args[0])
			fmt.Println(s)
		}
	},
	Version: "v0.1.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
