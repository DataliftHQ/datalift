// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: audit/v1/audit.proto

package auditv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "go.datalift.io/datalift/api/audit/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AuditAPIName is the fully-qualified name of the AuditAPI service.
	AuditAPIName = "datalift.audit.v1.AuditAPI"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuditAPIGetEventsProcedure is the fully-qualified name of the AuditAPI's GetEvents RPC.
	AuditAPIGetEventsProcedure = "/datalift.audit.v1.AuditAPI/GetEvents"
	// AuditAPIGetEventProcedure is the fully-qualified name of the AuditAPI's GetEvent RPC.
	AuditAPIGetEventProcedure = "/datalift.audit.v1.AuditAPI/GetEvent"
)

// AuditAPIClient is a client for the datalift.audit.v1.AuditAPI service.
type AuditAPIClient interface {
	GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.GetEventsResponse], error)
	GetEvent(context.Context, *connect_go.Request[v1.GetEventRequest]) (*connect_go.Response[v1.GetEventResponse], error)
}

// NewAuditAPIClient constructs a client for the datalift.audit.v1.AuditAPI service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuditAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuditAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &auditAPIClient{
		getEvents: connect_go.NewClient[v1.GetEventsRequest, v1.GetEventsResponse](
			httpClient,
			baseURL+AuditAPIGetEventsProcedure,
			opts...,
		),
		getEvent: connect_go.NewClient[v1.GetEventRequest, v1.GetEventResponse](
			httpClient,
			baseURL+AuditAPIGetEventProcedure,
			opts...,
		),
	}
}

// auditAPIClient implements AuditAPIClient.
type auditAPIClient struct {
	getEvents *connect_go.Client[v1.GetEventsRequest, v1.GetEventsResponse]
	getEvent  *connect_go.Client[v1.GetEventRequest, v1.GetEventResponse]
}

// GetEvents calls datalift.audit.v1.AuditAPI.GetEvents.
func (c *auditAPIClient) GetEvents(ctx context.Context, req *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.GetEventsResponse], error) {
	return c.getEvents.CallUnary(ctx, req)
}

// GetEvent calls datalift.audit.v1.AuditAPI.GetEvent.
func (c *auditAPIClient) GetEvent(ctx context.Context, req *connect_go.Request[v1.GetEventRequest]) (*connect_go.Response[v1.GetEventResponse], error) {
	return c.getEvent.CallUnary(ctx, req)
}

// AuditAPIHandler is an implementation of the datalift.audit.v1.AuditAPI service.
type AuditAPIHandler interface {
	GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.GetEventsResponse], error)
	GetEvent(context.Context, *connect_go.Request[v1.GetEventRequest]) (*connect_go.Response[v1.GetEventResponse], error)
}

// NewAuditAPIHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuditAPIHandler(svc AuditAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	auditAPIGetEventsHandler := connect_go.NewUnaryHandler(
		AuditAPIGetEventsProcedure,
		svc.GetEvents,
		opts...,
	)
	auditAPIGetEventHandler := connect_go.NewUnaryHandler(
		AuditAPIGetEventProcedure,
		svc.GetEvent,
		opts...,
	)
	return "/datalift.audit.v1.AuditAPI/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuditAPIGetEventsProcedure:
			auditAPIGetEventsHandler.ServeHTTP(w, r)
		case AuditAPIGetEventProcedure:
			auditAPIGetEventHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuditAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedAuditAPIHandler struct{}

func (UnimplementedAuditAPIHandler) GetEvents(context.Context, *connect_go.Request[v1.GetEventsRequest]) (*connect_go.Response[v1.GetEventsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("datalift.audit.v1.AuditAPI.GetEvents is not implemented"))
}

func (UnimplementedAuditAPIHandler) GetEvent(context.Context, *connect_go.Request[v1.GetEventRequest]) (*connect_go.Response[v1.GetEventResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("datalift.audit.v1.AuditAPI.GetEvent is not implemented"))
}
