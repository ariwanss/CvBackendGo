package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CvItem struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	Section      string             `json:"section" form:"section" bson:"section,omitempty"`
	Title        string             `json:"title" form:"title" bson:"title,omitempty"`
	DateStart    time.Time          `json:"dateStart" form:"dateStart" bson:"dateStart,omitempty"`
	DateEnd      time.Time          `json:"dateEnd" form:"dateEnd" bson:"dateEnd,omitempty"`
	Link         string             `json:"link" form:"link" bson:"link,omitempty"`
	Descriptions []string           `json:"descriptions" form:"descriptions" bson:"descriptions,omitempty"`
}
