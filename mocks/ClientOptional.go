// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	rpc "github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

// ClientOptional is an autogenerated mock type for the ClientOptional type
type ClientOptional struct {
	mock.Mock
}

// Execute provides a mock function with given fields: o
func (_m *ClientOptional) Execute(o *rpc.ClientOptions) {
	_m.Called(o)
}
