package controller

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/just1689/remote-pi/controller/io"
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"github.com/sirupsen/logrus"
)

func StartDaemon(config model.Config) {

	if config.EnableGPIO {
		err := io.Startup()
		if err != nil {
			logrus.Fatalln(fmt.Sprintf("Could not start the GPIO library: %s", err.Error()))
		}
		io.SubscribePubSub(config, handleMessage)
	} else {
		io.SubscribePubSub(config, handleMessageNoGPIO)
	}

}

func handleMessage(_ context.Context, message *pubsub.Message) {
	util.LogMsg(string(message.Data))

	var pinMessage = model.PinMessage{}
	if err := util.BytesToDecoder(message.Data).Decode(&pinMessage); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	io.PinToggle(pinMessage.PinID, pinMessage.On)
	message.Ack()

}

func handleMessageNoGPIO(_ context.Context, message *pubsub.Message) {
	util.LogMsg(string(message.Data))
	message.Ack()

}
