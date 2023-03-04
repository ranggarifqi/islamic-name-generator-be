package v1Name

import "github.com/ranggarifqi/islamic-name-generator-be/src/name"

type GenerateNameDTO struct {
	Gender              name.Gender `json:"gender" validate:"required"`
	ShouldUseMiddleName *bool       `json:"shouldUseMiddleName" validate:"required,boolean"`
	ShouldUseLastName   *bool       `json:"shouldUseLastName" validate:"required,boolean"`
}
