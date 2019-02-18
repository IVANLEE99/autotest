// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: stages.proto

/*
Package stages is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package stages

import (
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"github.com/ory/hydra/sdk/go/hydra"
	"github.com/ory/hydra/sdk/go/hydra/swagger"
	"github.com/yourhe/grpc-gateway/policy"
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
var HydraClient hydra.SDK

func request_StagesSvic_CreateStages_0(ctx context.Context, marshaler runtime.Marshaler, client StagesSvicClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateStagesRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateStages(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_StagesSvic_GetStages_0 = &utilities.DoubleArray{Encoding: map[string]int{"id": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_StagesSvic_GetStages_0(ctx context.Context, marshaler runtime.Marshaler, client StagesSvicClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SearchStagesRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_StagesSvic_GetStages_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetStages(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_StagesSvic_ListStages_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_StagesSvic_ListStages_0(ctx context.Context, marshaler runtime.Marshaler, client StagesSvicClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SearchStagesRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_StagesSvic_ListStages_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ListStages(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_StagesSvic_UpdateStages_0(ctx context.Context, marshaler runtime.Marshaler, client StagesSvicClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UpdateStagesRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.UpdateStages(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_StagesSvic_RemoveStages_0(ctx context.Context, marshaler runtime.Marshaler, client StagesSvicClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateStagesRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.RemoveStages(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	headerAuthorize = "authorization"
)

// func exampleAuthFunc(rule *policy.PolicyRule, w http.ResponseWriter, req *http.Request, pathParams map[string]string) (context.Context, error) {
func exampleAuthFunc(rule *policy.PolicyRule, fn func(w http.ResponseWriter, req *http.Request, pathParams map[string]string)) func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
	return func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		fn(w, req, pathParams)
		return
		val := req.Header.Get(headerAuthorize)
		var expectedScheme = "bearer"
		if val == "" {
			// return "", grpc.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
			// runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			runtime.DefaultOtherErrorHandler(w, req, "Request unauthenticated with v:"+val, http.StatusUnauthorized)
			return

		}
		splits := strings.SplitN(val, " ", 2)
		if len(splits) < 2 {
			runtime.DefaultOtherErrorHandler(w, req, "Bad authorization string", http.StatusUnauthorized)

			// return "", grpc.Errorf(codes.Unauthenticated, "Bad authorization string")
			return
		}
		if strings.ToLower(splits[0]) != strings.ToLower(expectedScheme) {
			runtime.DefaultOtherErrorHandler(w, req, "Request unauthenticated with ", http.StatusUnauthorized)
			return
			// return "", grpc.Errorf(codes.Unauthenticated, "Request unauthenticated with "+expectedScheme)
		}
		token := splits[1]
		Accesstoken, _, err := HydraClient.DoesWardenAllowTokenAccessRequest(swagger.WardenTokenAccessRequest{
			Action:   rule.Action,
			Resource: rule.Resources,
			Token:    token,
		})
		if err != nil {
			runtime.DefaultOtherErrorHandler(w, req, err.Error(), http.StatusUnauthorized)
			return
		}
		if Accesstoken.Allowed != true {
			runtime.DefaultOtherErrorHandler(w, req, "Request unauthenticated with not Allowed", http.StatusUnauthorized)
			return
		}
		fn(w, req, pathParams)
	}
}

// RegisterStagesSvicHandlerFromEndpoint is same as RegisterStagesSvicHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterStagesSvicHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterStagesSvicHandler(ctx, mux, conn)
}

// RegisterStagesSvicHandler registers the http handlers for service StagesSvic to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterStagesSvicHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterStagesSvicHandlerClient(ctx, mux, NewStagesSvicClient(conn))
}

// RegisterStagesSvicHandler registers the http handlers for service StagesSvic to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "StagesSvicClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "StagesSvicClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "StagesSvicClient" to call the correct interceptors.
func RegisterStagesSvicHandlerClient(ctx context.Context, mux *runtime.ServeMux, client StagesSvicClient) error {

	mux.Handle("POST", pattern_StagesSvic_CreateStages_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {

		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}

		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StagesSvic_CreateStages_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StagesSvic_CreateStages_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_StagesSvic_GetStages_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {

		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}

		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StagesSvic_GetStages_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StagesSvic_GetStages_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_StagesSvic_ListStages_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {

		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}

		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StagesSvic_ListStages_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StagesSvic_ListStages_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_StagesSvic_UpdateStages_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {

		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}

		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StagesSvic_UpdateStages_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StagesSvic_UpdateStages_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_StagesSvic_RemoveStages_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {

		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}

		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_StagesSvic_RemoveStages_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_StagesSvic_RemoveStages_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_StagesSvic_CreateStages_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "test", "stages"}, ""))

	pattern_StagesSvic_GetStages_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "test", "stages", "id"}, ""))

	pattern_StagesSvic_ListStages_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "test", "stages"}, ""))

	pattern_StagesSvic_UpdateStages_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "test", "stages", "id"}, ""))

	pattern_StagesSvic_RemoveStages_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "test", "stages"}, ""))
)

var (
	forward_StagesSvic_CreateStages_0 = runtime.ForwardResponseMessage

	forward_StagesSvic_GetStages_0 = runtime.ForwardResponseMessage

	forward_StagesSvic_ListStages_0 = runtime.ForwardResponseMessage

	forward_StagesSvic_UpdateStages_0 = runtime.ForwardResponseMessage

	forward_StagesSvic_RemoveStages_0 = runtime.ForwardResponseMessage
)

// type RequestPolicyMap map[string]*policy.PolicyRule
type RequestPolicyMap struct {
	PolicyMap   map[string]*policy.PolicyRule
	PatternList []ServiceMothedMap
}
type ServiceMothedMap struct {
	Name    string
	Pattern *runtime.Pattern
}

var (
	RPM = RequestPolicyMap{
		PolicyMap: map[string]*policy.PolicyRule{},
		// PatternList => []ServiceMothedMap
		PatternList: []ServiceMothedMap{

			{
				Name:    "/v1/test/stages",
				Pattern: &pattern_StagesSvic_CreateStages_0,
			},

			{
				Name:    "/v1/test/stages/{id}",
				Pattern: &pattern_StagesSvic_GetStages_0,
			},

			{
				Name:    "/v1/test/stages",
				Pattern: &pattern_StagesSvic_ListStages_0,
			},

			{
				Name:    "/v1/test/stages/{id}",
				Pattern: &pattern_StagesSvic_UpdateStages_0,
			},

			{
				Name:    "/v1/test/stages",
				Pattern: &pattern_StagesSvic_RemoveStages_0,
			},
		},
	}
)