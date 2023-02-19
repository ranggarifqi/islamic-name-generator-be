package name

import (
	"time"

	"github.com/ranggarifqi/islamic-name-generator-be/helper"
)

func ConstructDummyName(payload Name) Name {
	name, _ := helper.GetStructValue(payload, "Name", "Fulan")
	nameTypes, _ := helper.GetStructValue(payload, "NameTypes", []NameType{FIRST_NAME, MIDDLE_NAME})
	gender, _ := helper.GetStructValue(payload, "Gender", IKHWAN)
	meanings, _ := helper.GetStructValue(payload, "Meanings", []string{"Baik", "Penyayang"})

	return Name{
		ID:        "someId",
		Name:      *name,
		NameTypes: *nameTypes,
		Gender:    *gender,
		Meanings:  *meanings,
		CreatedAt: time.Now(),
	}
}
