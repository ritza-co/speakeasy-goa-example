package main

import (
	band "app/gen/band"
	bandpb "app/gen/grpc/band/pb"
	bandsvr "app/gen/grpc/band/server"
	orderpb "app/gen/grpc/order/pb"
	ordersvr "app/gen/grpc/order/server"
	order "app/gen/order"
	"context"
	"log"
	"net"
	"net/url"
	"sync"

	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, orderEndpoints *order.Endpoints, bandEndpoints *band.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		orderServer *ordersvr.Server
		bandServer  *bandsvr.Server
	)
	{
		orderServer = ordersvr.New(orderEndpoints, nil)
		bandServer = bandsvr.New(bandEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	orderpb.RegisterOrderServer(srv, orderServer)
	bandpb.RegisterBandServer(srv, bandServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
