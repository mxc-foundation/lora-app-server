// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: settings.proto

/*
Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package api

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_SettingsServiceM2M_GetSettings_0(ctx context.Context, marshaler runtime.Marshaler, client SettingsServiceM2MClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSettingsRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetSettings(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SettingsServiceM2M_GetSettings_0(ctx context.Context, marshaler runtime.Marshaler, server SettingsServiceM2MServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSettingsRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetSettings(ctx, &protoReq)
	return msg, metadata, err

}

func request_SettingsServiceM2M_ModifySettings_0(ctx context.Context, marshaler runtime.Marshaler, client SettingsServiceM2MClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifySettingsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ModifySettings(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SettingsServiceM2M_ModifySettings_0(ctx context.Context, marshaler runtime.Marshaler, server SettingsServiceM2MServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ModifySettingsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ModifySettings(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSettingsServiceM2MHandlerServer registers the http handlers for service SettingsServiceM2M to "mux".
// UnaryRPC     :call SettingsServiceM2MServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterSettingsServiceM2MHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SettingsServiceM2MServer) error {

	mux.Handle("GET", pattern_SettingsServiceM2M_GetSettings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SettingsServiceM2M_GetSettings_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SettingsServiceM2M_GetSettings_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_SettingsServiceM2M_ModifySettings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SettingsServiceM2M_ModifySettings_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SettingsServiceM2M_ModifySettings_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSettingsServiceM2MHandlerFromEndpoint is same as RegisterSettingsServiceM2MHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSettingsServiceM2MHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSettingsServiceM2MHandler(ctx, mux, conn)
}

// RegisterSettingsServiceM2MHandler registers the http handlers for service SettingsServiceM2M to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSettingsServiceM2MHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSettingsServiceM2MHandlerClient(ctx, mux, NewSettingsServiceM2MClient(conn))
}

// RegisterSettingsServiceM2MHandlerClient registers the http handlers for service SettingsServiceM2M
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SettingsServiceM2MClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SettingsServiceM2MClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SettingsServiceM2MClient" to call the correct interceptors.
func RegisterSettingsServiceM2MHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SettingsServiceM2MClient) error {

	mux.Handle("GET", pattern_SettingsServiceM2M_GetSettings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SettingsServiceM2M_GetSettings_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SettingsServiceM2M_GetSettings_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_SettingsServiceM2M_ModifySettings_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SettingsServiceM2M_ModifySettings_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SettingsServiceM2M_ModifySettings_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SettingsServiceM2M_GetSettings_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "settings"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_SettingsServiceM2M_ModifySettings_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"api", "settings"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_SettingsServiceM2M_GetSettings_0 = runtime.ForwardResponseMessage

	forward_SettingsServiceM2M_ModifySettings_0 = runtime.ForwardResponseMessage
)
