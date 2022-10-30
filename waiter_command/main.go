package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	waitTime    time.Duration
	platform    string
	eventType   string
	payload     string
	artifactURL string
)

func init() {
	flag.DurationVar(&waitTime, "wait", 10*time.Second, "wait time")
	flag.StringVar(&platform, "platform", "", "Description")
	flag.StringVar(&eventType, "event-type", "", "Description")
	flag.StringVar(&payload, "payload", "", "Description")
	flag.StringVar(&artifactURL, "artifact-url", "", "Description")
	flag.Parse()
}

func main() {
	fmt.Printf("Waiting for configured amount of time: %d\n", waitTime)
	time.Sleep(waitTime)
	fmt.Println("Done with waiting. Have fun!")
}
