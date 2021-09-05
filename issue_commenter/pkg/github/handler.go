package github

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
	kgithub "github.com/krok-o/command-sdk/github"
	"golang.org/x/oauth2"
)

// Handler handles comments made on Github.
type Handler interface {
	Handle() error
}

// GithubHandler defines a Github based handler.
type GithubHandler struct {
	Payload string
	Message string
	Token   string
}

// NewGithubHandler creates a new handler for Events from Github.
func NewGithubHandler(payload string, message string, token string) *GithubHandler {
	return &GithubHandler{Payload: payload, Message: message}
}

func (h *GithubHandler) Handle() error {
	// TODO: This needs to be extracted into its own package.
	// Needs to deal with Issues, right now it only handles PullRequests.
	parser, err := kgithub.NewParser(h.Payload)
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
		&oauth2.Token{AccessToken: h.Token},
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
	h.Message += " in response to comment: " + comment.GetBody()

	if _, _, err := client.Issues.CreateComment(context.Background(), owner, repo, int(issueID), &github.IssueComment{Body: &h.Message}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add comment: %s", err)
		return err
	}
	return nil
}
