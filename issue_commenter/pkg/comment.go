package pkg

import (
	"fmt"

	cgithub "issue-commenter/pkg/github"
)

type Config struct {
	Platform  string
	EventType string
	Payload   string
	Token     string
	Message   string
	BotLogin  string
}

// Commenter takes a PR request and comments on it, if the previous
// comment had a trigger key word.
type Commenter struct {
	Config
}

// NewCommenter creates a new Commenter.
func NewCommenter(cfg Config) *Commenter {
	return &Commenter{Config: cfg}
}

// Comment sends a message to a PR.
func (c *Commenter) Comment() error {
	switch c.Platform {
	case "github":
		githubHandler := cgithub.NewGithubHandler(c.Payload, c.Message, c.Token, c.BotLogin)
		if err := githubHandler.Handle(); err != nil {
			return fmt.Errorf("failed to handle event: %w", err)
		}
	}
	return nil
}
