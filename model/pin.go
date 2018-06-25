package model

type PinMessage struct {
	PinID int  `json:"pinId" bson:"pinId"`
	On    bool `json:"on" bson:"on"`
}
