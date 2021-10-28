package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Receiver struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User string             `bson:"user"`
	Name string             `bson:"name"`
}

type Template struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Raw    string             `bson:"raw"`
	Params []string           `bson:"params"`
}
