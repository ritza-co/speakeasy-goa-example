// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
)

// NewMultiplyPayload builds a calc service multiply endpoint payload.
func NewMultiplyPayload(a int, b int) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{}
	v.A = a
	v.B = b

	return v
}