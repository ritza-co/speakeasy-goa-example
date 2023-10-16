// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order client HTTP transport
//
// Command:
// $ goa gen app/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the order service endpoint HTTP clients.
type Client struct {
	// Tea Doer is the HTTP client used to make requests to the tea endpoint.
	TeaDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the order service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		TeaDoer:             doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Tea returns an endpoint that makes HTTP requests to the order service tea
// server.
func (c *Client) Tea() goa.Endpoint {
	var (
		encodeRequest  = EncodeTeaRequest(c.encoder)
		decodeResponse = DecodeTeaResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildTeaRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TeaDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("order", "tea", err)
		}
		return decodeResponse(resp)
	}
}
