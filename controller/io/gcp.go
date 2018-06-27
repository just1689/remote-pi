package io

import (
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func Subscribe(projectID string, topicName string, subName string, credentialsFile string, handleMessage func(context.Context, *pubsub.Message)) (err error) {

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Println("Failed to create pub sub client", err.Error())
		return
	}
	sub := client.Subscription(subName)
	sub.Receive(ctx, handleMessage)

	return

}
