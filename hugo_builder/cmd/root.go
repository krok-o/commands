package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Hugo blog builder",
		Run:   runRootCmd,
	}
	rootArgs struct {
		platform  string
		eventType string
		payload   string
	}
)

func init() {
	flag := rootCmd.Flags()
	// Server Configs
	flag.StringVar(&rootArgs.platform, "platform", "", "--platform github")
	flag.StringVar(&rootArgs.eventType, "event-type", "", "--event-type The type of the event occurred.")
	flag.StringVar(&rootArgs.payload, "payload", "", "--payload received from the platform.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
