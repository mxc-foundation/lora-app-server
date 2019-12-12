// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: super_node.proto

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

func request_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(ctx context.Context, marshaler runtime.Marshaler, client SuperNodeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSuperNodeActiveMoneyAccountRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	val, ok = pathParams["money_abbr"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "money_abbr")
	}

	e, err = runtime.Enum(val, Money_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "money_abbr", err)
	}

	protoReq.MoneyAbbr = Money(e)

	msg, err := client.GetSuperNodeActiveMoneyAccount(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(ctx context.Context, marshaler runtime.Marshaler, server SuperNodeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSuperNodeActiveMoneyAccountRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	val, ok = pathParams["money_abbr"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "money_abbr")
	}

	e, err = runtime.Enum(val, Money_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "money_abbr", err)
	}

	protoReq.MoneyAbbr = Money(e)

	msg, err := server.GetSuperNodeActiveMoneyAccount(ctx, &protoReq)
	return msg, metadata, err

}

func request_SuperNodeService_AddSuperNodeMoneyAccount_0(ctx context.Context, marshaler runtime.Marshaler, client SuperNodeServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddSuperNodeMoneyAccountRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	val, ok = pathParams["money_abbr"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "money_abbr")
	}

	e, err = runtime.Enum(val, Money_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "money_abbr", err)
	}

	protoReq.MoneyAbbr = Money(e)

	msg, err := client.AddSuperNodeMoneyAccount(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SuperNodeService_AddSuperNodeMoneyAccount_0(ctx context.Context, marshaler runtime.Marshaler, server SuperNodeServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AddSuperNodeMoneyAccountRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		e   int32
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["org_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "org_id")
	}

	protoReq.OrgId, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "org_id", err)
	}

	val, ok = pathParams["money_abbr"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "money_abbr")
	}

	e, err = runtime.Enum(val, Money_value)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "money_abbr", err)
	}

	protoReq.MoneyAbbr = Money(e)

	msg, err := server.AddSuperNodeMoneyAccount(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSuperNodeServiceHandlerServer registers the http handlers for service SuperNodeService to "mux".
// UnaryRPC     :call SuperNodeServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterSuperNodeServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SuperNodeServiceServer) error {

	mux.Handle("GET", pattern_SuperNodeService_GetSuperNodeActiveMoneyAccount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SuperNodeService_AddSuperNodeMoneyAccount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SuperNodeService_AddSuperNodeMoneyAccount_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SuperNodeService_AddSuperNodeMoneyAccount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSuperNodeServiceHandlerFromEndpoint is same as RegisterSuperNodeServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSuperNodeServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterSuperNodeServiceHandler(ctx, mux, conn)
}

// RegisterSuperNodeServiceHandler registers the http handlers for service SuperNodeService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSuperNodeServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSuperNodeServiceHandlerClient(ctx, mux, NewSuperNodeServiceClient(conn))
}

// RegisterSuperNodeServiceHandlerClient registers the http handlers for service SuperNodeService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SuperNodeServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SuperNodeServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SuperNodeServiceClient" to call the correct interceptors.
func RegisterSuperNodeServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SuperNodeServiceClient) error {

	mux.Handle("GET", pattern_SuperNodeService_GetSuperNodeActiveMoneyAccount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SuperNodeService_GetSuperNodeActiveMoneyAccount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_SuperNodeService_AddSuperNodeMoneyAccount_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SuperNodeService_AddSuperNodeMoneyAccount_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SuperNodeService_AddSuperNodeMoneyAccount_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SuperNodeService_GetSuperNodeActiveMoneyAccount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "supernode", "org_id", "active-account", "money_abbr"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_SuperNodeService_AddSuperNodeMoneyAccount_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "supernode", "org_id", "add-account", "money_abbr"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_SuperNodeService_GetSuperNodeActiveMoneyAccount_0 = runtime.ForwardResponseMessage

	forward_SuperNodeService_AddSuperNodeMoneyAccount_0 = runtime.ForwardResponseMessage
)
