// Code generated by mockery v2.20.0.

package name

import (
	"fmt"

	mock "github.com/stretchr/testify/mock"
)

// MockINameRepository is an autogenerated mock type for the INameRepository type
type MockINameRepository struct {
	mock.Mock
}

// create provides a mock function with given fields: payload
func (_m *MockINameRepository) create(payload Name) (*Name, error) {
	ret := _m.Called(payload)

	var r0 *Name
	var r1 error
	if rf, ok := ret.Get(0).(func(Name) (*Name, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(Name) *Name); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			fmt.Printf("ret.Get(0) = %v", ret.Get(0))
			temp := ret.Get(0).(Name)
			r0 = &temp
		}
	}

	if rf, ok := ret.Get(1).(func(Name) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockINameRepository_create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'create'
type MockINameRepository_create_Call struct {
	*mock.Call
}

// create is a helper method to define mock.On call
//   - payload Name

func (_c *MockINameRepository_create_Call) Run(run func(payload Name)) *MockINameRepository_create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Name))
	})
	return _c
}

func (_c *MockINameRepository_create_Call) Return(_a0 *Name, _a1 error) *MockINameRepository_create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockINameRepository_create_Call) RunAndReturn(run func(Name) (*Name, error)) *MockINameRepository_create_Call {
	_c.Call.Return(run)
	return _c
}

// findBy provides a mock function with given fields: filter
func (_m *MockINameRepository) findBy(filter FindByFilter) (*[]Name, error) {
	ret := _m.Called(filter)

	var r0 *[]Name
	var r1 error
	if rf, ok := ret.Get(0).(func(FindByFilter) (*[]Name, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(FindByFilter) *[]Name); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]Name)
		}
	}

	if rf, ok := ret.Get(1).(func(FindByFilter) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockINameRepository_findBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'findBy'
type MockINameRepository_findBy_Call struct {
	*mock.Call
}

// findBy is a helper method to define mock.On call
//   - filter FindByFilter

func (_c *MockINameRepository_findBy_Call) Run(run func(filter FindByFilter)) *MockINameRepository_findBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(FindByFilter))
	})
	return _c
}

func (_c *MockINameRepository_findBy_Call) Return(_a0 *[]Name, _a1 error) *MockINameRepository_findBy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockINameRepository_findBy_Call) RunAndReturn(run func(FindByFilter) (*[]Name, error)) *MockINameRepository_findBy_Call {
	_c.Call.Return(run)
	return _c
}

// findById provides a mock function with given fields: id
func (_m *MockINameRepository) findById(id string) (*Name, error) {
	ret := _m.Called(id)

	var r0 *Name
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*Name, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *Name); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Name)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockINameRepository_findById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'findById'
type MockINameRepository_findById_Call struct {
	*mock.Call
}

// findById is a helper method to define mock.On call
//   - id string

func (_c *MockINameRepository_findById_Call) Run(run func(id string)) *MockINameRepository_findById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockINameRepository_findById_Call) Return(_a0 *Name, _a1 error) *MockINameRepository_findById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockINameRepository_findById_Call) RunAndReturn(run func(string) (*Name, error)) *MockINameRepository_findById_Call {
	_c.Call.Return(run)
	return _c
}

// updateById provides a mock function with given fields: id, payload
func (_m *MockINameRepository) updateById(id string, payload Name) (*Name, error) {
	ret := _m.Called(id, payload)

	var r0 *Name
	var r1 error
	if rf, ok := ret.Get(0).(func(string, Name) (*Name, error)); ok {
		return rf(id, payload)
	}
	if rf, ok := ret.Get(0).(func(string, Name) *Name); ok {
		r0 = rf(id, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Name)
		}
	}

	if rf, ok := ret.Get(1).(func(string, Name) error); ok {
		r1 = rf(id, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockINameRepository_updateById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'updateById'
type MockINameRepository_updateById_Call struct {
	*mock.Call
}

// updateById is a helper method to define mock.On call
//   - id string
//   - payload Name

func (_c *MockINameRepository_updateById_Call) Run(run func(id string, payload Name)) *MockINameRepository_updateById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(Name))
	})
	return _c
}

func (_c *MockINameRepository_updateById_Call) Return(_a0 *Name, _a1 error) *MockINameRepository_updateById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockINameRepository_updateById_Call) RunAndReturn(run func(string, Name) (*Name, error)) *MockINameRepository_updateById_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockINameRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockINameRepository creates a new instance of MockINameRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockINameRepository(t mockConstructorTestingTNewMockINameRepository) *MockINameRepository {
	mock := &MockINameRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
