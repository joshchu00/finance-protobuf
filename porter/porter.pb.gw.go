// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: porter/porter.proto

/*
Package porter is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package porter

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_PorterV1_GetSymbol_0(ctx context.Context, marshaler runtime.Marshaler, client PorterV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetSymbolRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["exchange"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "exchange")
	}

	protoReq.Exchange, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "exchange", err)
	}

	val, ok = pathParams["symbol"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "symbol")
	}

	protoReq.Symbol, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "symbol", err)
	}

	val, ok = pathParams["period"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "period")
	}

	protoReq.Period, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "period", err)
	}

	msg, err := client.GetSymbol(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterPorterV1HandlerFromEndpoint is same as RegisterPorterV1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPorterV1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterPorterV1Handler(ctx, mux, conn)
}

// RegisterPorterV1Handler registers the http handlers for service PorterV1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPorterV1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPorterV1HandlerClient(ctx, mux, NewPorterV1Client(conn))
}

// RegisterPorterV1HandlerClient registers the http handlers for service PorterV1
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PorterV1Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PorterV1Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PorterV1Client" to call the correct interceptors.
func RegisterPorterV1HandlerClient(ctx context.Context, mux *runtime.ServeMux, client PorterV1Client) error {

	mux.Handle("GET", pattern_PorterV1_GetSymbol_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PorterV1_GetSymbol_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PorterV1_GetSymbol_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PorterV1_GetSymbol_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 4}, []string{"porter", "v1", "exchange", "symbol", "period"}, ""))
)

var (
	forward_PorterV1_GetSymbol_0 = runtime.ForwardResponseMessage
)
