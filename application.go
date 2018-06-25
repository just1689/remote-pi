package main

import (
	"fmt"
	"github.com/just1689/remote-pi/controller/daemon"
	"github.com/just1689/remote-pi/controller/io"
)

const configFile = "config.json"

func main() {
	fmt.Println("Application starting")
	start()
}

func start() {

	config := io.ReadConfig(configFile)
	fmt.Println("..config was loaded from file")
	daemon.StartDaemon(config)
}
