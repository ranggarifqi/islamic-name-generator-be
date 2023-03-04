package name

import (
	"fmt"
	"math/rand"

	"github.com/ranggarifqi/islamic-name-generator-be/helper"
	"github.com/samber/lo"
)

type Service struct {
	nameRepository INameRepository
	randomizer     *rand.Rand
}

func NewService(nameRepository INameRepository, randomizer *rand.Rand) INameService {
	return &Service{
		nameRepository: nameRepository,
		randomizer:     randomizer,
	}
}

func (s *Service) UpsertName(payload Name) (*Name, error) {
	// Find in collection that has the same name.
	foundNames, err := s.nameRepository.FindBy(
		FindByFilter{
			Name:      payload.Name,
			Gender:    payload.Gender,
			NameTypes: payload.NameTypes,
		},
	)
	if err != nil {
		return nil, err
	}

	// If no data found, then just create it
	if len(*foundNames) == 0 {
		newName, err := s.nameRepository.Create(payload)
		if err != nil {
			return nil, err
		}
		return newName, nil
	}

	name := (*foundNames)[0]

	newMeanings := append(name.Meanings, payload.Meanings...)
	newMeanings = helper.RemoveSliceDuplicate(newMeanings)

	newNameTypes := append(name.NameTypes, payload.NameTypes...)
	newNameTypes = helper.RemoveSliceDuplicate(newNameTypes)

	updatedName, err := s.nameRepository.UpdateById(name.ID, Name{
		NameTypes: newNameTypes,
		Meanings:  newMeanings,
	})
	if err != nil {
		return nil, err
	}

	return updatedName, nil
}

func (s *Service) GenerateName(payload GenerateNameDTO) (map[NameType]Name, error) {
	// Construct name types
	nameTypes := ConstructNameTypes(payload.ShouldUseLastName, payload.ShouldUseMiddleName)

	// Get By filter
	names, err := s.nameRepository.FindBy(FindByFilter{
		Gender:    payload.Gender,
		NameTypes: nameTypes,
	})
	if err != nil {
		return map[NameType]Name{}, err
	}

	result := map[NameType]Name{}

	for _, nameType := range nameTypes {
		// Filter array, get FIRST_NAME
		names := lo.Filter(*names, func(item Name, idx int) bool {
			return lo.Contains(item.NameTypes, nameType)
		})

		exceptions := lo.Values(result)
		exceptionsStr := lo.Map(exceptions, func(item Name, idx int) string {
			return item.Name
		})

		// Randomize First Name
		choosenName := s.ChooseRandomizedName(names, exceptionsStr)

		result[nameType] = choosenName
	}

	return result, nil
}

func (s *Service) ChooseRandomizedName(nameArr []Name, exceptions []string) Name {
	tempArr := make([]Name, len(nameArr))
	copy(tempArr, nameArr)

	if len(exceptions) > 0 {
		tempArr = lo.Filter(tempArr, func(item Name, index int) bool {
			return !lo.Contains(exceptions, item.Name)
		})
	}

	fmt.Printf("tempArr = %v\n", tempArr)

	arrLength := len(tempArr)
	choosenIdx := s.randomizer.Intn(arrLength)

	return tempArr[choosenIdx]
}

/** Helper Funcs */
func ConstructNameTypes(shouldUseMiddleName bool, shouldUseLastName bool) []NameType {
	nameTypes := []NameType{FIRST_NAME}

	if shouldUseMiddleName {
		nameTypes = append(nameTypes, MIDDLE_NAME)
	}

	if shouldUseLastName {
		nameTypes = append(nameTypes, LAST_NAME)
	}

	return nameTypes
}
