// Code generated by goa v3.13.2, DO NOT EDIT.
//
// band gRPC client types
//
// Command:
// $ goa gen app/design

package client

import (
	band "app/gen/band"
	bandpb "app/gen/grpc/band/pb"
)

// NewProtoPlayRequest builds the gRPC request type from the payload of the
// "play" endpoint of the "band" service.
func NewProtoPlayRequest(payload *band.PlayPayload) *bandpb.PlayRequest {
	message := &bandpb.PlayRequest{
		Style: payload.Style,
	}
	return message
}