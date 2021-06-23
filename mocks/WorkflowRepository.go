// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "github.com/tsundata/assistant/internal/pkg/model"
)

// WorkflowRepository is an autogenerated mock type for the WorkflowRepository type
type WorkflowRepository struct {
	mock.Mock
}

// CreateTrigger provides a mock function with given fields: trigger
func (_m *WorkflowRepository) CreateTrigger(trigger model.Trigger) (int64, error) {
	ret := _m.Called(trigger)

	var r0 int64
	if rf, ok := ret.Get(0).(func(model.Trigger) int64); ok {
		r0 = rf(trigger)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Trigger) error); ok {
		r1 = rf(trigger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTriggerByMessageID provides a mock function with given fields: messageID
func (_m *WorkflowRepository) DeleteTriggerByMessageID(messageID int64) error {
	ret := _m.Called(messageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(messageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTriggerByFlag provides a mock function with given fields: t, flag
func (_m *WorkflowRepository) GetTriggerByFlag(t string, flag string) (model.Trigger, error) {
	ret := _m.Called(t, flag)

	var r0 model.Trigger
	if rf, ok := ret.Get(0).(func(string, string) model.Trigger); ok {
		r0 = rf(t, flag)
	} else {
		r0 = ret.Get(0).(model.Trigger)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(t, flag)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTriggersByType provides a mock function with given fields: t
func (_m *WorkflowRepository) ListTriggersByType(t string) ([]model.Trigger, error) {
	ret := _m.Called(t)

	var r0 []model.Trigger
	if rf, ok := ret.Get(0).(func(string) []model.Trigger); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
