package io

import (
	"log"

	"cloud.google.com/go/pubsub"
	"encoding/json"
	"fmt"
	"github.com/just1689/remote-pi/model"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func SubscribePubSub(config model.AppConfig, handleMessage func(context.Context, *pubsub.Message)) (err error) {

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, config.ProjectID, option.WithCredentialsFile(config.OutputSubscription.CredentialsFile))
	if err != nil {
		log.Println("Failed to create pub sub client", err.Error())
		return
	}
	sub := client.Subscription(config.OutputSubscription.SubscriptionName)
	sub.Receive(ctx, handleMessage)

	return

}

func GetPubTopic(config model.AppConfig, c model.InputSubscription) (topic *pubsub.Topic, err error) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.ProjectID, option.WithCredentialsFile(config.OutputSubscription.CredentialsFile))
	if err != nil {
		logrus.Errorln(fmt.Sprintf("Failed to create pub sub client", err.Error()))
		return
	}
	topic = client.Topic(c.Topic)
	return

}

func PubToTopic(topic *pubsub.Topic, i interface{}) (err error) {
	ctx := context.Background()
	b, err := json.Marshal(i)
	if err != nil {
		logrus.Errorln(fmt.Sprintf("Failed to marshal i: %s", err))
		return
	}
	_, err = topic.Publish(ctx, &pubsub.Message{Data: b}).Get(ctx)
	return

}
