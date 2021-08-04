package pkg

import (
	"fmt"

	"github.com/slack-go/slack"

	"github.com/krok-o/command-sdk/github"
)

// Notifier notifies.
type Notifier struct {
	Payload string
	Token   string
}

// Notify notifies.
func (n Notifier) Notify() error {
	parser, err := github.NewParser(n.Payload)
	if err != nil {
		return err
	}
	fmt.Println("notifying: ", parser.RepoName())

	api := slack.New(n.Token)
	channelID, ts, err := api.PostMessage(fmt.Sprintf("repository %q received event", parser.RepoName()), slack.MsgOptionText("test", false))
	if err != nil {
		return err
	}
	fmt.Printf("message posted to channel %s at %s\n", channelID, ts)
	return nil
}
