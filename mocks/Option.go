// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	app "github.com/tsundata/assistant/internal/pkg/app"
)

// Option is an autogenerated mock type for the Option type
type Option struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *Option) Execute(_a0 *app.Application) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*app.Application) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
