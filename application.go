package main

import (
	"fmt"
	"github.com/just1689/remote-pi/controller/daemon"
)

func main() {
	fmt.Println("Application starting")

	projectID := ""
	topicName := ""
	credentialsFile := ""

	daemon.StartDaemon(projectID, topicName, credentialsFile)
}
