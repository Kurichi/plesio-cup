// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: schema/tree.proto

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

const (
	TreeService_InitTree_FullMethodName        = "/tree.TreeService/InitTree"
	TreeService_PlantTree_FullMethodName       = "/tree.TreeService/PlantTree"
	TreeService_GetTreeByUserId_FullMethodName = "/tree.TreeService/GetTreeByUserId"
	TreeService_GetMyTree_FullMethodName       = "/tree.TreeService/GetMyTree"
	TreeService_GetTreeRanking_FullMethodName  = "/tree.TreeService/GetTreeRanking"
)

// TreeServiceClient is the client API for TreeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TreeServiceClient interface {
	InitTree(ctx context.Context, in *PlantTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error)
	PlantTree(ctx context.Context, in *PlantTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error)
	GetTreeByUserId(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error)
	GetMyTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error)
	GetTreeRanking(ctx context.Context, in *GetTreeRankingRequest, opts ...grpc.CallOption) (*GetTreeRankingResponse, error)
}

type treeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTreeServiceClient(cc grpc.ClientConnInterface) TreeServiceClient {
	return &treeServiceClient{cc}
}

func (c *treeServiceClient) InitTree(ctx context.Context, in *PlantTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error) {
	out := new(GetTreeResponse)
	err := c.cc.Invoke(ctx, TreeService_InitTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) PlantTree(ctx context.Context, in *PlantTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error) {
	out := new(GetTreeResponse)
	err := c.cc.Invoke(ctx, TreeService_PlantTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) GetTreeByUserId(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error) {
	out := new(GetTreeResponse)
	err := c.cc.Invoke(ctx, TreeService_GetTreeByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) GetMyTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetTreeResponse, error) {
	out := new(GetTreeResponse)
	err := c.cc.Invoke(ctx, TreeService_GetMyTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treeServiceClient) GetTreeRanking(ctx context.Context, in *GetTreeRankingRequest, opts ...grpc.CallOption) (*GetTreeRankingResponse, error) {
	out := new(GetTreeRankingResponse)
	err := c.cc.Invoke(ctx, TreeService_GetTreeRanking_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TreeServiceServer is the server API for TreeService service.
// All implementations must embed UnimplementedTreeServiceServer
// for forward compatibility
type TreeServiceServer interface {
	InitTree(context.Context, *PlantTreeRequest) (*GetTreeResponse, error)
	PlantTree(context.Context, *PlantTreeRequest) (*GetTreeResponse, error)
	GetTreeByUserId(context.Context, *GetTreeRequest) (*GetTreeResponse, error)
	GetMyTree(context.Context, *GetTreeRequest) (*GetTreeResponse, error)
	GetTreeRanking(context.Context, *GetTreeRankingRequest) (*GetTreeRankingResponse, error)
	mustEmbedUnimplementedTreeServiceServer()
}

// UnimplementedTreeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTreeServiceServer struct {
}

func (UnimplementedTreeServiceServer) InitTree(context.Context, *PlantTreeRequest) (*GetTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitTree not implemented")
}
func (UnimplementedTreeServiceServer) PlantTree(context.Context, *PlantTreeRequest) (*GetTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlantTree not implemented")
}
func (UnimplementedTreeServiceServer) GetTreeByUserId(context.Context, *GetTreeRequest) (*GetTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTreeByUserId not implemented")
}
func (UnimplementedTreeServiceServer) GetMyTree(context.Context, *GetTreeRequest) (*GetTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyTree not implemented")
}
func (UnimplementedTreeServiceServer) GetTreeRanking(context.Context, *GetTreeRankingRequest) (*GetTreeRankingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTreeRanking not implemented")
}
func (UnimplementedTreeServiceServer) mustEmbedUnimplementedTreeServiceServer() {}

// UnsafeTreeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TreeServiceServer will
// result in compilation errors.
type UnsafeTreeServiceServer interface {
	mustEmbedUnimplementedTreeServiceServer()
}

func RegisterTreeServiceServer(s grpc.ServiceRegistrar, srv TreeServiceServer) {
	s.RegisterService(&TreeService_ServiceDesc, srv)
}

func _TreeService_InitTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).InitTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TreeService_InitTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).InitTree(ctx, req.(*PlantTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_PlantTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlantTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).PlantTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TreeService_PlantTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).PlantTree(ctx, req.(*PlantTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_GetTreeByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).GetTreeByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TreeService_GetTreeByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).GetTreeByUserId(ctx, req.(*GetTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_GetMyTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).GetMyTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TreeService_GetMyTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).GetMyTree(ctx, req.(*GetTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TreeService_GetTreeRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRankingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreeServiceServer).GetTreeRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TreeService_GetTreeRanking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreeServiceServer).GetTreeRanking(ctx, req.(*GetTreeRankingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TreeService_ServiceDesc is the grpc.ServiceDesc for TreeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TreeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tree.TreeService",
	HandlerType: (*TreeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitTree",
			Handler:    _TreeService_InitTree_Handler,
		},
		{
			MethodName: "PlantTree",
			Handler:    _TreeService_PlantTree_Handler,
		},
		{
			MethodName: "GetTreeByUserId",
			Handler:    _TreeService_GetTreeByUserId_Handler,
		},
		{
			MethodName: "GetMyTree",
			Handler:    _TreeService_GetMyTree_Handler,
		},
		{
			MethodName: "GetTreeRanking",
			Handler:    _TreeService_GetTreeRanking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/tree.proto",
}
