package cmd

import (
	"fmt"
	"issue-commenter/pkg"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Issue commenter",
		Run:   runRootCmd,
	}
	rootArgs struct {
		commenter pkg.Config
	}
)

func init() {
	flag := rootCmd.Flags()
	// Server Configs
	flag.StringVar(&rootArgs.commenter.Platform, "platform", "", "--platform github")
	flag.StringVar(&rootArgs.commenter.EventType, "event-type", "", "--event-type The type of the event occurred")
	flag.StringVar(&rootArgs.commenter.Payload, "payload", "", "--payload received from the platform")
	flag.StringVar(&rootArgs.commenter.Token, "token", "", "--token contains the token to access the api of the repo")
	flag.StringVar(&rootArgs.commenter.Message, "message", "", "--message the message to send")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	c := pkg.NewCommenter(rootArgs.commenter)
	if err := c.Comment(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to leave comment: %s", err)
		os.Exit(1)
	}
	fmt.Println("All done. Good bye!")
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
