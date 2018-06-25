package model

type PinMessage struct {
	PinID string `json:"pinId" bson:"pinId"`
	On    bool   `json:"on" bson:"on"`
}
