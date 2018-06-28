package main

import (
	"fmt"
	"github.com/just1689/remote-pi/controller"
	"github.com/just1689/remote-pi/controller/io"
)

const configFile = "app-config.json"

func main() {
	fmt.Println("Application starting v1.0")
	start()
}

func start() {

	config := io.ReadConfig(configFile)
	fmt.Println("..config was loaded from file")
	controller.StartDaemon(config)
}
