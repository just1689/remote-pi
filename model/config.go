package model

type Config struct {
	ProjectID       string `json:"projectId" bson:"projectId"`
	TopicName       string `json:"topicName" bson:"topicName"`
	CredentialsFile string `json:"credentialsFile" bson:"credentialsFile"`
}
