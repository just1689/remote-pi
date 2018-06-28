package main

import (
	"github.com/just1689/remote-pi/controller"
	"github.com/just1689/remote-pi/controller/io"
	"github.com/sirupsen/logrus"
)

const configFile = "app-config.json"

func main() {
	logrus.Println("Application starting v1.0")
	start()
}

func start() {

	config := io.ReadConfig(configFile)
	logrus.Println("..config was loaded from file")
	controller.StartDaemon(config)
}
