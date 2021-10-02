package UserDS

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	SecondName string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	CreatedOn  time.Time          `json:"createdon" bson:"createdon"`
	UpdateOn   time.Time          `json:"updateon" bson:"updateon"`
	Mobile     string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Active     bool               `json:"active,omitempty" bson:"active,omitempty"`
	Age        AgeDS              `json:"age,omitempty"  bson:"age,omitempty"`
}

type AgeDS struct {
	Value    int    `json:"age,omitempty" bson:"age,omitempty"`
	Interval string `json:"interval,omitempty" bson:"interval,omitempty"`
}
