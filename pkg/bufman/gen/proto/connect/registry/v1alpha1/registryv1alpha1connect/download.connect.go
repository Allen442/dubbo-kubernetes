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
// Source: registry/v1alpha1/download.proto

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
	// DownloadServiceName is the fully-qualified name of the DownloadService service.
	DownloadServiceName = "bufman.dubbo.apache.org.registry.v1alpha1.DownloadService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// DownloadServiceDownloadProcedure is the fully-qualified name of the DownloadService's Download
	// RPC.
	DownloadServiceDownloadProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.DownloadService/Download"
	// DownloadServiceDownloadManifestAndBlobsProcedure is the fully-qualified name of the
	// DownloadService's DownloadManifestAndBlobs RPC.
	DownloadServiceDownloadManifestAndBlobsProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.DownloadService/DownloadManifestAndBlobs"
)

// DownloadServiceClient is a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.DownloadService service.
type DownloadServiceClient interface {
	// Download downloads a BSR module.
	// NOTE: Newer clients should use DownloadManifestAndBlobs instead.
	Download(context.Context, *connect_go.Request[v1alpha1.DownloadRequest]) (*connect_go.Response[v1alpha1.DownloadResponse], error)
	// DownloadManifestAndBlobs downloads a module in the manifest+blobs encoding format.
	DownloadManifestAndBlobs(context.Context, *connect_go.Request[v1alpha1.DownloadManifestAndBlobsRequest]) (*connect_go.Response[v1alpha1.DownloadManifestAndBlobsResponse], error)
}

// NewDownloadServiceClient constructs a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.DownloadService service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewDownloadServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) DownloadServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &downloadServiceClient{
		download: connect_go.NewClient[v1alpha1.DownloadRequest, v1alpha1.DownloadResponse](
			httpClient,
			baseURL+DownloadServiceDownloadProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		downloadManifestAndBlobs: connect_go.NewClient[v1alpha1.DownloadManifestAndBlobsRequest, v1alpha1.DownloadManifestAndBlobsResponse](
			httpClient,
			baseURL+DownloadServiceDownloadManifestAndBlobsProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// downloadServiceClient implements DownloadServiceClient.
type downloadServiceClient struct {
	download                 *connect_go.Client[v1alpha1.DownloadRequest, v1alpha1.DownloadResponse]
	downloadManifestAndBlobs *connect_go.Client[v1alpha1.DownloadManifestAndBlobsRequest, v1alpha1.DownloadManifestAndBlobsResponse]
}

// Download calls bufman.dubbo.apache.org.registry.v1alpha1.DownloadService.Download.
func (c *downloadServiceClient) Download(ctx context.Context, req *connect_go.Request[v1alpha1.DownloadRequest]) (*connect_go.Response[v1alpha1.DownloadResponse], error) {
	return c.download.CallUnary(ctx, req)
}

// DownloadManifestAndBlobs calls
// bufman.dubbo.apache.org.registry.v1alpha1.DownloadService.DownloadManifestAndBlobs.
func (c *downloadServiceClient) DownloadManifestAndBlobs(ctx context.Context, req *connect_go.Request[v1alpha1.DownloadManifestAndBlobsRequest]) (*connect_go.Response[v1alpha1.DownloadManifestAndBlobsResponse], error) {
	return c.downloadManifestAndBlobs.CallUnary(ctx, req)
}

// DownloadServiceHandler is an implementation of the
// bufman.dubbo.apache.org.registry.v1alpha1.DownloadService service.
type DownloadServiceHandler interface {
	// Download downloads a BSR module.
	// NOTE: Newer clients should use DownloadManifestAndBlobs instead.
	Download(context.Context, *connect_go.Request[v1alpha1.DownloadRequest]) (*connect_go.Response[v1alpha1.DownloadResponse], error)
	// DownloadManifestAndBlobs downloads a module in the manifest+blobs encoding format.
	DownloadManifestAndBlobs(context.Context, *connect_go.Request[v1alpha1.DownloadManifestAndBlobsRequest]) (*connect_go.Response[v1alpha1.DownloadManifestAndBlobsResponse], error)
}

// NewDownloadServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewDownloadServiceHandler(svc DownloadServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	downloadServiceDownloadHandler := connect_go.NewUnaryHandler(
		DownloadServiceDownloadProcedure,
		svc.Download,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	downloadServiceDownloadManifestAndBlobsHandler := connect_go.NewUnaryHandler(
		DownloadServiceDownloadManifestAndBlobsProcedure,
		svc.DownloadManifestAndBlobs,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/bufman.dubbo.apache.org.registry.v1alpha1.DownloadService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case DownloadServiceDownloadProcedure:
			downloadServiceDownloadHandler.ServeHTTP(w, r)
		case DownloadServiceDownloadManifestAndBlobsProcedure:
			downloadServiceDownloadManifestAndBlobsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedDownloadServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedDownloadServiceHandler struct{}

func (UnimplementedDownloadServiceHandler) Download(context.Context, *connect_go.Request[v1alpha1.DownloadRequest]) (*connect_go.Response[v1alpha1.DownloadResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.DownloadService.Download is not implemented"))
}

func (UnimplementedDownloadServiceHandler) DownloadManifestAndBlobs(context.Context, *connect_go.Request[v1alpha1.DownloadManifestAndBlobsRequest]) (*connect_go.Response[v1alpha1.DownloadManifestAndBlobsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.DownloadService.DownloadManifestAndBlobs is not implemented"))
}
