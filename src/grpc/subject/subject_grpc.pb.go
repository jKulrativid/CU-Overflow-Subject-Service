// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: subject.proto

package subject

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

// SubjectServiceClient is the client API for SubjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubjectServiceClient interface {
	PaginateSubjects(ctx context.Context, in *PaginateSubjectRequest, opts ...grpc.CallOption) (*PaginateSubjectResponse, error)
	GetSubjectById(ctx context.Context, in *GetSubjectByIdRequest, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error)
	ValidateSubjectId(ctx context.Context, in *ValidateSubjectIdRequest, opts ...grpc.CallOption) (*ValidateSubjectIdResponse, error)
	CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*CreateSubjectResponse, error)
	UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*UpdateSubjectResponse, error)
	DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*DeleteSubjectResponse, error)
	ValidateSection(ctx context.Context, in *ValidateSectionRequest, opts ...grpc.CallOption) (*ValidateSectionResponse, error)
	CreateSection(ctx context.Context, in *CreateSectionRequest, opts ...grpc.CallOption) (*CreateSectionResponse, error)
	UpdateSection(ctx context.Context, in *UpdateSectionRequest, opts ...grpc.CallOption) (*UpdateSectionResponse, error)
	DeleteSection(ctx context.Context, in *DeleteSectionRequest, opts ...grpc.CallOption) (*DeleteSectionResponse, error)
	PaginatePostBySubject(ctx context.Context, in *PaginatePostBySubjectRequest, opts ...grpc.CallOption) (*PaginatePostBySubjectResponse, error)
	PaginateFileBySubject(ctx context.Context, in *PaginateFileBySubjectRequest, opts ...grpc.CallOption) (*PaginateFileBySubjectResponse, error)
}

type subjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubjectServiceClient(cc grpc.ClientConnInterface) SubjectServiceClient {
	return &subjectServiceClient{cc}
}

func (c *subjectServiceClient) PaginateSubjects(ctx context.Context, in *PaginateSubjectRequest, opts ...grpc.CallOption) (*PaginateSubjectResponse, error) {
	out := new(PaginateSubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/PaginateSubjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) GetSubjectById(ctx context.Context, in *GetSubjectByIdRequest, opts ...grpc.CallOption) (*GetSubjectByIdResponse, error) {
	out := new(GetSubjectByIdResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/GetSubjectById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) ValidateSubjectId(ctx context.Context, in *ValidateSubjectIdRequest, opts ...grpc.CallOption) (*ValidateSubjectIdResponse, error) {
	out := new(ValidateSubjectIdResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/ValidateSubjectId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*CreateSubjectResponse, error) {
	out := new(CreateSubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/CreateSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) UpdateSubject(ctx context.Context, in *UpdateSubjectRequest, opts ...grpc.CallOption) (*UpdateSubjectResponse, error) {
	out := new(UpdateSubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/UpdateSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*DeleteSubjectResponse, error) {
	out := new(DeleteSubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/DeleteSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) ValidateSection(ctx context.Context, in *ValidateSectionRequest, opts ...grpc.CallOption) (*ValidateSectionResponse, error) {
	out := new(ValidateSectionResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/ValidateSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) CreateSection(ctx context.Context, in *CreateSectionRequest, opts ...grpc.CallOption) (*CreateSectionResponse, error) {
	out := new(CreateSectionResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/CreateSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) UpdateSection(ctx context.Context, in *UpdateSectionRequest, opts ...grpc.CallOption) (*UpdateSectionResponse, error) {
	out := new(UpdateSectionResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/UpdateSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) DeleteSection(ctx context.Context, in *DeleteSectionRequest, opts ...grpc.CallOption) (*DeleteSectionResponse, error) {
	out := new(DeleteSectionResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/DeleteSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) PaginatePostBySubject(ctx context.Context, in *PaginatePostBySubjectRequest, opts ...grpc.CallOption) (*PaginatePostBySubjectResponse, error) {
	out := new(PaginatePostBySubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/PaginatePostBySubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subjectServiceClient) PaginateFileBySubject(ctx context.Context, in *PaginateFileBySubjectRequest, opts ...grpc.CallOption) (*PaginateFileBySubjectResponse, error) {
	out := new(PaginateFileBySubjectResponse)
	err := c.cc.Invoke(ctx, "/subject.SubjectService/PaginateFileBySubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubjectServiceServer is the server API for SubjectService service.
// All implementations must embed UnimplementedSubjectServiceServer
// for forward compatibility
type SubjectServiceServer interface {
	PaginateSubjects(context.Context, *PaginateSubjectRequest) (*PaginateSubjectResponse, error)
	GetSubjectById(context.Context, *GetSubjectByIdRequest) (*GetSubjectByIdResponse, error)
	ValidateSubjectId(context.Context, *ValidateSubjectIdRequest) (*ValidateSubjectIdResponse, error)
	CreateSubject(context.Context, *CreateSubjectRequest) (*CreateSubjectResponse, error)
	UpdateSubject(context.Context, *UpdateSubjectRequest) (*UpdateSubjectResponse, error)
	DeleteSubject(context.Context, *DeleteSubjectRequest) (*DeleteSubjectResponse, error)
	ValidateSection(context.Context, *ValidateSectionRequest) (*ValidateSectionResponse, error)
	CreateSection(context.Context, *CreateSectionRequest) (*CreateSectionResponse, error)
	UpdateSection(context.Context, *UpdateSectionRequest) (*UpdateSectionResponse, error)
	DeleteSection(context.Context, *DeleteSectionRequest) (*DeleteSectionResponse, error)
	PaginatePostBySubject(context.Context, *PaginatePostBySubjectRequest) (*PaginatePostBySubjectResponse, error)
	PaginateFileBySubject(context.Context, *PaginateFileBySubjectRequest) (*PaginateFileBySubjectResponse, error)
	mustEmbedUnimplementedSubjectServiceServer()
}

// UnimplementedSubjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubjectServiceServer struct {
}

func (UnimplementedSubjectServiceServer) PaginateSubjects(context.Context, *PaginateSubjectRequest) (*PaginateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaginateSubjects not implemented")
}
func (UnimplementedSubjectServiceServer) GetSubjectById(context.Context, *GetSubjectByIdRequest) (*GetSubjectByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectById not implemented")
}
func (UnimplementedSubjectServiceServer) ValidateSubjectId(context.Context, *ValidateSubjectIdRequest) (*ValidateSubjectIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSubjectId not implemented")
}
func (UnimplementedSubjectServiceServer) CreateSubject(context.Context, *CreateSubjectRequest) (*CreateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) UpdateSubject(context.Context, *UpdateSubjectRequest) (*UpdateSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubject not implemented")
}
func (UnimplementedSubjectServiceServer) DeleteSubject(context.Context, *DeleteSubjectRequest) (*DeleteSubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubject not implemented")
}
func (UnimplementedSubjectServiceServer) ValidateSection(context.Context, *ValidateSectionRequest) (*ValidateSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateSection not implemented")
}
func (UnimplementedSubjectServiceServer) CreateSection(context.Context, *CreateSectionRequest) (*CreateSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSection not implemented")
}
func (UnimplementedSubjectServiceServer) UpdateSection(context.Context, *UpdateSectionRequest) (*UpdateSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSection not implemented")
}
func (UnimplementedSubjectServiceServer) DeleteSection(context.Context, *DeleteSectionRequest) (*DeleteSectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSection not implemented")
}
func (UnimplementedSubjectServiceServer) PaginatePostBySubject(context.Context, *PaginatePostBySubjectRequest) (*PaginatePostBySubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaginatePostBySubject not implemented")
}
func (UnimplementedSubjectServiceServer) PaginateFileBySubject(context.Context, *PaginateFileBySubjectRequest) (*PaginateFileBySubjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaginateFileBySubject not implemented")
}
func (UnimplementedSubjectServiceServer) mustEmbedUnimplementedSubjectServiceServer() {}

// UnsafeSubjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubjectServiceServer will
// result in compilation errors.
type UnsafeSubjectServiceServer interface {
	mustEmbedUnimplementedSubjectServiceServer()
}

func RegisterSubjectServiceServer(s grpc.ServiceRegistrar, srv SubjectServiceServer) {
	s.RegisterService(&SubjectService_ServiceDesc, srv)
}

func _SubjectService_PaginateSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).PaginateSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/PaginateSubjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).PaginateSubjects(ctx, req.(*PaginateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_GetSubjectById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).GetSubjectById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/GetSubjectById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).GetSubjectById(ctx, req.(*GetSubjectByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_ValidateSubjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSubjectIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).ValidateSubjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/ValidateSubjectId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).ValidateSubjectId(ctx, req.(*ValidateSubjectIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_CreateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).CreateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/CreateSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).CreateSubject(ctx, req.(*CreateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_UpdateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/UpdateSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).UpdateSubject(ctx, req.(*UpdateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_DeleteSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/DeleteSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).DeleteSubject(ctx, req.(*DeleteSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_ValidateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).ValidateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/ValidateSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).ValidateSection(ctx, req.(*ValidateSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_CreateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).CreateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/CreateSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).CreateSection(ctx, req.(*CreateSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_UpdateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).UpdateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/UpdateSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).UpdateSection(ctx, req.(*UpdateSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_DeleteSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).DeleteSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/DeleteSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).DeleteSection(ctx, req.(*DeleteSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_PaginatePostBySubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginatePostBySubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).PaginatePostBySubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/PaginatePostBySubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).PaginatePostBySubject(ctx, req.(*PaginatePostBySubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubjectService_PaginateFileBySubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginateFileBySubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubjectServiceServer).PaginateFileBySubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/subject.SubjectService/PaginateFileBySubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubjectServiceServer).PaginateFileBySubject(ctx, req.(*PaginateFileBySubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubjectService_ServiceDesc is the grpc.ServiceDesc for SubjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subject.SubjectService",
	HandlerType: (*SubjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PaginateSubjects",
			Handler:    _SubjectService_PaginateSubjects_Handler,
		},
		{
			MethodName: "GetSubjectById",
			Handler:    _SubjectService_GetSubjectById_Handler,
		},
		{
			MethodName: "ValidateSubjectId",
			Handler:    _SubjectService_ValidateSubjectId_Handler,
		},
		{
			MethodName: "CreateSubject",
			Handler:    _SubjectService_CreateSubject_Handler,
		},
		{
			MethodName: "UpdateSubject",
			Handler:    _SubjectService_UpdateSubject_Handler,
		},
		{
			MethodName: "DeleteSubject",
			Handler:    _SubjectService_DeleteSubject_Handler,
		},
		{
			MethodName: "ValidateSection",
			Handler:    _SubjectService_ValidateSection_Handler,
		},
		{
			MethodName: "CreateSection",
			Handler:    _SubjectService_CreateSection_Handler,
		},
		{
			MethodName: "UpdateSection",
			Handler:    _SubjectService_UpdateSection_Handler,
		},
		{
			MethodName: "DeleteSection",
			Handler:    _SubjectService_DeleteSection_Handler,
		},
		{
			MethodName: "PaginatePostBySubject",
			Handler:    _SubjectService_PaginatePostBySubject_Handler,
		},
		{
			MethodName: "PaginateFileBySubject",
			Handler:    _SubjectService_PaginateFileBySubject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subject.proto",
}
