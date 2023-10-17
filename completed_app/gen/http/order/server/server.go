// Code generated by goa v3.13.2, DO NOT EDIT.
//
// order HTTP server
//
// Command:
// $ goa gen app/design

package server

import (
	order "app/gen/order"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the order service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	Tea                http.Handler
	GenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the order service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *order.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Tea", "GET", "/tea"},
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		Tea:                NewTeaHandler(e.Tea, mux, decoder, encoder, errhandler, formatter),
		GenHTTPOpenapiJSON: http.FileServer(fileSystemGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "order" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Tea = m(s.Tea)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return order.MethodNames[:] }

// Mount configures the mux to serve the order endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountTeaHandler(mux, h.Tea)
	MountGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/./gen/http/openapi.json", h.GenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the order endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountTeaHandler configures the mux to serve the "order" service "tea"
// endpoint.
func MountTeaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/tea", f)
}

// NewTeaHandler creates a HTTP handler which loads the HTTP request and calls
// the "order" service "tea" endpoint.
func NewTeaHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeTeaRequest(mux, decoder)
		encodeResponse = EncodeTeaResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "tea")
		ctx = context.WithValue(ctx, goa.ServiceKey, "order")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}
