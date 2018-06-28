package io

import (
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"io/ioutil"
	"log"
)

func ReadConfig(f string) (config model.AppConfig) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal("Failed to read file app-config.json: ", err.Error())
	}

	if err := util.BytesToDecoder(bytes).Decode(&config); err != nil {
		log.Fatal("Failed to load app-config.json: ", err.Error())
	}
	return
}
