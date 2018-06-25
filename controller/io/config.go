package io

import (
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"io/ioutil"
	"log"
)

func ReadConfig(f string) (config model.Config) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal("Failed to read file config.json: ", err.Error())
	}

	if err := util.BytesToDecoder(bytes).Decode(&config); err != nil {
		log.Fatal("Failed to load config.json: ", err.Error())
	}
	return
}
