// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: task.proto

package task

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
	Task_CreateCuratorialTask_FullMethodName    = "/task.Task/CreateCuratorialTask"
	Task_QueryTaskList_FullMethodName           = "/task.Task/QueryTaskList"
	Task_QueryTaskDetails_FullMethodName        = "/task.Task/QueryTaskDetails"
	Task_QueryUserLaunchTaskList_FullMethodName = "/task.Task/QueryUserLaunchTaskList"
	Task_CreateLabel_FullMethodName             = "/task.Task/CreateLabel"
	Task_DeleteLabel_FullMethodName             = "/task.Task/DeleteLabel"
	Task_QueryLabelList_FullMethodName          = "/task.Task/QueryLabelList"
	Task_PerformTask_FullMethodName             = "/task.Task/PerformTask"
	Task_VoluntarilyTaskSchedule_FullMethodName = "/task.Task/VoluntarilyTaskSchedule"
	Task_AmendTreasureTask_FullMethodName       = "/task.Task/AmendTreasureTask"
	Task_ChangeTreasureTask_FullMethodName      = "/task.Task/ChangeTreasureTask"
	Task_QueryTreasureTaskList_FullMethodName   = "/task.Task/QueryTreasureTaskList"
	Task_QuerySubtaskStyle_FullMethodName       = "/task.Task/QuerySubtaskStyle"
	Task_AmendAssociatedSubtask_FullMethodName  = "/task.Task/AmendAssociatedSubtask"
	Task_DeleteAssociatedSubtask_FullMethodName = "/task.Task/DeleteAssociatedSubtask"
	Task_QueryAssociatedSubtask_FullMethodName  = "/task.Task/QueryAssociatedSubtask"
	Task_AmendChestCollection_FullMethodName    = "/task.Task/AmendChestCollection"
	Task_QueryChestCollection_FullMethodName    = "/task.Task/QueryChestCollection"
	Task_CreateUserPowerTask_FullMethodName     = "/task.Task/CreateUserPowerTask"
	Task_CreateSubtaskStyle_FullMethodName      = "/task.Task/CreateSubtaskStyle"
	Task_CreateAssistanceTask_FullMethodName    = "/task.Task/CreateAssistanceTask"
	Task_QueryAssistanceTask_FullMethodName     = "/task.Task/QueryAssistanceTask"
	Task_Ping_FullMethodName                    = "/task.Task/ping"
)

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	// 策展任务相关
	CreateCuratorialTask(ctx context.Context, in *CreatePublishTaskInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryTaskList(ctx context.Context, in *PublishTaskInput, opts ...grpc.CallOption) (*RePublishTask, error)
	QueryTaskDetails(ctx context.Context, in *TaskDetailsInput, opts ...grpc.CallOption) (*ReTaskDetails, error)
	QueryUserLaunchTaskList(ctx context.Context, in *UserLaunchTaskListInput, opts ...grpc.CallOption) (*RePublishTask, error)
	CreateLabel(ctx context.Context, in *CreateLabelInput, opts ...grpc.CallOption) (*Mistake, error)
	DeleteLabel(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryLabelList(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReLabelListOut, error)
	PerformTask(ctx context.Context, in *PerformTaskInput, opts ...grpc.CallOption) (*Mistake, error)
	VoluntarilyTaskSchedule(ctx context.Context, in *VoluntarilyTaskScheduleInput, opts ...grpc.CallOption) (*Mistake, error)
	// 每日任务
	AmendTreasureTask(ctx context.Context, in *TreasureTaskSrtInput, opts ...grpc.CallOption) (*Mistake, error)
	ChangeTreasureTask(ctx context.Context, in *TreasureTaskInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryTreasureTaskList(ctx context.Context, in *TreasureTaskListInput, opts ...grpc.CallOption) (*ReTreasureTaskSrt, error)
	QuerySubtaskStyle(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReSubtaskStyle, error)
	AmendAssociatedSubtask(ctx context.Context, in *AssociatedSubtaskSrt, opts ...grpc.CallOption) (*Mistake, error)
	DeleteAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReAssociatedSubtask, error)
	AmendChestCollection(ctx context.Context, in *AmendChestCollectionInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryChestCollection(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReChestCollectionSrt, error)
	CreateUserPowerTask(ctx context.Context, in *CreateUserPowerTaskInput, opts ...grpc.CallOption) (*Mistake, error)
	CreateSubtaskStyle(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*Mistake, error)
	CreateAssistanceTask(ctx context.Context, in *CreateUserPublishingAssistanceTaskInput, opts ...grpc.CallOption) (*Mistake, error)
	QueryAssistanceTask(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*UserPublishingAssistanceTask, error)
	Ping(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) CreateCuratorialTask(ctx context.Context, in *CreatePublishTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_CreateCuratorialTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryTaskList(ctx context.Context, in *PublishTaskInput, opts ...grpc.CallOption) (*RePublishTask, error) {
	out := new(RePublishTask)
	err := c.cc.Invoke(ctx, Task_QueryTaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryTaskDetails(ctx context.Context, in *TaskDetailsInput, opts ...grpc.CallOption) (*ReTaskDetails, error) {
	out := new(ReTaskDetails)
	err := c.cc.Invoke(ctx, Task_QueryTaskDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryUserLaunchTaskList(ctx context.Context, in *UserLaunchTaskListInput, opts ...grpc.CallOption) (*RePublishTask, error) {
	out := new(RePublishTask)
	err := c.cc.Invoke(ctx, Task_QueryUserLaunchTaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) CreateLabel(ctx context.Context, in *CreateLabelInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_CreateLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) DeleteLabel(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_DeleteLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryLabelList(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReLabelListOut, error) {
	out := new(ReLabelListOut)
	err := c.cc.Invoke(ctx, Task_QueryLabelList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) PerformTask(ctx context.Context, in *PerformTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_PerformTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) VoluntarilyTaskSchedule(ctx context.Context, in *VoluntarilyTaskScheduleInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_VoluntarilyTaskSchedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) AmendTreasureTask(ctx context.Context, in *TreasureTaskSrtInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_AmendTreasureTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) ChangeTreasureTask(ctx context.Context, in *TreasureTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_ChangeTreasureTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryTreasureTaskList(ctx context.Context, in *TreasureTaskListInput, opts ...grpc.CallOption) (*ReTreasureTaskSrt, error) {
	out := new(ReTreasureTaskSrt)
	err := c.cc.Invoke(ctx, Task_QueryTreasureTaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QuerySubtaskStyle(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReSubtaskStyle, error) {
	out := new(ReSubtaskStyle)
	err := c.cc.Invoke(ctx, Task_QuerySubtaskStyle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) AmendAssociatedSubtask(ctx context.Context, in *AssociatedSubtaskSrt, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_AmendAssociatedSubtask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) DeleteAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_DeleteAssociatedSubtask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReAssociatedSubtask, error) {
	out := new(ReAssociatedSubtask)
	err := c.cc.Invoke(ctx, Task_QueryAssociatedSubtask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) AmendChestCollection(ctx context.Context, in *AmendChestCollectionInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_AmendChestCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryChestCollection(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReChestCollectionSrt, error) {
	out := new(ReChestCollectionSrt)
	err := c.cc.Invoke(ctx, Task_QueryChestCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) CreateUserPowerTask(ctx context.Context, in *CreateUserPowerTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_CreateUserPowerTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) CreateSubtaskStyle(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_CreateSubtaskStyle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) CreateAssistanceTask(ctx context.Context, in *CreateUserPublishingAssistanceTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_CreateAssistanceTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) QueryAssistanceTask(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*UserPublishingAssistanceTask, error) {
	out := new(UserPublishingAssistanceTask)
	err := c.cc.Invoke(ctx, Task_QueryAssistanceTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) Ping(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	out := new(Mistake)
	err := c.cc.Invoke(ctx, Task_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	// 策展任务相关
	CreateCuratorialTask(context.Context, *CreatePublishTaskInput) (*Mistake, error)
	QueryTaskList(context.Context, *PublishTaskInput) (*RePublishTask, error)
	QueryTaskDetails(context.Context, *TaskDetailsInput) (*ReTaskDetails, error)
	QueryUserLaunchTaskList(context.Context, *UserLaunchTaskListInput) (*RePublishTask, error)
	CreateLabel(context.Context, *CreateLabelInput) (*Mistake, error)
	DeleteLabel(context.Context, *TaskIDInquireInput) (*Mistake, error)
	QueryLabelList(context.Context, *UserIDInquireInput) (*ReLabelListOut, error)
	PerformTask(context.Context, *PerformTaskInput) (*Mistake, error)
	VoluntarilyTaskSchedule(context.Context, *VoluntarilyTaskScheduleInput) (*Mistake, error)
	// 每日任务
	AmendTreasureTask(context.Context, *TreasureTaskSrtInput) (*Mistake, error)
	ChangeTreasureTask(context.Context, *TreasureTaskInput) (*Mistake, error)
	QueryTreasureTaskList(context.Context, *TreasureTaskListInput) (*ReTreasureTaskSrt, error)
	QuerySubtaskStyle(context.Context, *TaskIDInquireInput) (*ReSubtaskStyle, error)
	AmendAssociatedSubtask(context.Context, *AssociatedSubtaskSrt) (*Mistake, error)
	DeleteAssociatedSubtask(context.Context, *TaskIDInquireInput) (*Mistake, error)
	QueryAssociatedSubtask(context.Context, *TaskIDInquireInput) (*ReAssociatedSubtask, error)
	AmendChestCollection(context.Context, *AmendChestCollectionInput) (*Mistake, error)
	QueryChestCollection(context.Context, *UserIDInquireInput) (*ReChestCollectionSrt, error)
	CreateUserPowerTask(context.Context, *CreateUserPowerTaskInput) (*Mistake, error)
	CreateSubtaskStyle(context.Context, *UserIDInquireInput) (*Mistake, error)
	CreateAssistanceTask(context.Context, *CreateUserPublishingAssistanceTaskInput) (*Mistake, error)
	QueryAssistanceTask(context.Context, *UserIDInquireInput) (*UserPublishingAssistanceTask, error)
	Ping(context.Context, *TaskIDInquireInput) (*Mistake, error)
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) CreateCuratorialTask(context.Context, *CreatePublishTaskInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCuratorialTask not implemented")
}
func (UnimplementedTaskServer) QueryTaskList(context.Context, *PublishTaskInput) (*RePublishTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTaskList not implemented")
}
func (UnimplementedTaskServer) QueryTaskDetails(context.Context, *TaskDetailsInput) (*ReTaskDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTaskDetails not implemented")
}
func (UnimplementedTaskServer) QueryUserLaunchTaskList(context.Context, *UserLaunchTaskListInput) (*RePublishTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserLaunchTaskList not implemented")
}
func (UnimplementedTaskServer) CreateLabel(context.Context, *CreateLabelInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabel not implemented")
}
func (UnimplementedTaskServer) DeleteLabel(context.Context, *TaskIDInquireInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLabel not implemented")
}
func (UnimplementedTaskServer) QueryLabelList(context.Context, *UserIDInquireInput) (*ReLabelListOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLabelList not implemented")
}
func (UnimplementedTaskServer) PerformTask(context.Context, *PerformTaskInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformTask not implemented")
}
func (UnimplementedTaskServer) VoluntarilyTaskSchedule(context.Context, *VoluntarilyTaskScheduleInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoluntarilyTaskSchedule not implemented")
}
func (UnimplementedTaskServer) AmendTreasureTask(context.Context, *TreasureTaskSrtInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmendTreasureTask not implemented")
}
func (UnimplementedTaskServer) ChangeTreasureTask(context.Context, *TreasureTaskInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTreasureTask not implemented")
}
func (UnimplementedTaskServer) QueryTreasureTaskList(context.Context, *TreasureTaskListInput) (*ReTreasureTaskSrt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTreasureTaskList not implemented")
}
func (UnimplementedTaskServer) QuerySubtaskStyle(context.Context, *TaskIDInquireInput) (*ReSubtaskStyle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySubtaskStyle not implemented")
}
func (UnimplementedTaskServer) AmendAssociatedSubtask(context.Context, *AssociatedSubtaskSrt) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmendAssociatedSubtask not implemented")
}
func (UnimplementedTaskServer) DeleteAssociatedSubtask(context.Context, *TaskIDInquireInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssociatedSubtask not implemented")
}
func (UnimplementedTaskServer) QueryAssociatedSubtask(context.Context, *TaskIDInquireInput) (*ReAssociatedSubtask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAssociatedSubtask not implemented")
}
func (UnimplementedTaskServer) AmendChestCollection(context.Context, *AmendChestCollectionInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AmendChestCollection not implemented")
}
func (UnimplementedTaskServer) QueryChestCollection(context.Context, *UserIDInquireInput) (*ReChestCollectionSrt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryChestCollection not implemented")
}
func (UnimplementedTaskServer) CreateUserPowerTask(context.Context, *CreateUserPowerTaskInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPowerTask not implemented")
}
func (UnimplementedTaskServer) CreateSubtaskStyle(context.Context, *UserIDInquireInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubtaskStyle not implemented")
}
func (UnimplementedTaskServer) CreateAssistanceTask(context.Context, *CreateUserPublishingAssistanceTaskInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssistanceTask not implemented")
}
func (UnimplementedTaskServer) QueryAssistanceTask(context.Context, *UserIDInquireInput) (*UserPublishingAssistanceTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAssistanceTask not implemented")
}
func (UnimplementedTaskServer) Ping(context.Context, *TaskIDInquireInput) (*Mistake, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_CreateCuratorialTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublishTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateCuratorialTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateCuratorialTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateCuratorialTask(ctx, req.(*CreatePublishTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryTaskList(ctx, req.(*PublishTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryTaskDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskDetailsInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryTaskDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryTaskDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryTaskDetails(ctx, req.(*TaskDetailsInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryUserLaunchTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLaunchTaskListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryUserLaunchTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryUserLaunchTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryUserLaunchTaskList(ctx, req.(*UserLaunchTaskListInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_CreateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateLabel(ctx, req.(*CreateLabelInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_DeleteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).DeleteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_DeleteLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).DeleteLabel(ctx, req.(*TaskIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryLabelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryLabelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryLabelList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryLabelList(ctx, req.(*UserIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_PerformTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).PerformTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_PerformTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).PerformTask(ctx, req.(*PerformTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_VoluntarilyTaskSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoluntarilyTaskScheduleInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).VoluntarilyTaskSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_VoluntarilyTaskSchedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).VoluntarilyTaskSchedule(ctx, req.(*VoluntarilyTaskScheduleInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_AmendTreasureTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TreasureTaskSrtInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).AmendTreasureTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_AmendTreasureTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).AmendTreasureTask(ctx, req.(*TreasureTaskSrtInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_ChangeTreasureTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TreasureTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).ChangeTreasureTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_ChangeTreasureTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).ChangeTreasureTask(ctx, req.(*TreasureTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryTreasureTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TreasureTaskListInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryTreasureTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryTreasureTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryTreasureTaskList(ctx, req.(*TreasureTaskListInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QuerySubtaskStyle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QuerySubtaskStyle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QuerySubtaskStyle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QuerySubtaskStyle(ctx, req.(*TaskIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_AmendAssociatedSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssociatedSubtaskSrt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).AmendAssociatedSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_AmendAssociatedSubtask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).AmendAssociatedSubtask(ctx, req.(*AssociatedSubtaskSrt))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_DeleteAssociatedSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).DeleteAssociatedSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_DeleteAssociatedSubtask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).DeleteAssociatedSubtask(ctx, req.(*TaskIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryAssociatedSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryAssociatedSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryAssociatedSubtask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryAssociatedSubtask(ctx, req.(*TaskIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_AmendChestCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmendChestCollectionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).AmendChestCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_AmendChestCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).AmendChestCollection(ctx, req.(*AmendChestCollectionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryChestCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryChestCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryChestCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryChestCollection(ctx, req.(*UserIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_CreateUserPowerTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPowerTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateUserPowerTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateUserPowerTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateUserPowerTask(ctx, req.(*CreateUserPowerTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_CreateSubtaskStyle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateSubtaskStyle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateSubtaskStyle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateSubtaskStyle(ctx, req.(*UserIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_CreateAssistanceTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserPublishingAssistanceTaskInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).CreateAssistanceTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_CreateAssistanceTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).CreateAssistanceTask(ctx, req.(*CreateUserPublishingAssistanceTaskInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_QueryAssistanceTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).QueryAssistanceTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_QueryAssistanceTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).QueryAssistanceTask(ctx, req.(*UserIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskIDInquireInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Task_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).Ping(ctx, req.(*TaskIDInquireInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCuratorialTask",
			Handler:    _Task_CreateCuratorialTask_Handler,
		},
		{
			MethodName: "QueryTaskList",
			Handler:    _Task_QueryTaskList_Handler,
		},
		{
			MethodName: "QueryTaskDetails",
			Handler:    _Task_QueryTaskDetails_Handler,
		},
		{
			MethodName: "QueryUserLaunchTaskList",
			Handler:    _Task_QueryUserLaunchTaskList_Handler,
		},
		{
			MethodName: "CreateLabel",
			Handler:    _Task_CreateLabel_Handler,
		},
		{
			MethodName: "DeleteLabel",
			Handler:    _Task_DeleteLabel_Handler,
		},
		{
			MethodName: "QueryLabelList",
			Handler:    _Task_QueryLabelList_Handler,
		},
		{
			MethodName: "PerformTask",
			Handler:    _Task_PerformTask_Handler,
		},
		{
			MethodName: "VoluntarilyTaskSchedule",
			Handler:    _Task_VoluntarilyTaskSchedule_Handler,
		},
		{
			MethodName: "AmendTreasureTask",
			Handler:    _Task_AmendTreasureTask_Handler,
		},
		{
			MethodName: "ChangeTreasureTask",
			Handler:    _Task_ChangeTreasureTask_Handler,
		},
		{
			MethodName: "QueryTreasureTaskList",
			Handler:    _Task_QueryTreasureTaskList_Handler,
		},
		{
			MethodName: "QuerySubtaskStyle",
			Handler:    _Task_QuerySubtaskStyle_Handler,
		},
		{
			MethodName: "AmendAssociatedSubtask",
			Handler:    _Task_AmendAssociatedSubtask_Handler,
		},
		{
			MethodName: "DeleteAssociatedSubtask",
			Handler:    _Task_DeleteAssociatedSubtask_Handler,
		},
		{
			MethodName: "QueryAssociatedSubtask",
			Handler:    _Task_QueryAssociatedSubtask_Handler,
		},
		{
			MethodName: "AmendChestCollection",
			Handler:    _Task_AmendChestCollection_Handler,
		},
		{
			MethodName: "QueryChestCollection",
			Handler:    _Task_QueryChestCollection_Handler,
		},
		{
			MethodName: "CreateUserPowerTask",
			Handler:    _Task_CreateUserPowerTask_Handler,
		},
		{
			MethodName: "CreateSubtaskStyle",
			Handler:    _Task_CreateSubtaskStyle_Handler,
		},
		{
			MethodName: "CreateAssistanceTask",
			Handler:    _Task_CreateAssistanceTask_Handler,
		},
		{
			MethodName: "QueryAssistanceTask",
			Handler:    _Task_QueryAssistanceTask_Handler,
		},
		{
			MethodName: "ping",
			Handler:    _Task_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}
