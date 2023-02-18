package name

import (
	"time"
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
	ID        string     `bson:"_id,omitempty"`
	Name      string     `bson:"name"`
	Gender    Gender     `bson:"gender"`
	NameTypes []NameType `bson:"nameTypes"`
	Meanings  []string   `bson:"meanings"`
	CreatedAt time.Time  `bson:"createdAt"`
}
