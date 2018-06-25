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

func StartDaemon(projectID string, topicName string, credentialsFile string) {
	gpio.Startup()
	gcp.Subscribe(projectID, topicName, credentialsFile, handleMessage)
}

func handleMessage(_ context.Context, message *pubsub.Message) {
	var pinMessage = model.PinMessage{}
	if err := util.BytesToDecoder(message.Data).Decode(&pinMessage); err != nil {
		logrus.Errorln(fmt.Sprintf("There was a problem decoding the post message: %s", err.Error()))
	}
	gpio.PinToggle(pinMessage.PinID, pinMessage.On)

}
