package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/krok-o/command-sdk/github"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("wrong number of arguments received")
		os.Exit(1)
	}
	platform := args[1]
	fmt.Println("Received platform: ", platform)
	var in string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in += scanner.Text()
	}
}

func notifyViaSlack(payload string) error {
	repoName, err := github.ExtractRepoName(payload)
	if err != nil {
		return err
	}
	fmt.Println("notifying: ", repoName)
	return nil
}
