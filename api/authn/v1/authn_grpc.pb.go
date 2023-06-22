// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: authn/v1/authn.proto

package authnv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthnAPI_Login_FullMethodName       = "/datalift.authn.v1.AuthnAPI/Login"
	AuthnAPI_Callback_FullMethodName    = "/datalift.authn.v1.AuthnAPI/Callback"
	AuthnAPI_CreateToken_FullMethodName = "/datalift.authn.v1.AuthnAPI/CreateToken"
)

// AuthnAPIClient is the client API for AuthnAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthnAPIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
}

type authnAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnAPIClient(cc grpc.ClientConnInterface) AuthnAPIClient {
	return &authnAPIClient{cc}
}

func (c *authnAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthnAPI_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnAPIClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, AuthnAPI_Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnAPIClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, AuthnAPI_CreateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnAPIServer is the server API for AuthnAPI service.
// All implementations should embed UnimplementedAuthnAPIServer
// for forward compatibility
type AuthnAPIServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
}

// UnimplementedAuthnAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAuthnAPIServer struct {
}

func (UnimplementedAuthnAPIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthnAPIServer) Callback(context.Context, *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (UnimplementedAuthnAPIServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}

// UnsafeAuthnAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthnAPIServer will
// result in compilation errors.
type UnsafeAuthnAPIServer interface {
	mustEmbedUnimplementedAuthnAPIServer()
}

func RegisterAuthnAPIServer(s grpc.ServiceRegistrar, srv AuthnAPIServer) {
	s.RegisterService(&AuthnAPI_ServiceDesc, srv)
}

func _AuthnAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnAPI_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnAPI_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnAPI_Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnAPI_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthnAPI_CreateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthnAPI_ServiceDesc is the grpc.ServiceDesc for AuthnAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthnAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datalift.authn.v1.AuthnAPI",
	HandlerType: (*AuthnAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthnAPI_Login_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _AuthnAPI_Callback_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _AuthnAPI_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authn/v1/authn.proto",
}
