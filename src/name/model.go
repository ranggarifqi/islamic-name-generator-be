package name

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender string

const (
	Ikhwan Gender = "Ikhwan"
	Akhwat Gender = "Akhwat"
)

type NameType string

const (
	FirstName  NameType = "FirstName"
	MiddleName NameType = "MiddleName"
	LastName   NameType = "LastName"
)

type Name struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Gender    Gender             `bson:"gender"`
	NameTypes []NameType         `bson:"nameTypes"`
	Meanings  []string           `bson:"meanings"`
	CreatedAt time.Time          `bson:"createdAt"`
}
