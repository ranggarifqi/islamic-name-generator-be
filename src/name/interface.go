package name

/** Service Interfaces */
type INameService interface {
	UpsertName(payload Name) (*Name, error)
	GenerateName(payload GenerateNameDTO) (map[NameType]Name, error)
}

type GenerateNameDTO struct {
	Gender              Gender
	ShouldUseMiddleName bool
	ShouldUseLastName   bool
}

/** Repo Interfaces */
type INameRepository interface {
	FindBy(filter FindByFilter) (*[]Name, error)
	FindById(id string) (*Name, error)

	Create(payload Name) (*Name, error)

	UpdateById(id string, payload Name) (*Name, error)
}

type FindByFilter struct {
	Name      string
	Gender    Gender
	NameTypes []NameType
}
