package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	csdk "github.com/krok-o/command-sdk"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Input Command",
		Run:   runRootCmd,
	}
	rootArgs struct {
		options     csdk.Options
		extraOption string
	}
)

func init() {
	csdk.AddRequiredFlagsToCommand(rootCmd, &rootArgs.options)

	rootCmd.Flags().StringVar(&rootArgs.extraOption, "extra-flags", "", "Extra flag provided by the outputter command.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Options: ", rootArgs)
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
