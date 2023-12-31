// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order gRPC client
//
// Command:
// $ goa gen app/design

package client

import (
	orderpb "app/gen/grpc/order/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli orderpb.OrderClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the order service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: orderpb.NewOrderClient(cc),
		opts:    opts,
	}
}

// Tea calls the "Tea" function in orderpb.OrderClient interface.
func (c *Client) Tea() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildTeaFunc(c.grpccli, c.opts...),
			EncodeTeaRequest,
			DecodeTeaResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
