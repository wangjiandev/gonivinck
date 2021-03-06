// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package product

import (
	"context"

	"mall/service/product/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRequest  = pb.CreateRequest
	CreateResponse = pb.CreateResponse
	DetailRequest  = pb.DetailRequest
	DetailResponse = pb.DetailResponse
	RemoveRequest  = pb.RemoveRequest
	RemoveResponse = pb.RemoveResponse
	UpdateRequest  = pb.UpdateRequest
	UpdateResponse = pb.UpdateResponse

	Product interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultProduct) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultProduct) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultProduct) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}
