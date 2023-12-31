// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order service
//
// Command:
// $ goa gen app/design

package order

import (
	"context"
)

// A waiter that brings drinks.
type Service interface {
	// Order a cup of tea.
	Tea(context.Context, *TeaPayload) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "order"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"tea"}

// TeaPayload is the payload type of the order service tea method.
type TeaPayload struct {
	// Whether to have green tea instead of normal.
	IsGreen *bool
	// Number of spoons of sugar.
	NumberSugars *int
	// Whether to have milk.
	IncludeMilk *bool
}
