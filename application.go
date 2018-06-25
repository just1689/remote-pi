package main

import (
	"fmt"
	"github.com/just1689/remote-pi/controller/daemon"
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"io/ioutil"
	"log"
)

const configFile = "config.json"

func main() {
	fmt.Println("Application starting")

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Failed to read file config.json: ", err.Error())
	}

	var config model.Config
	if err := util.BytesToDecoder(bytes).Decode(&config); err != nil {
		log.Fatal("Failed to load config.json: ", err.Error())
	}

	daemon.StartDaemon(config)
}
