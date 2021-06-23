// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"

	sdk "github.com/tsundata/assistant/internal/pkg/sdk"
)

// OAuthProvider is an autogenerated mock type for the OAuthProvider type
type OAuthProvider struct {
	mock.Mock
}

// AuthorizeURL provides a mock function with given fields:
func (_m *OAuthProvider) AuthorizeURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAccessToken provides a mock function with given fields: code
func (_m *OAuthProvider) GetAccessToken(code string) (interface{}, error) {
	ret := _m.Called(code)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Redirect provides a mock function with given fields: c, gateway
func (_m *OAuthProvider) Redirect(c *fiber.Ctx, gateway *sdk.GatewayClient) error {
	ret := _m.Called(c, gateway)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *sdk.GatewayClient) error); ok {
		r0 = rf(c, gateway)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreAccessToken provides a mock function with given fields: c, gateway
func (_m *OAuthProvider) StoreAccessToken(c *fiber.Ctx, gateway *sdk.GatewayClient) error {
	ret := _m.Called(c, gateway)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *sdk.GatewayClient) error); ok {
		r0 = rf(c, gateway)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
