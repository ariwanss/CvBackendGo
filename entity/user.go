package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" form:"username" bson:"username,omitempty"`
	Password string             `json:"-" form:"password" bson:"password,omitempty"`
	Name     string             `json:"name" form:"name" bson:"name,omitempty"`
	DOB      time.Time          `json:"dob" form:"dob" bson:"dob,omitempty"`
	Email    string             `json:"email" form:"email" bson:"email,omitempty"`
	Phone    string             `json:"phone" form:"phone" bson:"phone,omitempty"`
	Address  string             `json:"address" form:"address" bson:"address,omitempty"`
	GitHub   string             `json:"github" form:"github" bson:"github,omitempty"`
	LinkedIn string             `json:"linkedin" form:"linkedin" bson:"linkedin,omitempty"`
	Summary  string             `json:"summary" form:"summary" bson:"summary,omitempty"`
}
