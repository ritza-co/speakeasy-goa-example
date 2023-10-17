// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Multiply goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Multiply: NewMultiplyEndpoint(s),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Multiply = m(e.Multiply)
}

// NewMultiplyEndpoint returns an endpoint function that calls the method
// "multiply" of service "calc".
func NewMultiplyEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MultiplyPayload)
		return s.Multiply(ctx, p)
	}
}