// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/organizations_v1.proto

package protos

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

// OrganizationsClient is the client API for Organizations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationsClient interface {
	GetOrganizations(ctx context.Context, in *OrganizationPageRequest, opts ...grpc.CallOption) (*OrganizationPageReply, error)
	GetOrganizationById(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error)
	GetOrganizationByCode(ctx context.Context, in *OrganizationCodeRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error)
	GenerateCode(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationCodeReply, error)
	CreateOrganization(ctx context.Context, in *OrganizationObjectRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error)
	UpdateOrganization(ctx context.Context, in *OrganizationObjectRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error)
	DeleteOrganizationById(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error)
}

type organizationsClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationsClient(cc grpc.ClientConnInterface) OrganizationsClient {
	return &organizationsClient{cc}
}

func (c *organizationsClient) GetOrganizations(ctx context.Context, in *OrganizationPageRequest, opts ...grpc.CallOption) (*OrganizationPageReply, error) {
	out := new(OrganizationPageReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/get_organizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) GetOrganizationById(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error) {
	out := new(OrganizationObjectReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/get_organization_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) GetOrganizationByCode(ctx context.Context, in *OrganizationCodeRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error) {
	out := new(OrganizationObjectReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/get_organization_by_code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) GenerateCode(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationCodeReply, error) {
	out := new(OrganizationCodeReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/generate_code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) CreateOrganization(ctx context.Context, in *OrganizationObjectRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error) {
	out := new(OrganizationObjectReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/create_organization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) UpdateOrganization(ctx context.Context, in *OrganizationObjectRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error) {
	out := new(OrganizationObjectReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/update_organization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsClient) DeleteOrganizationById(ctx context.Context, in *OrgIdRequest, opts ...grpc.CallOption) (*OrganizationObjectReply, error) {
	out := new(OrganizationObjectReply)
	err := c.cc.Invoke(ctx, "/organizations_v1.Organizations/delete_organization_by_id", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationsServer is the server API for Organizations service.
// All implementations must embed UnimplementedOrganizationsServer
// for forward compatibility
type OrganizationsServer interface {
	GetOrganizations(context.Context, *OrganizationPageRequest) (*OrganizationPageReply, error)
	GetOrganizationById(context.Context, *OrgIdRequest) (*OrganizationObjectReply, error)
	GetOrganizationByCode(context.Context, *OrganizationCodeRequest) (*OrganizationObjectReply, error)
	GenerateCode(context.Context, *OrgIdRequest) (*OrganizationCodeReply, error)
	CreateOrganization(context.Context, *OrganizationObjectRequest) (*OrganizationObjectReply, error)
	UpdateOrganization(context.Context, *OrganizationObjectRequest) (*OrganizationObjectReply, error)
	DeleteOrganizationById(context.Context, *OrgIdRequest) (*OrganizationObjectReply, error)
	mustEmbedUnimplementedOrganizationsServer()
}

// UnimplementedOrganizationsServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationsServer struct {
}

func (UnimplementedOrganizationsServer) GetOrganizations(context.Context, *OrganizationPageRequest) (*OrganizationPageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizations not implemented")
}
func (UnimplementedOrganizationsServer) GetOrganizationById(context.Context, *OrgIdRequest) (*OrganizationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationById not implemented")
}
func (UnimplementedOrganizationsServer) GetOrganizationByCode(context.Context, *OrganizationCodeRequest) (*OrganizationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationByCode not implemented")
}
func (UnimplementedOrganizationsServer) GenerateCode(context.Context, *OrgIdRequest) (*OrganizationCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCode not implemented")
}
func (UnimplementedOrganizationsServer) CreateOrganization(context.Context, *OrganizationObjectRequest) (*OrganizationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizationsServer) UpdateOrganization(context.Context, *OrganizationObjectRequest) (*OrganizationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrganization not implemented")
}
func (UnimplementedOrganizationsServer) DeleteOrganizationById(context.Context, *OrgIdRequest) (*OrganizationObjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrganizationById not implemented")
}
func (UnimplementedOrganizationsServer) mustEmbedUnimplementedOrganizationsServer() {}

// UnsafeOrganizationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationsServer will
// result in compilation errors.
type UnsafeOrganizationsServer interface {
	mustEmbedUnimplementedOrganizationsServer()
}

func RegisterOrganizationsServer(s grpc.ServiceRegistrar, srv OrganizationsServer) {
	s.RegisterService(&Organizations_ServiceDesc, srv)
}

func _Organizations_GetOrganizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationPageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).GetOrganizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/get_organizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).GetOrganizations(ctx, req.(*OrganizationPageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_GetOrganizationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).GetOrganizationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/get_organization_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).GetOrganizationById(ctx, req.(*OrgIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_GetOrganizationByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).GetOrganizationByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/get_organization_by_code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).GetOrganizationByCode(ctx, req.(*OrganizationCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_GenerateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).GenerateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/generate_code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).GenerateCode(ctx, req.(*OrgIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/create_organization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).CreateOrganization(ctx, req.(*OrganizationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_UpdateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrganizationObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).UpdateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/update_organization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).UpdateOrganization(ctx, req.(*OrganizationObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organizations_DeleteOrganizationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServer).DeleteOrganizationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organizations_v1.Organizations/delete_organization_by_id",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServer).DeleteOrganizationById(ctx, req.(*OrgIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Organizations_ServiceDesc is the grpc.ServiceDesc for Organizations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Organizations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organizations_v1.Organizations",
	HandlerType: (*OrganizationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get_organizations",
			Handler:    _Organizations_GetOrganizations_Handler,
		},
		{
			MethodName: "get_organization_by_id",
			Handler:    _Organizations_GetOrganizationById_Handler,
		},
		{
			MethodName: "get_organization_by_code",
			Handler:    _Organizations_GetOrganizationByCode_Handler,
		},
		{
			MethodName: "generate_code",
			Handler:    _Organizations_GenerateCode_Handler,
		},
		{
			MethodName: "create_organization",
			Handler:    _Organizations_CreateOrganization_Handler,
		},
		{
			MethodName: "update_organization",
			Handler:    _Organizations_UpdateOrganization_Handler,
		},
		{
			MethodName: "delete_organization_by_id",
			Handler:    _Organizations_DeleteOrganizationById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/organizations_v1.proto",
}
