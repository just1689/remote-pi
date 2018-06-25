package daemon

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/just1689/remote-pi/controller/io/gcp"
	"github.com/just1689/remote-pi/controller/io/gpio"
	"github.com/just1689/remote-pi/model"
	"github.com/just1689/remote-pi/util"
	"github.com/sirupsen/logrus"
)

func StartDaemon(config model.Config) {
	if config.EnableGPIO {
		gpio.Startup()
		gcp.Subscribe(config.ProjectID, config.TopicName, config.SubscriptionName, config.CredentialsFile, handleMessage)
	} else {
		gcp.Subscribe(config.ProjectID, config.TopicName, config.SubscriptionName, config.CredentialsFile, handleMessageNoGPIO)
	}

}

func handleMessage(_ context.Context, message *pubsub.Message) {

	var pinMessage = model.PinMessage{}
	if err := util.BytesToDecoder(message.Data).Decode(&pinMessage); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	gpio.PinToggle(pinMessage.PinID, pinMessage.On)
	message.Ack()

}

func handleMessageNoGPIO(_ context.Context, message *pubsub.Message) {
	message.Ack()

	fmt.Println(">>> New message")
	fmt.Println(string(message.Data))
	fmt.Println("<<< EOM")

}
