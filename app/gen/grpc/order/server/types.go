// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order gRPC server types
//
// Command:
// $ goa gen app/design

package server

import (
	orderpb "app/gen/grpc/order/pb"
	order "app/gen/order"
)

// NewTeaPayload builds the payload of the "tea" endpoint of the "order"
// service from the gRPC request type.
func NewTeaPayload(message *orderpb.TeaRequest) *order.TeaPayload {
	v := &order.TeaPayload{
		IsGreen:     message.IsGreen,
		IncludeMilk: message.IncludeMilk,
	}
	if message.NumberSugars != nil {
		numberSugars := int(*message.NumberSugars)
		v.NumberSugars = &numberSugars
	}
	return v
}

// NewProtoTeaResponse builds the gRPC response type from the result of the
// "tea" endpoint of the "order" service.
func NewProtoTeaResponse(result string) *orderpb.TeaResponse {
	message := &orderpb.TeaResponse{}
	message.Field = result
	return message
}