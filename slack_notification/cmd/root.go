package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slack-notification/pkg"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Slack notifier command",
		Run:   runRootCmd,
	}
	rootArgs struct {
		platform  string
		token     string
		channel   string
		message   string
		eventType string
	}
)

func init() {
	flag := rootCmd.Flags()
	// Server Configs
	flag.StringVar(&rootArgs.platform, "platform", "", "--platform github")
	flag.StringVar(&rootArgs.token, "token", "", "--token slack-token")
	flag.StringVar(&rootArgs.channel, "channel", "", "--channel ")
	flag.StringVar(&rootArgs.message, "message", "", "--message Event occurred.")
	flag.StringVar(&rootArgs.eventType, "event-type", "", "--event-type The type of the event occurred.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	var in string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in += scanner.Text()
	}
	n := pkg.Notifier{
		Payload:   in,
		Token:     rootArgs.token,
		Message:   rootArgs.message,
		Channel:   rootArgs.channel,
		EventType: rootArgs.eventType,
	}
	if err := n.Notify(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to send notification: %s", err)
		os.Exit(1)
	}
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
