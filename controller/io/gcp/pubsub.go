package gcp

import (
	"log"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"time"
)

func Subscribe(projectID string, topicName string, credentialsFile string, subName string, handleMessage func(context.Context, *pubsub.Message)) (err error) {

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Println("Failed to create pub sub client", err.Error())
		return
	}

	topic := client.Topic(topicName)

	sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})
	if err != nil {
		log.Println("Failed to create subscription", err.Error())
		return
	}
	sub.Receive(ctx, handleMessage)

	return

}
