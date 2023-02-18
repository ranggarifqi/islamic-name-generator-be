package name

import (
	"time"
)

type Gender string

const (
	IKHWAN Gender = "Ikhwan"
	AKWHAT Gender = "Akhwat"
)

type NameType string

const (
	FIRST_NAME  NameType = "FirstName"
	MIDDLE_NAME NameType = "MiddleName"
	LAST_NAME   NameType = "LastName"
)

type Name struct {
	ID        string     `bson:"_id,omitempty"`
	Name      string     `bson:"name"`
	Gender    Gender     `bson:"gender"`
	NameTypes []NameType `bson:"nameTypes"`
	Meanings  []string   `bson:"meanings"`
	CreatedAt time.Time  `bson:"createdAt"`
}
