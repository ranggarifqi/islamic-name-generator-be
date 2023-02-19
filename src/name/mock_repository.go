package name

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) findBy(filter FindByFilter) (*[]Name, error) {
	args := r.Called(filter)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	dummyName := ConstructDummyName(Name{})

	return &[]Name{dummyName}, nil
}

func (r *MockRepository) findById(id string) (*Name, error) {
	args := r.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	dummyName := ConstructDummyName(Name{})

	return &dummyName, nil
}

func (r *MockRepository) create(payload Name) (*Name, error) {
	args := r.Called(payload)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	dummyName := ConstructDummyName(payload)

	return &dummyName, nil
}
