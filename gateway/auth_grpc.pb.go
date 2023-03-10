// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: auth.proto

package grpc

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	SignUpAdministrator(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignUpUser(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) SignUpAdministrator(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SignUpAdministrator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignUpUser(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SignUpUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/VerifyEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/grpc.Auth/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	SignUpAdministrator(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignUpUser(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) SignUpAdministrator(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpAdministrator not implemented")
}
func (UnimplementedAuthServer) SignUpUser(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUpUser not implemented")
}
func (UnimplementedAuthServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedAuthServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_SignUpAdministrator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUpAdministrator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SignUpAdministrator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUpAdministrator(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignUpUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignUpUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SignUpUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignUpUser(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/VerifyEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Auth/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUpAdministrator",
			Handler:    _Auth_SignUpAdministrator_Handler,
		},
		{
			MethodName: "SignUpUser",
			Handler:    _Auth_SignUpUser_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Auth_SignIn_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _Auth_VerifyEmail_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Auth_ResetPassword_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Auth_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
