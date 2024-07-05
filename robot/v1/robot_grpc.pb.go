// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: robot/v1/robot.proto

package v1

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

// RobotServiceClient is the client API for RobotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotServiceClient interface {
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error)
	GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error)
	// ResourceNames returns the list of all resources.
	ResourceNames(ctx context.Context, in *ResourceNamesRequest, opts ...grpc.CallOption) (*ResourceNamesResponse, error)
	// ResourceRPCSubtypes returns the list of all resource types.
	ResourceRPCSubtypes(ctx context.Context, in *ResourceRPCSubtypesRequest, opts ...grpc.CallOption) (*ResourceRPCSubtypesResponse, error)
	CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*CancelOperationResponse, error)
	BlockForOperation(ctx context.Context, in *BlockForOperationRequest, opts ...grpc.CallOption) (*BlockForOperationResponse, error)
	// DiscoverComponents returns the list of discovered component configurations.
	DiscoverComponents(ctx context.Context, in *DiscoverComponentsRequest, opts ...grpc.CallOption) (*DiscoverComponentsResponse, error)
	FrameSystemConfig(ctx context.Context, in *FrameSystemConfigRequest, opts ...grpc.CallOption) (*FrameSystemConfigResponse, error)
	TransformPose(ctx context.Context, in *TransformPoseRequest, opts ...grpc.CallOption) (*TransformPoseResponse, error)
	TransformPCD(ctx context.Context, in *TransformPCDRequest, opts ...grpc.CallOption) (*TransformPCDResponse, error)
	// Deprecated: Do not use.
	// GetStatus returns the list of all statuses requested. An empty request signifies all resources.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	// Deprecated: Do not use.
	// StreamStatus periodically sends the status of all statuses requested. An empty request signifies all resources.
	StreamStatus(ctx context.Context, in *StreamStatusRequest, opts ...grpc.CallOption) (RobotService_StreamStatusClient, error)
	// StopAll will stop all current and outstanding operations for the robot and stops all actuators and movement
	StopAll(ctx context.Context, in *StopAllRequest, opts ...grpc.CallOption) (*StopAllResponse, error)
	// StartSession creates a new session that expects at least one heartbeat within the returned window.
	// If the window lapses, any resources that have safety heart monitored methods, where this session was
	// the last caller on the resource, will be stopped.
	StartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (*StartSessionResponse, error)
	// SendSessionHeartbeat sends a heartbeat to the given session. If the session has expired, a
	// SESSION_EXPIRED error will be returned.
	SendSessionHeartbeat(ctx context.Context, in *SendSessionHeartbeatRequest, opts ...grpc.CallOption) (*SendSessionHeartbeatResponse, error)
	// Log sends logs to be logged by this robot. Currently used for module logging.
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	// GetCloudMetadata returns app-related information about the robot.
	GetCloudMetadata(ctx context.Context, in *GetCloudMetadataRequest, opts ...grpc.CallOption) (*GetCloudMetadataResponse, error)
	RestartModule(ctx context.Context, in *RestartModuleRequest, opts ...grpc.CallOption) (*RestartModuleResponse, error)
	// Shutdown shuts down the robot.
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
	// GetMachineStatus returns the current status of the robot.
	GetMachineStatus(ctx context.Context, in *GetMachineStatusRequest, opts ...grpc.CallOption) (*GetMachineStatusResponse, error)
}

type robotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotServiceClient(cc grpc.ClientConnInterface) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error) {
	out := new(GetOperationsResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error) {
	out := new(GetSessionsResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/GetSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ResourceNames(ctx context.Context, in *ResourceNamesRequest, opts ...grpc.CallOption) (*ResourceNamesResponse, error) {
	out := new(ResourceNamesResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/ResourceNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ResourceRPCSubtypes(ctx context.Context, in *ResourceRPCSubtypesRequest, opts ...grpc.CallOption) (*ResourceRPCSubtypesResponse, error) {
	out := new(ResourceRPCSubtypesResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/ResourceRPCSubtypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) CancelOperation(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*CancelOperationResponse, error) {
	out := new(CancelOperationResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/CancelOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) BlockForOperation(ctx context.Context, in *BlockForOperationRequest, opts ...grpc.CallOption) (*BlockForOperationResponse, error) {
	out := new(BlockForOperationResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/BlockForOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) DiscoverComponents(ctx context.Context, in *DiscoverComponentsRequest, opts ...grpc.CallOption) (*DiscoverComponentsResponse, error) {
	out := new(DiscoverComponentsResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/DiscoverComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) FrameSystemConfig(ctx context.Context, in *FrameSystemConfigRequest, opts ...grpc.CallOption) (*FrameSystemConfigResponse, error) {
	out := new(FrameSystemConfigResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/FrameSystemConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) TransformPose(ctx context.Context, in *TransformPoseRequest, opts ...grpc.CallOption) (*TransformPoseResponse, error) {
	out := new(TransformPoseResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/TransformPose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) TransformPCD(ctx context.Context, in *TransformPCDRequest, opts ...grpc.CallOption) (*TransformPCDResponse, error) {
	out := new(TransformPCDResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/TransformPCD", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *robotServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *robotServiceClient) StreamStatus(ctx context.Context, in *StreamStatusRequest, opts ...grpc.CallOption) (RobotService_StreamStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &RobotService_ServiceDesc.Streams[0], "/viam.robot.v1.RobotService/StreamStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotServiceStreamStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotService_StreamStatusClient interface {
	Recv() (*StreamStatusResponse, error)
	grpc.ClientStream
}

type robotServiceStreamStatusClient struct {
	grpc.ClientStream
}

func (x *robotServiceStreamStatusClient) Recv() (*StreamStatusResponse, error) {
	m := new(StreamStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotServiceClient) StopAll(ctx context.Context, in *StopAllRequest, opts ...grpc.CallOption) (*StopAllResponse, error) {
	out := new(StopAllResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/StopAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) StartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (*StartSessionResponse, error) {
	out := new(StartSessionResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/StartSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) SendSessionHeartbeat(ctx context.Context, in *SendSessionHeartbeatRequest, opts ...grpc.CallOption) (*SendSessionHeartbeatResponse, error) {
	out := new(SendSessionHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/SendSessionHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) GetCloudMetadata(ctx context.Context, in *GetCloudMetadataRequest, opts ...grpc.CallOption) (*GetCloudMetadataResponse, error) {
	out := new(GetCloudMetadataResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/GetCloudMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) RestartModule(ctx context.Context, in *RestartModuleRequest, opts ...grpc.CallOption) (*RestartModuleResponse, error) {
	out := new(RestartModuleResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/RestartModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) GetMachineStatus(ctx context.Context, in *GetMachineStatusRequest, opts ...grpc.CallOption) (*GetMachineStatusResponse, error) {
	out := new(GetMachineStatusResponse)
	err := c.cc.Invoke(ctx, "/viam.robot.v1.RobotService/GetMachineStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServiceServer is the server API for RobotService service.
// All implementations must embed UnimplementedRobotServiceServer
// for forward compatibility
type RobotServiceServer interface {
	GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error)
	GetSessions(context.Context, *GetSessionsRequest) (*GetSessionsResponse, error)
	// ResourceNames returns the list of all resources.
	ResourceNames(context.Context, *ResourceNamesRequest) (*ResourceNamesResponse, error)
	// ResourceRPCSubtypes returns the list of all resource types.
	ResourceRPCSubtypes(context.Context, *ResourceRPCSubtypesRequest) (*ResourceRPCSubtypesResponse, error)
	CancelOperation(context.Context, *CancelOperationRequest) (*CancelOperationResponse, error)
	BlockForOperation(context.Context, *BlockForOperationRequest) (*BlockForOperationResponse, error)
	// DiscoverComponents returns the list of discovered component configurations.
	DiscoverComponents(context.Context, *DiscoverComponentsRequest) (*DiscoverComponentsResponse, error)
	FrameSystemConfig(context.Context, *FrameSystemConfigRequest) (*FrameSystemConfigResponse, error)
	TransformPose(context.Context, *TransformPoseRequest) (*TransformPoseResponse, error)
	TransformPCD(context.Context, *TransformPCDRequest) (*TransformPCDResponse, error)
	// Deprecated: Do not use.
	// GetStatus returns the list of all statuses requested. An empty request signifies all resources.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	// Deprecated: Do not use.
	// StreamStatus periodically sends the status of all statuses requested. An empty request signifies all resources.
	StreamStatus(*StreamStatusRequest, RobotService_StreamStatusServer) error
	// StopAll will stop all current and outstanding operations for the robot and stops all actuators and movement
	StopAll(context.Context, *StopAllRequest) (*StopAllResponse, error)
	// StartSession creates a new session that expects at least one heartbeat within the returned window.
	// If the window lapses, any resources that have safety heart monitored methods, where this session was
	// the last caller on the resource, will be stopped.
	StartSession(context.Context, *StartSessionRequest) (*StartSessionResponse, error)
	// SendSessionHeartbeat sends a heartbeat to the given session. If the session has expired, a
	// SESSION_EXPIRED error will be returned.
	SendSessionHeartbeat(context.Context, *SendSessionHeartbeatRequest) (*SendSessionHeartbeatResponse, error)
	// Log sends logs to be logged by this robot. Currently used for module logging.
	Log(context.Context, *LogRequest) (*LogResponse, error)
	// GetCloudMetadata returns app-related information about the robot.
	GetCloudMetadata(context.Context, *GetCloudMetadataRequest) (*GetCloudMetadataResponse, error)
	RestartModule(context.Context, *RestartModuleRequest) (*RestartModuleResponse, error)
	// Shutdown shuts down the robot.
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
	// GetMachineStatus returns the current status of the robot.
	GetMachineStatus(context.Context, *GetMachineStatusRequest) (*GetMachineStatusResponse, error)
	mustEmbedUnimplementedRobotServiceServer()
}

// UnimplementedRobotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotServiceServer struct {
}

func (UnimplementedRobotServiceServer) GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedRobotServiceServer) GetSessions(context.Context, *GetSessionsRequest) (*GetSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessions not implemented")
}
func (UnimplementedRobotServiceServer) ResourceNames(context.Context, *ResourceNamesRequest) (*ResourceNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceNames not implemented")
}
func (UnimplementedRobotServiceServer) ResourceRPCSubtypes(context.Context, *ResourceRPCSubtypesRequest) (*ResourceRPCSubtypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceRPCSubtypes not implemented")
}
func (UnimplementedRobotServiceServer) CancelOperation(context.Context, *CancelOperationRequest) (*CancelOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOperation not implemented")
}
func (UnimplementedRobotServiceServer) BlockForOperation(context.Context, *BlockForOperationRequest) (*BlockForOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockForOperation not implemented")
}
func (UnimplementedRobotServiceServer) DiscoverComponents(context.Context, *DiscoverComponentsRequest) (*DiscoverComponentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverComponents not implemented")
}
func (UnimplementedRobotServiceServer) FrameSystemConfig(context.Context, *FrameSystemConfigRequest) (*FrameSystemConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FrameSystemConfig not implemented")
}
func (UnimplementedRobotServiceServer) TransformPose(context.Context, *TransformPoseRequest) (*TransformPoseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformPose not implemented")
}
func (UnimplementedRobotServiceServer) TransformPCD(context.Context, *TransformPCDRequest) (*TransformPCDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformPCD not implemented")
}
func (UnimplementedRobotServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedRobotServiceServer) StreamStatus(*StreamStatusRequest, RobotService_StreamStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStatus not implemented")
}
func (UnimplementedRobotServiceServer) StopAll(context.Context, *StopAllRequest) (*StopAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAll not implemented")
}
func (UnimplementedRobotServiceServer) StartSession(context.Context, *StartSessionRequest) (*StartSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSession not implemented")
}
func (UnimplementedRobotServiceServer) SendSessionHeartbeat(context.Context, *SendSessionHeartbeatRequest) (*SendSessionHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSessionHeartbeat not implemented")
}
func (UnimplementedRobotServiceServer) Log(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedRobotServiceServer) GetCloudMetadata(context.Context, *GetCloudMetadataRequest) (*GetCloudMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCloudMetadata not implemented")
}
func (UnimplementedRobotServiceServer) RestartModule(context.Context, *RestartModuleRequest) (*RestartModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartModule not implemented")
}
func (UnimplementedRobotServiceServer) Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedRobotServiceServer) GetMachineStatus(context.Context, *GetMachineStatusRequest) (*GetMachineStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachineStatus not implemented")
}
func (UnimplementedRobotServiceServer) mustEmbedUnimplementedRobotServiceServer() {}

// UnsafeRobotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotServiceServer will
// result in compilation errors.
type UnsafeRobotServiceServer interface {
	mustEmbedUnimplementedRobotServiceServer()
}

func RegisterRobotServiceServer(s grpc.ServiceRegistrar, srv RobotServiceServer) {
	s.RegisterService(&RobotService_ServiceDesc, srv)
}

func _RobotService_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_GetSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/GetSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetSessions(ctx, req.(*GetSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ResourceNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ResourceNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/ResourceNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ResourceNames(ctx, req.(*ResourceNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ResourceRPCSubtypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceRPCSubtypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ResourceRPCSubtypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/ResourceRPCSubtypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ResourceRPCSubtypes(ctx, req.(*ResourceRPCSubtypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_CancelOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).CancelOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/CancelOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).CancelOperation(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_BlockForOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockForOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).BlockForOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/BlockForOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).BlockForOperation(ctx, req.(*BlockForOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_DiscoverComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).DiscoverComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/DiscoverComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).DiscoverComponents(ctx, req.(*DiscoverComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_FrameSystemConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FrameSystemConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).FrameSystemConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/FrameSystemConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).FrameSystemConfig(ctx, req.(*FrameSystemConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_TransformPose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformPoseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).TransformPose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/TransformPose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).TransformPose(ctx, req.(*TransformPoseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_TransformPCD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformPCDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).TransformPCD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/TransformPCD",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).TransformPCD(ctx, req.(*TransformPCDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_StreamStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotServiceServer).StreamStatus(m, &robotServiceStreamStatusServer{stream})
}

type RobotService_StreamStatusServer interface {
	Send(*StreamStatusResponse) error
	grpc.ServerStream
}

type robotServiceStreamStatusServer struct {
	grpc.ServerStream
}

func (x *robotServiceStreamStatusServer) Send(m *StreamStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RobotService_StopAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).StopAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/StopAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).StopAll(ctx, req.(*StopAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_StartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).StartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/StartSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).StartSession(ctx, req.(*StartSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_SendSessionHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSessionHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).SendSessionHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/SendSessionHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).SendSessionHeartbeat(ctx, req.(*SendSessionHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_GetCloudMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetCloudMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/GetCloudMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetCloudMetadata(ctx, req.(*GetCloudMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_RestartModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).RestartModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/RestartModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).RestartModule(ctx, req.(*RestartModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_GetMachineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMachineStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).GetMachineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.robot.v1.RobotService/GetMachineStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).GetMachineStatus(ctx, req.(*GetMachineStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotService_ServiceDesc is the grpc.ServiceDesc for RobotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.robot.v1.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperations",
			Handler:    _RobotService_GetOperations_Handler,
		},
		{
			MethodName: "GetSessions",
			Handler:    _RobotService_GetSessions_Handler,
		},
		{
			MethodName: "ResourceNames",
			Handler:    _RobotService_ResourceNames_Handler,
		},
		{
			MethodName: "ResourceRPCSubtypes",
			Handler:    _RobotService_ResourceRPCSubtypes_Handler,
		},
		{
			MethodName: "CancelOperation",
			Handler:    _RobotService_CancelOperation_Handler,
		},
		{
			MethodName: "BlockForOperation",
			Handler:    _RobotService_BlockForOperation_Handler,
		},
		{
			MethodName: "DiscoverComponents",
			Handler:    _RobotService_DiscoverComponents_Handler,
		},
		{
			MethodName: "FrameSystemConfig",
			Handler:    _RobotService_FrameSystemConfig_Handler,
		},
		{
			MethodName: "TransformPose",
			Handler:    _RobotService_TransformPose_Handler,
		},
		{
			MethodName: "TransformPCD",
			Handler:    _RobotService_TransformPCD_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _RobotService_GetStatus_Handler,
		},
		{
			MethodName: "StopAll",
			Handler:    _RobotService_StopAll_Handler,
		},
		{
			MethodName: "StartSession",
			Handler:    _RobotService_StartSession_Handler,
		},
		{
			MethodName: "SendSessionHeartbeat",
			Handler:    _RobotService_SendSessionHeartbeat_Handler,
		},
		{
			MethodName: "Log",
			Handler:    _RobotService_Log_Handler,
		},
		{
			MethodName: "GetCloudMetadata",
			Handler:    _RobotService_GetCloudMetadata_Handler,
		},
		{
			MethodName: "RestartModule",
			Handler:    _RobotService_RestartModule_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _RobotService_Shutdown_Handler,
		},
		{
			MethodName: "GetMachineStatus",
			Handler:    _RobotService_GetMachineStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamStatus",
			Handler:       _RobotService_StreamStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "robot/v1/robot.proto",
}
