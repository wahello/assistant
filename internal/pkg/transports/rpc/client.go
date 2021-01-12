package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"time"
)

type ClientOptions struct {
	Wait            time.Duration
	Tag             string
	Etcd            string
	GrpcDialOptions []grpc.DialOption
}

func NewClientOptions(v *viper.Viper, tracer opentracing.Tracer) (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
	}

	o.GrpcDialOptions = append(o.GrpcDialOptions,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				otgrpc.OpenTracingClientInterceptor(tracer),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				otgrpc.OpenTracingStreamClientInterceptor(tracer),
			),
		),
	)

	return o, err
}

type ClientOptional func(o *ClientOptions)

func WithTimeout(d time.Duration) ClientOptional {
	return func(o *ClientOptions) {
		o.Wait = d
	}
}

func WithTag(tag string) ClientOptional {
	return func(o *ClientOptions) {
		o.Tag = tag
	}
}

type Client struct {
	o  *ClientOptions
	CC *grpc.ClientConn
}

func NewClient(o *ClientOptions) (*Client, error) {
	return &Client{
		o: o,
	}, nil
}

func (c *Client) Dial(service string, options ...ClientOptional) (*grpc.ClientConn, error) {
	o := &ClientOptions{
		Wait:            c.o.Wait,
		Tag:             c.o.Tag,
		Etcd:            c.o.Etcd,
		GrpcDialOptions: c.o.GrpcDialOptions,
	}
	for _, option := range options {
		option(o)
	}

	etcdCli, err := clientv3.NewFromURL(o.Etcd)
	if err != nil {
		return nil, err
	}
	re := &etcdnaming.GRPCResolver{Client: etcdCli}
	rr := grpc.RoundRobin(re)                                     // nolint
	gdOptions := append(o.GrpcDialOptions, grpc.WithBalancer(rr)) // nolint

	conn, err := grpc.Dial(service, gdOptions...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
