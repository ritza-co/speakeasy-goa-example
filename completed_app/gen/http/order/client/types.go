// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order HTTP client types
//
// Command:
// $ goa gen app/design

package client

import (
	order "app/gen/order"
)

// TeaRequestBody is the type of the "order" service "tea" endpoint HTTP
// request body.
type TeaRequestBody struct {
	// Whether to have green tea instead of normal.
	IsGreen *bool `form:"isGreen,omitempty" json:"isGreen,omitempty" xml:"isGreen,omitempty"`
	// Number of spoons of sugar.
	NumberSugars *int `form:"numberSugars,omitempty" json:"numberSugars,omitempty" xml:"numberSugars,omitempty"`
	// Whether to have milk.
	IncludeMilk *bool `form:"includeMilk,omitempty" json:"includeMilk,omitempty" xml:"includeMilk,omitempty"`
}

// NewTeaRequestBody builds the HTTP request body from the payload of the "tea"
// endpoint of the "order" service.
func NewTeaRequestBody(p *order.TeaPayload) *TeaRequestBody {
	body := &TeaRequestBody{
		IsGreen:      p.IsGreen,
		NumberSugars: p.NumberSugars,
		IncludeMilk:  p.IncludeMilk,
	}
	return body
}
