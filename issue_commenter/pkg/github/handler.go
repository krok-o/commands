package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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
	Payload  string
	Message  string
	Token    string
	BotLogin string
}

// NewGithubHandler creates a new handler for Events from Github.
func NewGithubHandler(payload string, message string, token string, botLogin string) *GithubHandler {
	return &GithubHandler{Payload: payload, Message: message, Token: token, BotLogin: botLogin}
}

func (h *GithubHandler) Handle() error {
	// TODO: This needs to be extracted into its own package.
	// Needs to deal with Issues, right now it only handles PullRequests.
	data, err := base64.StdEncoding.DecodeString(h.Payload)
	if err != nil {
		return fmt.Errorf("Failed to decode payload: %w", err)
	}
	parser, err := kgithub.NewParser(string(data))
	if err != nil {
		return fmt.Errorf("Failed to create parser: %w", err)
	}
	commenter := parser.CommenterLogin()
	if commenter == h.BotLogin {
		return nil
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
	issueNumber := parser.IssueNumber()
	comment, response, err := client.Issues.GetComment(context.Background(), owner, repo, commentID)
	if err != nil {
		return fmt.Errorf("failed to get comment: %w", err)
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read body: %s", err)
		}
		return fmt.Errorf("status code was not 2xx but: %d with error message: %s", response.StatusCode, string(content))
	}

	// TODO: analyse command / run a plugin / etc...
	h.Message += " in response to comment: " + comment.GetBody()

	if _, _, err := client.Issues.CreateComment(context.Background(), owner, repo, int(issueNumber), &github.IssueComment{Body: &h.Message}); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add comment: %s", err)
		return err
	}
	return nil
}
