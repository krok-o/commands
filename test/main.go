package main

import "fmt"

// Execute is the main entrypoint for this test plugin.
// Returns whatever the payload was for testing purposes.
func Execute(payload string, opts ...interface{}) (string, bool, error) {
	fmt.Println("Opts received: ", opts)
	return payload, true, nil
}
