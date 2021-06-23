// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/tsundata/assistant/api/pb"
)

// MessageClient is an autogenerated mock type for the MessageClient type
type MessageClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) Create(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.TextsReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.TextsReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.TextsReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.TextsReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateActionMessage provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) CreateActionMessage(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.StateReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) *pb.StateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.StateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) Delete(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.TextReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.TextReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.TextReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteWorkflowMessage provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) DeleteWorkflowMessage(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.StateReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.StateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.StateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) Get(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.MessageReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.MessageReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.MessageReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.MessageReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActionMessages provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) GetActionMessages(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.ActionReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.ActionReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) *pb.ActionReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ActionReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) List(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.MessageListReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.MessageListReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.MessageListReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.MessageListReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Run provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) Run(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.TextReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.TextReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.TextReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: ctx, in, opts
func (_m *MessageClient) Send(ctx context.Context, in *pb.MessageRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.StateReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) *pb.StateReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.StateReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.MessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
