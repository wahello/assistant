// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	template "html/template"

	mock "github.com/stretchr/testify/mock"
)

// Component is an autogenerated mock type for the Component type
type Component struct {
	mock.Mock
}

// GetContent provides a mock function with given fields:
func (_m *Component) GetContent() template.HTML {
	ret := _m.Called()

	var r0 template.HTML
	if rf, ok := ret.Get(0).(func() template.HTML); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(template.HTML)
	}

	return r0
}
