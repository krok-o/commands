package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Integration test",
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
	flag.StringVar(&rootArgs.eventType, "payload", "", "--payload The payload.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("platform: %s,event-type: %s,payload: %s", rootArgs.platform, rootArgs.eventType, rootArgs.payload)
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
