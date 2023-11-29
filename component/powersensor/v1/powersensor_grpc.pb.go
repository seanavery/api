// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: component/powersensor/v1/powersensor.proto

package v1

import (
	context "context"
	v1 "go.viam.com/api/common/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PowerSensorServiceClient is the client API for PowerSensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PowerSensorServiceClient interface {
	// GetVoltage returns the voltage reading of a power sensor in volts
	GetVoltage(ctx context.Context, in *GetVoltageRequest, opts ...grpc.CallOption) (*GetVoltageResponse, error)
	// GetCurrent returns the current reading of a power sensor in amperes
	GetCurrent(ctx context.Context, in *GetCurrentRequest, opts ...grpc.CallOption) (*GetCurrentResponse, error)
	// GetPower returns the power reading of a power sensor in watts
	GetPower(ctx context.Context, in *GetPowerRequest, opts ...grpc.CallOption) (*GetPowerResponse, error)
	// GetReadings returns the readings of a sensor of the underlying robot.
	GetReadings(ctx context.Context, in *v1.GetReadingsRequest, opts ...grpc.CallOption) (*v1.GetReadingsResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error)
}

type powerSensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPowerSensorServiceClient(cc grpc.ClientConnInterface) PowerSensorServiceClient {
	return &powerSensorServiceClient{cc}
}

func (c *powerSensorServiceClient) GetVoltage(ctx context.Context, in *GetVoltageRequest, opts ...grpc.CallOption) (*GetVoltageResponse, error) {
	out := new(GetVoltageResponse)
	err := c.cc.Invoke(ctx, "/viam.component.powersensor.v1.PowerSensorService/GetVoltage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerSensorServiceClient) GetCurrent(ctx context.Context, in *GetCurrentRequest, opts ...grpc.CallOption) (*GetCurrentResponse, error) {
	out := new(GetCurrentResponse)
	err := c.cc.Invoke(ctx, "/viam.component.powersensor.v1.PowerSensorService/GetCurrent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerSensorServiceClient) GetPower(ctx context.Context, in *GetPowerRequest, opts ...grpc.CallOption) (*GetPowerResponse, error) {
	out := new(GetPowerResponse)
	err := c.cc.Invoke(ctx, "/viam.component.powersensor.v1.PowerSensorService/GetPower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerSensorServiceClient) GetReadings(ctx context.Context, in *v1.GetReadingsRequest, opts ...grpc.CallOption) (*v1.GetReadingsResponse, error) {
	out := new(v1.GetReadingsResponse)
	err := c.cc.Invoke(ctx, "/viam.component.powersensor.v1.PowerSensorService/GetReadings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *powerSensorServiceClient) DoCommand(ctx context.Context, in *v1.DoCommandRequest, opts ...grpc.CallOption) (*v1.DoCommandResponse, error) {
	out := new(v1.DoCommandResponse)
	err := c.cc.Invoke(ctx, "/viam.component.powersensor.v1.PowerSensorService/DoCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PowerSensorServiceServer is the server API for PowerSensorService service.
// All implementations must embed UnimplementedPowerSensorServiceServer
// for forward compatibility
type PowerSensorServiceServer interface {
	// GetVoltage returns the voltage reading of a power sensor in volts
	GetVoltage(context.Context, *GetVoltageRequest) (*GetVoltageResponse, error)
	// GetCurrent returns the current reading of a power sensor in amperes
	GetCurrent(context.Context, *GetCurrentRequest) (*GetCurrentResponse, error)
	// GetPower returns the power reading of a power sensor in watts
	GetPower(context.Context, *GetPowerRequest) (*GetPowerResponse, error)
	// GetReadings returns the readings of a sensor of the underlying robot.
	GetReadings(context.Context, *v1.GetReadingsRequest) (*v1.GetReadingsResponse, error)
	// DoCommand sends/receives arbitrary commands
	DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error)
	mustEmbedUnimplementedPowerSensorServiceServer()
}

// UnimplementedPowerSensorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPowerSensorServiceServer struct {
}

func (UnimplementedPowerSensorServiceServer) GetVoltage(context.Context, *GetVoltageRequest) (*GetVoltageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoltage not implemented")
}
func (UnimplementedPowerSensorServiceServer) GetCurrent(context.Context, *GetCurrentRequest) (*GetCurrentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrent not implemented")
}
func (UnimplementedPowerSensorServiceServer) GetPower(context.Context, *GetPowerRequest) (*GetPowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPower not implemented")
}
func (UnimplementedPowerSensorServiceServer) GetReadings(context.Context, *v1.GetReadingsRequest) (*v1.GetReadingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadings not implemented")
}
func (UnimplementedPowerSensorServiceServer) DoCommand(context.Context, *v1.DoCommandRequest) (*v1.DoCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCommand not implemented")
}
func (UnimplementedPowerSensorServiceServer) mustEmbedUnimplementedPowerSensorServiceServer() {}

// UnsafePowerSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PowerSensorServiceServer will
// result in compilation errors.
type UnsafePowerSensorServiceServer interface {
	mustEmbedUnimplementedPowerSensorServiceServer()
}

func RegisterPowerSensorServiceServer(s grpc.ServiceRegistrar, srv PowerSensorServiceServer) {
	s.RegisterService(&PowerSensorService_ServiceDesc, srv)
}

func _PowerSensorService_GetVoltage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVoltageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerSensorServiceServer).GetVoltage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.powersensor.v1.PowerSensorService/GetVoltage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerSensorServiceServer).GetVoltage(ctx, req.(*GetVoltageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerSensorService_GetCurrent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerSensorServiceServer).GetCurrent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.powersensor.v1.PowerSensorService/GetCurrent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerSensorServiceServer).GetCurrent(ctx, req.(*GetCurrentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerSensorService_GetPower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerSensorServiceServer).GetPower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.powersensor.v1.PowerSensorService/GetPower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerSensorServiceServer).GetPower(ctx, req.(*GetPowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerSensorService_GetReadings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetReadingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerSensorServiceServer).GetReadings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.powersensor.v1.PowerSensorService/GetReadings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerSensorServiceServer).GetReadings(ctx, req.(*v1.GetReadingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PowerSensorService_DoCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DoCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PowerSensorServiceServer).DoCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.powersensor.v1.PowerSensorService/DoCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PowerSensorServiceServer).DoCommand(ctx, req.(*v1.DoCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PowerSensorService_ServiceDesc is the grpc.ServiceDesc for PowerSensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PowerSensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.component.powersensor.v1.PowerSensorService",
	HandlerType: (*PowerSensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVoltage",
			Handler:    _PowerSensorService_GetVoltage_Handler,
		},
		{
			MethodName: "GetCurrent",
			Handler:    _PowerSensorService_GetCurrent_Handler,
		},
		{
			MethodName: "GetPower",
			Handler:    _PowerSensorService_GetPower_Handler,
		},
		{
			MethodName: "GetReadings",
			Handler:    _PowerSensorService_GetReadings_Handler,
		},
		{
			MethodName: "DoCommand",
			Handler:    _PowerSensorService_DoCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "component/powersensor/v1/powersensor.proto",
}