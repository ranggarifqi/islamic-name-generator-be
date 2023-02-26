package name

import (
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

func (s *service) UpsertName(payload Name) (*Name, error) {
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

func (s *service) GenerateName(payload GenerateNameDTO) (string, error) {
	panic("not implemented") // TODO: Implement
}
