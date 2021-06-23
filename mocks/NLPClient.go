// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/tsundata/assistant/api/pb"
)

// NLPClient is an autogenerated mock type for the NLPClient type
type NLPClient struct {
	mock.Mock
}

// Pinyin provides a mock function with given fields: ctx, in, opts
func (_m *NLPClient) Pinyin(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.WordsReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.WordsReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) *pb.WordsReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.WordsReply)
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

// Segmentation provides a mock function with given fields: ctx, in, opts
func (_m *NLPClient) Segmentation(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.WordsReply, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.WordsReply
	if rf, ok := ret.Get(0).(func(context.Context, *pb.TextRequest, ...grpc.CallOption) *pb.WordsReply); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.WordsReply)
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
