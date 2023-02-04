// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: catalog.proto

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

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
	ReadCategoriesByProjectId(ctx context.Context, in *ReadCategoriesByProjectIdRequest, opts ...grpc.CallOption) (*ReadCategoriesByProjectIdResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	ReadProductsOfProject(ctx context.Context, in *ReadProductsOfProjectRequest, opts ...grpc.CallOption) (*ReadProductsOfProjectResponse, error)
	ReadProductsOfCategory(ctx context.Context, in *ReadProductsOfCategoryRequest, opts ...grpc.CallOption) (*ReadProductsOfCategoryResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*UpdateCategoryResponse, error) {
	out := new(UpdateCategoryResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) ReadCategoriesByProjectId(ctx context.Context, in *ReadCategoriesByProjectIdRequest, opts ...grpc.CallOption) (*ReadCategoriesByProjectIdResponse, error) {
	out := new(ReadCategoriesByProjectIdResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/ReadCategoriesByProjectId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) ReadProductsOfProject(ctx context.Context, in *ReadProductsOfProjectRequest, opts ...grpc.CallOption) (*ReadProductsOfProjectResponse, error) {
	out := new(ReadProductsOfProjectResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/ReadProductsOfProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) ReadProductsOfCategory(ctx context.Context, in *ReadProductsOfCategoryRequest, opts ...grpc.CallOption) (*ReadProductsOfCategoryResponse, error) {
	out := new(ReadProductsOfCategoryResponse)
	err := c.cc.Invoke(ctx, "/grpc.Catalog/ReadProductsOfCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
	ReadCategoriesByProjectId(context.Context, *ReadCategoriesByProjectIdRequest) (*ReadCategoriesByProjectIdResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	ReadProductsOfProject(context.Context, *ReadProductsOfProjectRequest) (*ReadProductsOfProjectResponse, error)
	ReadProductsOfCategory(context.Context, *ReadProductsOfCategoryRequest) (*ReadProductsOfCategoryResponse, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedCatalogServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*UpdateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedCatalogServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCatalogServer) ReadCategoriesByProjectId(context.Context, *ReadCategoriesByProjectIdRequest) (*ReadCategoriesByProjectIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCategoriesByProjectId not implemented")
}
func (UnimplementedCatalogServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedCatalogServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedCatalogServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedCatalogServer) ReadProductsOfProject(context.Context, *ReadProductsOfProjectRequest) (*ReadProductsOfProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProductsOfProject not implemented")
}
func (UnimplementedCatalogServer) ReadProductsOfCategory(context.Context, *ReadProductsOfCategoryRequest) (*ReadProductsOfCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProductsOfCategory not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_ReadCategoriesByProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCategoriesByProjectIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ReadCategoriesByProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/ReadCategoriesByProjectId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ReadCategoriesByProjectId(ctx, req.(*ReadCategoriesByProjectIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_ReadProductsOfProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProductsOfProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ReadProductsOfProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/ReadProductsOfProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ReadProductsOfProject(ctx, req.(*ReadProductsOfProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_ReadProductsOfCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadProductsOfCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).ReadProductsOfCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Catalog/ReadProductsOfCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).ReadProductsOfCategory(ctx, req.(*ReadProductsOfCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _Catalog_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _Catalog_UpdateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Catalog_DeleteCategory_Handler,
		},
		{
			MethodName: "ReadCategoriesByProjectId",
			Handler:    _Catalog_ReadCategoriesByProjectId_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _Catalog_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Catalog_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _Catalog_DeleteProduct_Handler,
		},
		{
			MethodName: "ReadProductsOfProject",
			Handler:    _Catalog_ReadProductsOfProject_Handler,
		},
		{
			MethodName: "ReadProductsOfCategory",
			Handler:    _Catalog_ReadProductsOfCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
