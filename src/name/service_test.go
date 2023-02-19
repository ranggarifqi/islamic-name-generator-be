package name

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NameService_upsertName(t *testing.T) {
	payload := Name{
		Name:      "fulan",
		Gender:    IKHWAN,
		NameTypes: []NameType{FIRST_NAME, MIDDLE_NAME},
		Meanings:  []string{"dia", "someone"},
	}

	t.Run("Should return early if there's an error when finding existing name", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository)

		repository.On("findBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(nil, errors.New("some DB error when finding the data"))

		result, err := service.upsertName(payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "some DB error when finding the data")
	})

	t.Run("Should create the data directly if there's no existing name", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository)

		repository.On("findBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{}, nil)

		expectedCreatedName := ConstructDummyName(payload)

		repository.On("create", payload).Return(expectedCreatedName, nil)

		result, err := service.upsertName(payload)

		assert.True(t, repository.AssertCalled(t, "create", payload))

		assert.Nil(t, err)
		assert.Equal(t, "someId", result.ID)
		assert.Equal(t, expectedCreatedName.Name, result.Name)
		assert.Equal(t, expectedCreatedName.Gender, result.Gender)
		assert.ElementsMatch(t, result.NameTypes, expectedCreatedName.NameTypes)
		assert.ElementsMatch(t, result.Meanings, expectedCreatedName.Meanings)

	})
	t.Run("Should return error on unsuccessful creation", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository)

		repository.On("findBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{}, nil)

		repository.On("create", payload).Return(nil, errors.New("Some DB Error when creating document"))

		result, err := service.upsertName(payload)

		assert.True(t, repository.AssertCalled(t, "create", payload))

		assert.Nil(t, result)
		assert.NotNil(t, err)

		assert.EqualError(t, err, "Some DB Error when creating document")
	})

	t.Run("Should update existing data if found", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository)

		existingData := ConstructDummyName(Name{
			ID:        "someId",
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: []NameType{MIDDLE_NAME},
			Meanings:  []string{"someone"},
		})

		repository.On("findBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{existingData}, nil)

		repository.On("updateById", existingData.ID, mock.Anything).Return(&Name{
			ID: existingData.ID,
		}, nil)

		result, err := service.upsertName(payload)

		assert.True(t, repository.AssertCalled(t, "updateById", existingData.ID, Name{
			NameTypes: []NameType{MIDDLE_NAME, FIRST_NAME},
			Meanings:  []string{"someone", "dia"},
		}))

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	// t.Run("should return an error on unsucesful update", func(t *testing.T) {})
}
