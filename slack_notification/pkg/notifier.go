package pkg

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"

	"github.com/krok-o/command-sdk/github"
)

// Notifier notifies.
type Notifier struct {
	Payload   string
	Token     string
	Message   string
	Channel   string
	EventType string
}

// Notify notifies.
func (n Notifier) Notify() error {
	parser, err := github.NewParser(n.Payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse payload: %s", err)
		return err
	}
	fmt.Printf("notifying: %s for event type: %s\n", parser.RepoName(), n.EventType)

	api := slack.New(n.Token)
	channelID, ts, err := api.PostMessage(n.Channel, slack.MsgOptionText(fmt.Sprintf("repository %q received event type %q", parser.RepoName(), n.EventType), false))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to post api message: %s", err)
		return err
	}
	fmt.Printf("message posted to channel %s at %s\n", channelID, ts)
	return nil
}
