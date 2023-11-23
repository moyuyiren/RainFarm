// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: usercenter.proto

package __

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
	UserCenter_Register_FullMethodName = "/rpc.UserCenter/Register"
	UserCenter_Login_FullMethodName    = "/rpc.UserCenter/Login"
	UserCenter_RePwd_FullMethodName    = "/rpc.UserCenter/RePwd"
	UserCenter_Detail_FullMethodName   = "/rpc.UserCenter/Detail"
	UserCenter_GetCode_FullMethodName  = "/rpc.UserCenter/GetCode"
)

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RePwd(ctx context.Context, in *RePwdRequest, opts ...grpc.CallOption) (*RePwdResponse, error)
	Detail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	GetCode(ctx context.Context, in *UserCodeRequest, opts ...grpc.CallOption) (*UserCodeResponse, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserCenter_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserCenter_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) RePwd(ctx context.Context, in *RePwdRequest, opts ...grpc.CallOption) (*RePwdResponse, error) {
	out := new(RePwdResponse)
	err := c.cc.Invoke(ctx, UserCenter_RePwd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) Detail(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, UserCenter_Detail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetCode(ctx context.Context, in *UserCodeRequest, opts ...grpc.CallOption) (*UserCodeResponse, error) {
	out := new(UserCodeResponse)
	err := c.cc.Invoke(ctx, UserCenter_GetCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	RePwd(context.Context, *RePwdRequest) (*RePwdResponse, error)
	Detail(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	GetCode(context.Context, *UserCodeRequest) (*UserCodeResponse, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserCenterServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserCenterServer) RePwd(context.Context, *RePwdRequest) (*RePwdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RePwd not implemented")
}
func (UnimplementedUserCenterServer) Detail(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedUserCenterServer) GetCode(context.Context, *UserCodeRequest) (*UserCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_RePwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RePwdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).RePwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_RePwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).RePwd(ctx, req.(*RePwdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Detail(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_GetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetCode(ctx, req.(*UserCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserCenter_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserCenter_Login_Handler,
		},
		{
			MethodName: "RePwd",
			Handler:    _UserCenter_RePwd_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _UserCenter_Detail_Handler,
		},
		{
			MethodName: "GetCode",
			Handler:    _UserCenter_GetCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}
