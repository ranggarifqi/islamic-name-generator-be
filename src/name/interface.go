package name

/** Service Interfaces */
type INameService interface {
	createName(payload *Name) (*Name, error)
	generateName(payload *GenerateNameDTO) (string, error)
}

type GenerateNameDTO struct {
	Gender              Gender
	ShouldUseMiddleName bool
	ShouldUseLastName   bool
}

/** Repo Interfaces */
type INameRepository interface {
	findBy(filter FindByFilter) (*[]Name, error)
	findById(id string) (*Name, error)

	create(payload Name) (*Name, error)
}

type FindByFilter struct {
	Gender    Gender
	NameTypes []NameType
}
