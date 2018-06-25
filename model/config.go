package model

type Config struct {
	ProjectID        string `json:"projectId" bson:"projectId"`
	TopicName        string `json:"topicName" bson:"topicName"`
	SubscriptionName string `json:"subscriptionName" bson:"subscriptionName"`
	CredentialsFile  string `json:"credentialsFile" bson:"credentialsFile"`
}
