// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/agent.proto

package service

import (
	context "context"
	fmt "fmt"
	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("flyteidl/service/agent.proto", fileDescriptor_f7d1dfd1fb77d2ef) }

var fileDescriptor_f7d1dfd1fb77d2ef = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xbf, 0x4f, 0x41, 0x25, 0x42, 0xa9, 0xa1, 0x74, 0x31, 0x14, 0x7f, 0x2a, 0xba, 0x73,
	0x82, 0x75, 0x2d, 0x52, 0x15, 0xba, 0x69, 0x37, 0x6a, 0x41, 0xdc, 0xa5, 0xd3, 0x63, 0x0c, 0x9d,
	0x26, 0xe3, 0xe4, 0xb4, 0x50, 0xc4, 0x8d, 0xb7, 0xe0, 0x6d, 0xb9, 0xf3, 0x16, 0xbc, 0x00, 0x2f,
	0x41, 0x9a, 0x99, 0x4c, 0x6b, 0x35, 0x75, 0x57, 0xfa, 0xbc, 0x79, 0xde, 0x30, 0xe7, 0x84, 0xd4,
	0xee, 0xe3, 0x09, 0x82, 0xec, 0xc7, 0xcc, 0x40, 0x3a, 0x96, 0x11, 0x30, 0x2e, 0x40, 0x61, 0x98,
	0xa4, 0x1a, 0x35, 0x2d, 0x3b, 0x1a, 0xe6, 0x34, 0xa8, 0x09, 0xad, 0x45, 0x0c, 0x8c, 0x27, 0x92,
	0x71, 0xa5, 0x34, 0x72, 0x94, 0x5a, 0x99, 0x2c, 0x1f, 0x04, 0x85, 0x8d, 0xf7, 0x87, 0x52, 0xcd,
	0xbb, 0x1a, 0x6f, 0xab, 0x64, 0xab, 0x69, 0x26, 0x2a, 0x6a, 0x4e, 0xff, 0xbc, 0xce, 0x7c, 0xb4,
	0x4b, 0xc8, 0x45, 0x0a, 0x1c, 0xe1, 0x86, 0x9b, 0x01, 0xdd, 0x0b, 0x8b, 0x42, 0x2b, 0x08, 0x67,
	0xec, 0x0a, 0x1e, 0x47, 0x60, 0x30, 0xa8, 0x2f, 0x8b, 0x98, 0x44, 0x2b, 0x03, 0xf5, 0x7f, 0xb4,
	0x4d, 0xd6, 0x5b, 0x80, 0xd6, 0xb9, 0xbd, 0x78, 0x20, 0x07, 0x4e, 0xb8, 0xe3, 0xe5, 0x85, 0xad,
	0x4b, 0xc8, 0x25, 0xc4, 0xe0, 0xbb, 0xe4, 0x8c, 0x79, 0x2f, 0x39, 0x1f, 0x29, 0xb4, 0x9c, 0x94,
	0xf2, 0xae, 0x0e, 0x60, 0x2a, 0x23, 0x43, 0x0f, 0x3c, 0x77, 0xc9, 0xb9, 0xd3, 0x1f, 0xfe, 0x15,
	0x2b, 0x2a, 0x6e, 0xc9, 0x66, 0xce, 0xda, 0x5a, 0x18, 0x5a, 0xf7, 0x1c, 0x9c, 0x42, 0x27, 0xdf,
	0x5f, 0x9a, 0x71, 0xe6, 0xc6, 0xe7, 0x7f, 0x52, 0xb1, 0x93, 0xec, 0x00, 0xf2, 0x3e, 0x47, 0xee,
	0x26, 0x3a, 0x20, 0x1b, 0x2d, 0x40, 0x8b, 0xe8, 0x6f, 0xdf, 0xd6, 0x12, 0x57, 0xb6, 0xeb, 0x0f,
	0xe4, 0x4d, 0xb5, 0x97, 0xf7, 0x8f, 0xd7, 0x95, 0x2a, 0xad, 0xd8, 0x95, 0x1b, 0x1f, 0x67, 0x3b,
	0xc5, 0x9e, 0x14, 0x1f, 0xc2, 0x33, 0x1d, 0x10, 0xd2, 0x96, 0x26, 0x3b, 0x62, 0x7e, 0x4e, 0x66,
	0xc6, 0xbc, 0x93, 0x99, 0x8f, 0xe4, 0x95, 0x55, 0x5b, 0x59, 0xa6, 0xa5, 0x6f, 0x95, 0xe6, 0xfc,
	0xec, 0xee, 0x54, 0x48, 0x7c, 0x18, 0xf5, 0xc2, 0x48, 0x0f, 0x99, 0xf5, 0xe8, 0x54, 0x64, 0x3f,
	0x58, 0xb1, 0xf9, 0x02, 0x14, 0x4b, 0x7a, 0x47, 0x42, 0xb3, 0xc5, 0xa7, 0xd5, 0x5b, 0xb3, 0x2f,
	0xe1, 0xe4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x64, 0x33, 0x1f, 0x2d, 0x75, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AsyncAgentServiceClient is the client API for AsyncAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AsyncAgentServiceClient interface {
	// Send a task create request to the agent server.
	CreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) (*admin.CreateTaskResponse, error)
	// Get job status.
	GetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) (*admin.GetTaskResponse, error)
	// Delete the task resource.
	DeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) (*admin.DeleteTaskResponse, error)
	// GetTaskMetrics returns one or more task execution metrics, if available.
	//
	// Errors include
	//  * OutOfRange if metrics are not available for the specified task time range
	//  * various other errors
	GetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) (*admin.GetTaskMetricsResponse, error)
	// GetTaskLogs returns task execution logs, if available.
	GetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) (*admin.GetTaskLogsResponse, error)
}

type asyncAgentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAsyncAgentServiceClient(cc *grpc.ClientConn) AsyncAgentServiceClient {
	return &asyncAgentServiceClient{cc}
}

func (c *asyncAgentServiceClient) CreateTask(ctx context.Context, in *admin.CreateTaskRequest, opts ...grpc.CallOption) (*admin.CreateTaskResponse, error) {
	out := new(admin.CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AsyncAgentService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTask(ctx context.Context, in *admin.GetTaskRequest, opts ...grpc.CallOption) (*admin.GetTaskResponse, error) {
	out := new(admin.GetTaskResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AsyncAgentService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) DeleteTask(ctx context.Context, in *admin.DeleteTaskRequest, opts ...grpc.CallOption) (*admin.DeleteTaskResponse, error) {
	out := new(admin.DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AsyncAgentService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTaskMetrics(ctx context.Context, in *admin.GetTaskMetricsRequest, opts ...grpc.CallOption) (*admin.GetTaskMetricsResponse, error) {
	out := new(admin.GetTaskMetricsResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AsyncAgentService/GetTaskMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *asyncAgentServiceClient) GetTaskLogs(ctx context.Context, in *admin.GetTaskLogsRequest, opts ...grpc.CallOption) (*admin.GetTaskLogsResponse, error) {
	out := new(admin.GetTaskLogsResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AsyncAgentService/GetTaskLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AsyncAgentServiceServer is the server API for AsyncAgentService service.
type AsyncAgentServiceServer interface {
	// Send a task create request to the agent server.
	CreateTask(context.Context, *admin.CreateTaskRequest) (*admin.CreateTaskResponse, error)
	// Get job status.
	GetTask(context.Context, *admin.GetTaskRequest) (*admin.GetTaskResponse, error)
	// Delete the task resource.
	DeleteTask(context.Context, *admin.DeleteTaskRequest) (*admin.DeleteTaskResponse, error)
	// GetTaskMetrics returns one or more task execution metrics, if available.
	//
	// Errors include
	//  * OutOfRange if metrics are not available for the specified task time range
	//  * various other errors
	GetTaskMetrics(context.Context, *admin.GetTaskMetricsRequest) (*admin.GetTaskMetricsResponse, error)
	// GetTaskLogs returns task execution logs, if available.
	GetTaskLogs(context.Context, *admin.GetTaskLogsRequest) (*admin.GetTaskLogsResponse, error)
}

// UnimplementedAsyncAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAsyncAgentServiceServer struct {
}

func (*UnimplementedAsyncAgentServiceServer) CreateTask(ctx context.Context, req *admin.CreateTaskRequest) (*admin.CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedAsyncAgentServiceServer) GetTask(ctx context.Context, req *admin.GetTaskRequest) (*admin.GetTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (*UnimplementedAsyncAgentServiceServer) DeleteTask(ctx context.Context, req *admin.DeleteTaskRequest) (*admin.DeleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (*UnimplementedAsyncAgentServiceServer) GetTaskMetrics(ctx context.Context, req *admin.GetTaskMetricsRequest) (*admin.GetTaskMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskMetrics not implemented")
}
func (*UnimplementedAsyncAgentServiceServer) GetTaskLogs(ctx context.Context, req *admin.GetTaskLogsRequest) (*admin.GetTaskLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLogs not implemented")
}

func RegisterAsyncAgentServiceServer(s *grpc.Server, srv AsyncAgentServiceServer) {
	s.RegisterService(&_AsyncAgentService_serviceDesc, srv)
}

func _AsyncAgentService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AsyncAgentService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).CreateTask(ctx, req.(*admin.CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AsyncAgentService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTask(ctx, req.(*admin.GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.DeleteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AsyncAgentService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).DeleteTask(ctx, req.(*admin.DeleteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTaskMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTaskMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AsyncAgentService/GetTaskMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTaskMetrics(ctx, req.(*admin.GetTaskMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AsyncAgentService_GetTaskLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetTaskLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AsyncAgentServiceServer).GetTaskLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AsyncAgentService/GetTaskLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AsyncAgentServiceServer).GetTaskLogs(ctx, req.(*admin.GetTaskLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AsyncAgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AsyncAgentService",
	HandlerType: (*AsyncAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _AsyncAgentService_CreateTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _AsyncAgentService_GetTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _AsyncAgentService_DeleteTask_Handler,
		},
		{
			MethodName: "GetTaskMetrics",
			Handler:    _AsyncAgentService_GetTaskMetrics_Handler,
		},
		{
			MethodName: "GetTaskLogs",
			Handler:    _AsyncAgentService_GetTaskLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/agent.proto",
}

// AgentMetadataServiceClient is the client API for AgentMetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentMetadataServiceClient interface {
	// Fetch a :ref:`ref_flyteidl.admin.Agent` definition.
	GetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) (*admin.GetAgentResponse, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Agent` definitions.
	ListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) (*admin.ListAgentsResponse, error)
}

type agentMetadataServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentMetadataServiceClient(cc *grpc.ClientConn) AgentMetadataServiceClient {
	return &agentMetadataServiceClient{cc}
}

func (c *agentMetadataServiceClient) GetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) (*admin.GetAgentResponse, error) {
	out := new(admin.GetAgentResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AgentMetadataService/GetAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentMetadataServiceClient) ListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) (*admin.ListAgentsResponse, error) {
	out := new(admin.ListAgentsResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.AgentMetadataService/ListAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentMetadataServiceServer is the server API for AgentMetadataService service.
type AgentMetadataServiceServer interface {
	// Fetch a :ref:`ref_flyteidl.admin.Agent` definition.
	GetAgent(context.Context, *admin.GetAgentRequest) (*admin.GetAgentResponse, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Agent` definitions.
	ListAgents(context.Context, *admin.ListAgentsRequest) (*admin.ListAgentsResponse, error)
}

// UnimplementedAgentMetadataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentMetadataServiceServer struct {
}

func (*UnimplementedAgentMetadataServiceServer) GetAgent(ctx context.Context, req *admin.GetAgentRequest) (*admin.GetAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (*UnimplementedAgentMetadataServiceServer) ListAgents(ctx context.Context, req *admin.ListAgentsRequest) (*admin.ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgents not implemented")
}

func RegisterAgentMetadataServiceServer(s *grpc.Server, srv AgentMetadataServiceServer) {
	s.RegisterService(&_AgentMetadataService_serviceDesc, srv)
}

func _AgentMetadataService_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentMetadataServiceServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AgentMetadataService/GetAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentMetadataServiceServer).GetAgent(ctx, req.(*admin.GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentMetadataService_ListAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentMetadataServiceServer).ListAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.AgentMetadataService/ListAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentMetadataServiceServer).ListAgents(ctx, req.(*admin.ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentMetadataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.AgentMetadataService",
	HandlerType: (*AgentMetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgent",
			Handler:    _AgentMetadataService_GetAgent_Handler,
		},
		{
			MethodName: "ListAgents",
			Handler:    _AgentMetadataService_ListAgents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/agent.proto",
}