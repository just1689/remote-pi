package io

import (
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/just1689/remote-pi/model"
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
