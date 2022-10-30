package cmd

import (
	"github.com/spf13/cobra"

	csdk "github.com/krok-o/command-sdk"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Input Command",
		RunE:  runRootCmd,
	}
	rootArgs struct {
		options csdk.Options
	}
)

func init() {
	csdk.AddRequiredFlagsToCommand(rootCmd, &rootArgs.options)
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) error {
	return csdk.Output(map[string]string{
		"extra-flag": "this is the added value",
	})
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
