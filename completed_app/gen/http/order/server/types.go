// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order HTTP server types
//
// Command:
// $ goa gen app/design

package server

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

// NewTeaPayload builds a order service tea endpoint payload.
func NewTeaPayload(body *TeaRequestBody) *order.TeaPayload {
	v := &order.TeaPayload{
		IsGreen:      body.IsGreen,
		NumberSugars: body.NumberSugars,
		IncludeMilk:  body.IncludeMilk,
	}

	return v
}
