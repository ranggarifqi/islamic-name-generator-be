package name

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NameService_UpsertName(t *testing.T) {
	payload := Name{
		Name:      "fulan",
		Gender:    IKHWAN,
		NameTypes: []NameType{FIRST_NAME, MIDDLE_NAME},
		Meanings:  []string{"dia", "someone"},
	}

	randomizer := rand.New(rand.NewSource(1))

	t.Run("Should return early if there's an error when finding existing name", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		repository.On("FindBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(nil, errors.New("some DB error when finding the data"))

		result, err := service.UpsertName(payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "some DB error when finding the data")
	})

	t.Run("Should create the data directly if there's no existing name", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		repository.On("FindBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{}, nil)

		expectedCreatedName := ConstructDummyName(payload)

		repository.On("Create", payload).Return(expectedCreatedName, nil)

		result, err := service.UpsertName(payload)

		repository.AssertCalled(t, "Create", payload)

		assert.Nil(t, err)
		assert.Equal(t, "someId", result.ID)
		assert.Equal(t, expectedCreatedName.Name, result.Name)
		assert.Equal(t, expectedCreatedName.Gender, result.Gender)
		assert.ElementsMatch(t, result.NameTypes, expectedCreatedName.NameTypes)
		assert.ElementsMatch(t, result.Meanings, expectedCreatedName.Meanings)

	})
	t.Run("Should return error on unsuccessful creation", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		repository.On("FindBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{}, nil)

		repository.On("Create", payload).Return(nil, errors.New("Some DB Error when creating document"))

		result, err := service.UpsertName(payload)

		repository.AssertCalled(t, "Create", payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)

		assert.EqualError(t, err, "Some DB Error when creating document")
	})

	t.Run("Should update existing data if found", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		existingData := ConstructDummyName(Name{
			ID:        "someId",
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: []NameType{MIDDLE_NAME},
			Meanings:  []string{"someone"},
		})

		repository.On("FindBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{existingData}, nil)

		repository.On("UpdateById", existingData.ID, mock.Anything).Return(&Name{
			ID: existingData.ID,
		}, nil)

		result, err := service.UpsertName(payload)

		repository.AssertCalled(t, "UpdateById", existingData.ID, Name{
			NameTypes: []NameType{FIRST_NAME, MIDDLE_NAME},
			Meanings:  []string{"dia", "someone"},
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("should return an error on unsucesful update", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		existingData := ConstructDummyName(Name{
			ID:        "someId",
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: []NameType{MIDDLE_NAME},
			Meanings:  []string{"someone"},
		})

		repository.On("FindBy", FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		}).Return(&[]Name{existingData}, nil)

		repository.On("UpdateById", existingData.ID, mock.Anything).Return(nil, errors.New("Some Error when updating the document"))

		result, err := service.UpsertName(payload)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Some Error when updating the document")
	})
}

func Test_NameService_GenerateName(t *testing.T) {
	randomizer := rand.New(rand.NewSource(1))

	dto := GenerateNameDTO{
		Gender:              IKHWAN,
		ShouldUseMiddleName: false,
		ShouldUseLastName:   true,
	}

	t.Run("should handle error on DB query error", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		repository.On("FindBy", mock.Anything).Return(nil, errors.New("some DB error when fetching the data"))

		result, err := service.GenerateName(dto)

		assert.EqualValues(t, map[NameType]Name{}, result)
		assert.NotNil(t, err)

		assert.EqualError(t, err, "some DB error when fetching the data")
	})

	t.Run("should generate first name & last name correctly", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		findByResults := []Name{
			{
				ID:        "firstName1",
				Name:      "andi",
				Gender:    IKHWAN,
				NameTypes: []NameType{FIRST_NAME},
			},
			{
				ID:        "middleName1",
				Name:      "putra",
				Gender:    IKHWAN,
				NameTypes: []NameType{MIDDLE_NAME},
			},
			{
				ID:        "lastName1",
				Name:      "sulistyo",
				Gender:    IKHWAN,
				NameTypes: []NameType{LAST_NAME},
			},
		}

		repository.On("FindBy", mock.Anything).Return(&findByResults, nil)

		result, err := service.GenerateName(dto)

		assert.Nil(t, err)
		assert.Equal(t, "andi", result[FIRST_NAME].Name)
		assert.Equal(t, "sulistyo", result[LAST_NAME].Name)
	})

	t.Run("should not pick the same name if it's been picked previously", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		findByResults := []Name{
			{
				ID:        "firstName1",
				Name:      "andi",
				Gender:    IKHWAN,
				NameTypes: []NameType{FIRST_NAME, LAST_NAME},
			},
			{
				ID:        "lastName1",
				Name:      "sulistyo",
				Gender:    IKHWAN,
				NameTypes: []NameType{LAST_NAME},
			},
		}

		repository.On("FindBy", mock.Anything).Return(&findByResults, nil)

		result, err := service.GenerateName(dto)

		assert.Nil(t, err)
		assert.Equal(t, "andi", result[FIRST_NAME].Name)
		assert.NotEqual(t, "andi", result[LAST_NAME].Name)
	})

	t.Run("Should return empty map if no name with the given nameType exists", func(t *testing.T) {
		repository := &MockINameRepository{}

		service := NewService(repository, randomizer)

		findByResults := []Name{
			{
				ID:        "firstName1",
				Name:      "andi",
				Gender:    IKHWAN,
				NameTypes: []NameType{MIDDLE_NAME},
			},
			{
				ID:        "lastName1",
				Name:      "sulistyo",
				Gender:    IKHWAN,
				NameTypes: []NameType{MIDDLE_NAME},
			},
		}

		repository.On("FindBy", mock.Anything).Return(&findByResults, nil)

		// Expect first name & last name being returned. But they don't exist.
		result, err := service.GenerateName(dto)

		assert.Nil(t, err)
		assert.EqualValues(t, map[NameType]Name{}, result)
	})
}

/** Helper Methods */
func Test_NameService_ChooseRandomizedName(t *testing.T) {
	namesArr := []Name{
		{
			Name: "anto",
		},
		{
			Name: "budi",
		},
		{
			Name: "cici",
		},
	}

	repository := &MockINameRepository{}
	randomizer := rand.New(rand.NewSource(1))

	t.Run("Should randomize the returned name without exception", func(t *testing.T) {
		service := &Service{
			randomizer:     randomizer,
			nameRepository: repository,
		}

		result := service.ChooseRandomizedName(namesArr, []string{})

		isCorrect := lo.ContainsBy(namesArr, func(item Name) bool {
			return item.Name == result.Name
		})

		assert.True(t, isCorrect)
	})

	t.Run("Should not return an excluded name", func(t *testing.T) {
		service := &Service{
			randomizer:     randomizer,
			nameRepository: repository,
		}

		result := service.ChooseRandomizedName(namesArr, []string{"anto"})

		assert.NotEqual(t, "anto", result.Name)
	})
}

/** Helper funcs */
func Test_ConstructNameTypes(t *testing.T) {
	type TestCase struct {
		shouldUseMiddleName bool
		shouldUseLastName   bool
		expected            []NameType
	}

	testCases := []TestCase{
		{
			shouldUseMiddleName: false,
			shouldUseLastName:   false,
			expected:            []NameType{FIRST_NAME},
		},
		{
			shouldUseMiddleName: true,
			shouldUseLastName:   false,
			expected:            []NameType{FIRST_NAME, MIDDLE_NAME},
		},
		{
			shouldUseMiddleName: false,
			shouldUseLastName:   true,
			expected:            []NameType{FIRST_NAME, LAST_NAME},
		},
		{
			shouldUseMiddleName: true,
			shouldUseLastName:   true,
			expected:            []NameType{FIRST_NAME, MIDDLE_NAME, LAST_NAME},
		},
	}

	for i, tc := range testCases {
		expectedStr := lo.Reduce(tc.expected, func(agg string, item NameType, idx int) string {
			return fmt.Sprintf("%v, %v", agg, item)
		}, "")

		t.Run(fmt.Sprintf("(%v) should return (%v) if shouldUseMiddleName = %v & shouldUseLastName = %v", i, expectedStr, tc.shouldUseMiddleName, tc.shouldUseLastName), func(t *testing.T) {
			result := ConstructNameTypes(tc.shouldUseMiddleName, tc.shouldUseLastName)
			assert.Len(t, result, len(tc.expected))
			assert.True(t, reflect.DeepEqual(tc.expected, result))
		})
	}
}
