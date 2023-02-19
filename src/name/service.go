package name

import (
	"fmt"

	"github.com/ranggarifqi/islamic-name-generator-be/helper"
)

type service struct {
	nameRepository INameRepository
}

func NewService(nameRepository INameRepository) INameService {
	return &service{
		nameRepository: nameRepository,
	}
}

func (s *service) upsertName(payload Name) (*Name, error) {
	// Find in collection that has the same name.
	foundNames, err := s.nameRepository.findBy(
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
		newName, err := s.nameRepository.create(payload)
		if err != nil {
			return nil, err
		}
		return newName, nil
	}

	name := (*foundNames)[0]

	newMeanings := append(name.Meanings, payload.Meanings...)
	fmt.Printf("newMeanings before remove dupe = %v\n", newMeanings)
	newMeanings = helper.RemoveSliceDuplicate(newMeanings)
	fmt.Printf("newMeanings after remove dupe = %v\n", newMeanings)

	newNameTypes := append(name.NameTypes, payload.NameTypes...)
	newNameTypes = helper.RemoveSliceDuplicate(newNameTypes)

	updatedName, err := s.nameRepository.updateById(name.ID, Name{
		NameTypes: newNameTypes,
		Meanings:  newMeanings,
	})
	if err != nil {
		return nil, err
	}

	return updatedName, nil
}

func (s *service) generateName(payload GenerateNameDTO) (string, error) {
	panic("not implemented") // TODO: Implement
}
