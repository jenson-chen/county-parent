// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: county.proto

package service

import (
	context "context"
	"github.com/jenson/county/core/models"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CountyRpc_Insert_FullMethodName              = "/countyService.countyRpc/Insert"
	CountyRpc_QueryDetailByEditor_FullMethodName = "/countyService.countyRpc/QueryDetailByEditor"
	CountyRpc_DeleteByEditor_FullMethodName      = "/countyService.countyRpc/DeleteByEditor"
	CountyRpc_UpdateByEditor_FullMethodName      = "/countyService.countyRpc/UpdateByEditor"
)

// CountyRpcClient is the client API for CountyRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountyRpcClient interface {
	Insert(ctx context.Context, in *models.CountyEntity, opts ...grpc.CallOption) (*models.NormalResult, error)
	QueryDetailByEditor(ctx context.Context, in *models.EditorRequest, opts ...grpc.CallOption) (*models.CountyEntity, error)
	DeleteByEditor(ctx context.Context, in *models.EditorRequest, opts ...grpc.CallOption) (*models.NormalResult, error)
	UpdateByEditor(ctx context.Context, in *models.UpdateRequest, opts ...grpc.CallOption) (*models.NormalResult, error)
}

type countyRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCountyRpcClient(cc grpc.ClientConnInterface) CountyRpcClient {
	return &countyRpcClient{cc}
}

func (c *countyRpcClient) Insert(ctx context.Context, in *models.CountyEntity, opts ...grpc.CallOption) (*models.NormalResult, error) {
	out := new(models.NormalResult)
	err := c.cc.Invoke(ctx, CountyRpc_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countyRpcClient) QueryDetailByEditor(ctx context.Context, in *models.EditorRequest, opts ...grpc.CallOption) (*models.CountyEntity, error) {
	out := new(models.CountyEntity)
	err := c.cc.Invoke(ctx, CountyRpc_QueryDetailByEditor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countyRpcClient) DeleteByEditor(ctx context.Context, in *models.EditorRequest, opts ...grpc.CallOption) (*models.NormalResult, error) {
	out := new(models.NormalResult)
	err := c.cc.Invoke(ctx, CountyRpc_DeleteByEditor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countyRpcClient) UpdateByEditor(ctx context.Context, in *models.UpdateRequest, opts ...grpc.CallOption) (*models.NormalResult, error) {
	out := new(models.NormalResult)
	err := c.cc.Invoke(ctx, CountyRpc_UpdateByEditor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountyRpcServer is the server API for CountyRpc service.
// All implementations must embed UnimplementedCountyRpcServer
// for forward compatibility
type CountyRpcServer interface {
	Insert(context.Context, *models.CountyEntity) (*models.NormalResult, error)
	QueryDetailByEditor(context.Context, *models.EditorRequest) (*models.CountyEntity, error)
	DeleteByEditor(context.Context, *models.EditorRequest) (*models.NormalResult, error)
	UpdateByEditor(context.Context, *models.UpdateRequest) (*models.NormalResult, error)
	mustEmbedUnimplementedCountyRpcServer()
}

// UnimplementedCountyRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCountyRpcServer struct {
}

func (UnimplementedCountyRpcServer) Insert(context.Context, *models.CountyEntity) (*models.NormalResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedCountyRpcServer) QueryDetailByEditor(context.Context, *models.EditorRequest) (*models.CountyEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDetailByEditor not implemented")
}
func (UnimplementedCountyRpcServer) DeleteByEditor(context.Context, *models.EditorRequest) (*models.NormalResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByEditor not implemented")
}
func (UnimplementedCountyRpcServer) UpdateByEditor(context.Context, *models.UpdateRequest) (*models.NormalResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByEditor not implemented")
}
func (UnimplementedCountyRpcServer) mustEmbedUnimplementedCountyRpcServer() {}

// UnsafeCountyRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountyRpcServer will
// result in compilation errors.
type UnsafeCountyRpcServer interface {
	mustEmbedUnimplementedCountyRpcServer()
}

func RegisterCountyRpcServer(s grpc.ServiceRegistrar, srv CountyRpcServer) {
	s.RegisterService(&CountyRpc_ServiceDesc, srv)
}

func _CountyRpc_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.CountyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountyRpcServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountyRpc_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountyRpcServer).Insert(ctx, req.(*models.CountyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountyRpc_QueryDetailByEditor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.EditorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountyRpcServer).QueryDetailByEditor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountyRpc_QueryDetailByEditor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountyRpcServer).QueryDetailByEditor(ctx, req.(*models.EditorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountyRpc_DeleteByEditor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.EditorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountyRpcServer).DeleteByEditor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountyRpc_DeleteByEditor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountyRpcServer).DeleteByEditor(ctx, req.(*models.EditorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountyRpc_UpdateByEditor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(models.UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountyRpcServer).UpdateByEditor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CountyRpc_UpdateByEditor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountyRpcServer).UpdateByEditor(ctx, req.(*models.UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountyRpc_ServiceDesc is the grpc.ServiceDesc for CountyRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountyRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "countyService.countyRpc",
	HandlerType: (*CountyRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _CountyRpc_Insert_Handler,
		},
		{
			MethodName: "QueryDetailByEditor",
			Handler:    _CountyRpc_QueryDetailByEditor_Handler,
		},
		{
			MethodName: "DeleteByEditor",
			Handler:    _CountyRpc_DeleteByEditor_Handler,
		},
		{
			MethodName: "UpdateByEditor",
			Handler:    _CountyRpc_UpdateByEditor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "county.proto",
}