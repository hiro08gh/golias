package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "golias",
	Short: "golias is setup alias path in bash config.",
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(
		addCmd(),
	)
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		_ = rootCmd.Help()
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
