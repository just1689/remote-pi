package main

import (
	"fmt"
	"github.com/just1689/remote-pi/controller/daemon"
)

func main() {
	fmt.Println("Application starting")

	projectID := "ex-remote-pi"
	topicName := "toggle"
	credentialsFile := "credentials.json"

	daemon.StartDaemonNoPi(projectID, topicName, credentialsFile)
}
