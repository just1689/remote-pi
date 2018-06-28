package model

type AppConfig struct {
	ProjectID          string              `json:"projectId" bson:"projectId"`
	CredentialsFile    string              `json:"credentialsFile" bson:"credentialsFile"`
	EnableGPIO         bool                `json:"enableGPIO" bson:"enableGPIO"`
	OutputSubscription OutputSubscription  `json:"outputSubscription" bson:"outputSubscription"`
	InputSubscriptions []InputSubscription `json:"inputSubscriptions" bson:"inputSubscriptions"`
}

/*
	OutputSubscription describes a queue topic and subscription to use as input
	to control output pins on the board
 */
type OutputSubscription struct {
	TopicName        string `json:"topicName" bson:"topicName"`
	SubscriptionName string `json:"subscriptionName" bson:"subscriptionName"`
}

/*
	InputSubscription is used to describe a subscription for a particular pin.
	This pin will be read every IntervalMs and written to queue TopicName
	on a subscription called SubscriptionName
 */
type InputSubscription struct {
	PinID            int    `json:"pinId" bson:"pinId"`
	TopicName        string `json:"topicName" bson:"topicName"`
	SubscriptionName string `json:"subscriptionName" bson:"subscriptionName"`
	IntervalMs       string `json:"intervalMs" bson:"intervalMs"`
}
