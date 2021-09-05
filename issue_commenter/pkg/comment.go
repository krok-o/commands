package pkg

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"

	kgithub "github.com/krok-o/command-sdk/github"
)

type Config struct {
	Platform  string
	EventType string
	Payload   string
	Token     string
	Message   string
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
		// TODO: This needs to be extracted into its own package.
		// Needs to deal with Issues, right now it only handles PullRequests.
		parser, err := kgithub.NewParser(c.Payload)
		if err != nil {
			return fmt.Errorf("Failed to create parser: %w", err)
		}
		fullName := parser.RepoName()
		if !strings.Contains(fullName, "/") {
			return fmt.Errorf("Failed to extract owner and repo from full name %s", fullName)
		}
		split := strings.Split(fullName, "/")
		owner := split[0]
		repo := split[1]

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: c.Token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client := github.NewClient(tc)
		commentID := parser.CommentID()
		issueID := parser.IssueID()
		comment, _, err := client.Issues.GetComment(context.Background(), owner, repo, commentID)
		if err != nil {
			return fmt.Errorf("failed to get comment: %w", err)
		}
		// TODO: analyse command / run a plugin / etc...
		c.Message += " in response to comment: " + comment.GetBody()

		if _, _, err := client.Issues.CreateComment(context.Background(), owner, repo, int(issueID), &github.IssueComment{Body: &c.Message}); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add comment: %s", err)
			return err
		}
	}
	return nil
}
