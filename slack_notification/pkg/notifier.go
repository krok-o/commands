package pkg

import (
	"fmt"

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
	fmt.Println("token: ", n.Token)
	fmt.Println("notifying: ", parser.RepoName())
	return nil
}
