package controller

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/just1689/remote-pi/controller/io"
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"github.com/sirupsen/logrus"
	"time"
)

func StartDaemon(config model.AppConfig) {

	if config.EnableGPIO {
		err := io.Startup()
		if err != nil {
			logrus.Fatalln(fmt.Sprintf("Could not start the GPIO library: %s", err.Error()))
		}
	}

	if len(config.InputSubscriptions) > 0 {
		logrus.Println("> Starting writers: %i", len(config.InputSubscriptions))
		go startWriter(config)
	}

	if config.OutputSubscription.On {
		logrus.Println("> Starting reader")
		startReader(config)
	}

}

func startReader(config model.AppConfig) {
	if config.EnableGPIO {
		io.SubscribePubSub(config, handleMessage)
	} else {
		io.SubscribePubSub(config, handleMessageNoGPIO)
	}
}

func startWriter(config model.AppConfig) {

	for _, c := range config.InputSubscriptions {
		go monitorPin(config, c)
	}
}

func monitorPin(config model.AppConfig, c model.InputSubscription) {
	topic, err := io.GetPubTopic(config, c)
	if err != nil {
		logrus.Fatalln(fmt.Sprintf("Could not subscribe to topic %s. Got error: %s", c.Topic, err.Error()))
	}
	sleepDur := time.Duration(c.IntervalMs) * time.Millisecond
	for {
		if !config.EnableGPIO {
			time.Sleep(sleepDur)
			logrus.Println(fmt.Sprintf("Would have read pin %v, but config told me not to", c.PinID))
			time.Sleep(sleepDur)
			continue
		}
		io.PubToTopic(
			topic,
			model.PinMessage{
				PinID: c.PinID,
				On:    io.PinRead(c.PinID)})
		time.Sleep(sleepDur)
	}

}

func handleMessage(_ context.Context, message *pubsub.Message) {
	logrus.Info(">>> New message:")
	logrus.Info(string(message.Data))

	var pinMessage = model.PinMessage{}
	if err := util.BytesToDecoder(message.Data).Decode(&pinMessage); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	io.PinToggle(pinMessage.PinID, pinMessage.On)
	message.Ack()

}

func handleMessageNoGPIO(_ context.Context, message *pubsub.Message) {
	logrus.Info(">>> New message:")
	logrus.Info(string(message.Data))
	message.Ack()

}
