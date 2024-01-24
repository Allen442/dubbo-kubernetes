// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: registry/v1alpha1/studio_request.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// StudioRequestServiceName is the fully-qualified name of the StudioRequestService service.
	StudioRequestServiceName = "bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StudioRequestServiceCreateStudioRequestProcedure is the fully-qualified name of the
	// StudioRequestService's CreateStudioRequest RPC.
	StudioRequestServiceCreateStudioRequestProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService/CreateStudioRequest"
	// StudioRequestServiceRenameStudioRequestProcedure is the fully-qualified name of the
	// StudioRequestService's RenameStudioRequest RPC.
	StudioRequestServiceRenameStudioRequestProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService/RenameStudioRequest"
	// StudioRequestServiceDeleteStudioRequestProcedure is the fully-qualified name of the
	// StudioRequestService's DeleteStudioRequest RPC.
	StudioRequestServiceDeleteStudioRequestProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService/DeleteStudioRequest"
	// StudioRequestServiceListStudioRequestsProcedure is the fully-qualified name of the
	// StudioRequestService's ListStudioRequests RPC.
	StudioRequestServiceListStudioRequestsProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService/ListStudioRequests"
)

// StudioRequestServiceClient is a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService service.
type StudioRequestServiceClient interface {
	// CreateStudioRequest registers a favorite Studio Requests to the caller's
	// BSR profile.
	CreateStudioRequest(context.Context, *connect_go.Request[v1alpha1.CreateStudioRequestRequest]) (*connect_go.Response[v1alpha1.CreateStudioRequestResponse], error)
	// RenameStudioRequest renames an existing Studio Request.
	RenameStudioRequest(context.Context, *connect_go.Request[v1alpha1.RenameStudioRequestRequest]) (*connect_go.Response[v1alpha1.RenameStudioRequestResponse], error)
	// DeleteStudioRequest removes a favorite Studio Request from the caller's BSR
	// profile.
	DeleteStudioRequest(context.Context, *connect_go.Request[v1alpha1.DeleteStudioRequestRequest]) (*connect_go.Response[v1alpha1.DeleteStudioRequestResponse], error)
	// ListStudioRequests shows the caller's favorited Studio Requests.
	ListStudioRequests(context.Context, *connect_go.Request[v1alpha1.ListStudioRequestsRequest]) (*connect_go.Response[v1alpha1.ListStudioRequestsResponse], error)
}

// NewStudioRequestServiceClient constructs a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStudioRequestServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StudioRequestServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &studioRequestServiceClient{
		createStudioRequest: connect_go.NewClient[v1alpha1.CreateStudioRequestRequest, v1alpha1.CreateStudioRequestResponse](
			httpClient,
			baseURL+StudioRequestServiceCreateStudioRequestProcedure,
			opts...,
		),
		renameStudioRequest: connect_go.NewClient[v1alpha1.RenameStudioRequestRequest, v1alpha1.RenameStudioRequestResponse](
			httpClient,
			baseURL+StudioRequestServiceRenameStudioRequestProcedure,
			opts...,
		),
		deleteStudioRequest: connect_go.NewClient[v1alpha1.DeleteStudioRequestRequest, v1alpha1.DeleteStudioRequestResponse](
			httpClient,
			baseURL+StudioRequestServiceDeleteStudioRequestProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
			connect_go.WithClientOptions(opts...),
		),
		listStudioRequests: connect_go.NewClient[v1alpha1.ListStudioRequestsRequest, v1alpha1.ListStudioRequestsResponse](
			httpClient,
			baseURL+StudioRequestServiceListStudioRequestsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// studioRequestServiceClient implements StudioRequestServiceClient.
type studioRequestServiceClient struct {
	createStudioRequest *connect_go.Client[v1alpha1.CreateStudioRequestRequest, v1alpha1.CreateStudioRequestResponse]
	renameStudioRequest *connect_go.Client[v1alpha1.RenameStudioRequestRequest, v1alpha1.RenameStudioRequestResponse]
	deleteStudioRequest *connect_go.Client[v1alpha1.DeleteStudioRequestRequest, v1alpha1.DeleteStudioRequestResponse]
	listStudioRequests  *connect_go.Client[v1alpha1.ListStudioRequestsRequest, v1alpha1.ListStudioRequestsResponse]
}

// CreateStudioRequest calls
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.CreateStudioRequest.
func (c *studioRequestServiceClient) CreateStudioRequest(ctx context.Context, req *connect_go.Request[v1alpha1.CreateStudioRequestRequest]) (*connect_go.Response[v1alpha1.CreateStudioRequestResponse], error) {
	return c.createStudioRequest.CallUnary(ctx, req)
}

// RenameStudioRequest calls
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.RenameStudioRequest.
func (c *studioRequestServiceClient) RenameStudioRequest(ctx context.Context, req *connect_go.Request[v1alpha1.RenameStudioRequestRequest]) (*connect_go.Response[v1alpha1.RenameStudioRequestResponse], error) {
	return c.renameStudioRequest.CallUnary(ctx, req)
}

// DeleteStudioRequest calls
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.DeleteStudioRequest.
func (c *studioRequestServiceClient) DeleteStudioRequest(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteStudioRequestRequest]) (*connect_go.Response[v1alpha1.DeleteStudioRequestResponse], error) {
	return c.deleteStudioRequest.CallUnary(ctx, req)
}

// ListStudioRequests calls
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.ListStudioRequests.
func (c *studioRequestServiceClient) ListStudioRequests(ctx context.Context, req *connect_go.Request[v1alpha1.ListStudioRequestsRequest]) (*connect_go.Response[v1alpha1.ListStudioRequestsResponse], error) {
	return c.listStudioRequests.CallUnary(ctx, req)
}

// StudioRequestServiceHandler is an implementation of the
// bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService service.
type StudioRequestServiceHandler interface {
	// CreateStudioRequest registers a favorite Studio Requests to the caller's
	// BSR profile.
	CreateStudioRequest(context.Context, *connect_go.Request[v1alpha1.CreateStudioRequestRequest]) (*connect_go.Response[v1alpha1.CreateStudioRequestResponse], error)
	// RenameStudioRequest renames an existing Studio Request.
	RenameStudioRequest(context.Context, *connect_go.Request[v1alpha1.RenameStudioRequestRequest]) (*connect_go.Response[v1alpha1.RenameStudioRequestResponse], error)
	// DeleteStudioRequest removes a favorite Studio Request from the caller's BSR
	// profile.
	DeleteStudioRequest(context.Context, *connect_go.Request[v1alpha1.DeleteStudioRequestRequest]) (*connect_go.Response[v1alpha1.DeleteStudioRequestResponse], error)
	// ListStudioRequests shows the caller's favorited Studio Requests.
	ListStudioRequests(context.Context, *connect_go.Request[v1alpha1.ListStudioRequestsRequest]) (*connect_go.Response[v1alpha1.ListStudioRequestsResponse], error)
}

// NewStudioRequestServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStudioRequestServiceHandler(svc StudioRequestServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	studioRequestServiceCreateStudioRequestHandler := connect_go.NewUnaryHandler(
		StudioRequestServiceCreateStudioRequestProcedure,
		svc.CreateStudioRequest,
		opts...,
	)
	studioRequestServiceRenameStudioRequestHandler := connect_go.NewUnaryHandler(
		StudioRequestServiceRenameStudioRequestProcedure,
		svc.RenameStudioRequest,
		opts...,
	)
	studioRequestServiceDeleteStudioRequestHandler := connect_go.NewUnaryHandler(
		StudioRequestServiceDeleteStudioRequestProcedure,
		svc.DeleteStudioRequest,
		connect_go.WithIdempotency(connect_go.IdempotencyIdempotent),
		connect_go.WithHandlerOptions(opts...),
	)
	studioRequestServiceListStudioRequestsHandler := connect_go.NewUnaryHandler(
		StudioRequestServiceListStudioRequestsProcedure,
		svc.ListStudioRequests,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StudioRequestServiceCreateStudioRequestProcedure:
			studioRequestServiceCreateStudioRequestHandler.ServeHTTP(w, r)
		case StudioRequestServiceRenameStudioRequestProcedure:
			studioRequestServiceRenameStudioRequestHandler.ServeHTTP(w, r)
		case StudioRequestServiceDeleteStudioRequestProcedure:
			studioRequestServiceDeleteStudioRequestHandler.ServeHTTP(w, r)
		case StudioRequestServiceListStudioRequestsProcedure:
			studioRequestServiceListStudioRequestsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStudioRequestServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStudioRequestServiceHandler struct{}

func (UnimplementedStudioRequestServiceHandler) CreateStudioRequest(context.Context, *connect_go.Request[v1alpha1.CreateStudioRequestRequest]) (*connect_go.Response[v1alpha1.CreateStudioRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.CreateStudioRequest is not implemented"))
}

func (UnimplementedStudioRequestServiceHandler) RenameStudioRequest(context.Context, *connect_go.Request[v1alpha1.RenameStudioRequestRequest]) (*connect_go.Response[v1alpha1.RenameStudioRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.RenameStudioRequest is not implemented"))
}

func (UnimplementedStudioRequestServiceHandler) DeleteStudioRequest(context.Context, *connect_go.Request[v1alpha1.DeleteStudioRequestRequest]) (*connect_go.Response[v1alpha1.DeleteStudioRequestResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.DeleteStudioRequest is not implemented"))
}

func (UnimplementedStudioRequestServiceHandler) ListStudioRequests(context.Context, *connect_go.Request[v1alpha1.ListStudioRequestsRequest]) (*connect_go.Response[v1alpha1.ListStudioRequestsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.StudioRequestService.ListStudioRequests is not implemented"))
}
