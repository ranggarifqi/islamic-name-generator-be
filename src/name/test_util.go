package name

import (
	"time"

	"github.com/ranggarifqi/islamic-name-generator-be/helper"
)

func ConstructDummyName(payload Name) Name {
	name, _ := helper.GetStructValue(payload, "Name", "Fulan")
	gender, _ := helper.GetStructValue(payload, "Gender", IKHWAN)

	if len(payload.NameTypes) == 0 {
		payload.NameTypes = []NameType{FIRST_NAME, MIDDLE_NAME}
	}

	if len(payload.Meanings) == 0 {
		payload.Meanings = []string{"Baik", "Penyayang"}
	}

	return Name{
		ID:        "someId",
		Name:      *name,
		NameTypes: payload.NameTypes,
		Gender:    *gender,
		Meanings:  payload.Meanings,
		CreatedAt: time.Now(),
	}
}
