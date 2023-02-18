package babyname

type Gender int64

const (
	Ikhwan Gender = iota
	Akhwat
)

func (g Gender) ToString() string {
	switch g {
	case Ikhwan:
		return "Ikhwan"
	case Akhwat:
		return "Akhwat"
	}

	return "Unknown"
}

type BabyName struct {
	Name   string `bson:"name"`
	Gender Gender `bson:"gender"`
}
