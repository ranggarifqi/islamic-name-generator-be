package name

import (
	"time"
)

type Gender string

const (
	IKHWAN Gender = "IKHWAN"
	AKWHAT Gender = "AKHWAT"
)

type NameType string

const (
	FIRST_NAME  NameType = "FIRST_NAME"
	MIDDLE_NAME NameType = "MIDDLE_NAME"
	LAST_NAME   NameType = "LAST_NAME"
)

type Name struct {
	ID        string     `bson:"_id,omitempty"`
	Name      string     `bson:"name"`
	Gender    Gender     `bson:"gender"`
	NameTypes []NameType `bson:"nameTypes"`
	Meanings  []string   `bson:"meanings"`
	CreatedAt time.Time  `bson:"createdAt"`
}
