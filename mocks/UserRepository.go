// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/tsundata/assistant/internal/pkg/model"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// ChangeRoleAttr provides a mock function with given fields: userID, attr, val
func (_m *UserRepository) ChangeRoleAttr(userID int, attr string, val int) error {
	ret := _m.Called(userID, attr, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, int) error); ok {
		r0 = rf(userID, attr, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeRoleExp provides a mock function with given fields: userID, exp
func (_m *UserRepository) ChangeRoleExp(userID int, exp int) error {
	ret := _m.Called(userID, exp)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(userID, exp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRole provides a mock function with given fields: userID
func (_m *UserRepository) GetRole(userID int) (model.Role, error) {
	ret := _m.Called(userID)

	var r0 model.Role
	if rf, ok := ret.Get(0).(func(int) model.Role); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(model.Role)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
