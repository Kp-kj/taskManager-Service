// Code generated by goctl. DO NOT EDIT.
// Source: task.proto

package taskclient

import (
	"context"

	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AmendChestCollectionInput    = task.AmendChestCollectionInput
	AssociatedSubtask            = task.AssociatedSubtask
	AssociatedSubtaskSeed        = task.AssociatedSubtaskSeed
	AssociatedSubtaskSrt         = task.AssociatedSubtaskSrt
	CreateLabelInput             = task.CreateLabelInput
	CreatePublishTaskInput       = task.CreatePublishTaskInput
	CreateUserPowerTaskInput     = task.CreateUserPowerTaskInput
	DetermineWhetherTaskComplete = task.DetermineWhetherTaskComplete
	Mistake                      = task.Mistake
	PaginationData               = task.PaginationData
	ParticipantBak               = task.ParticipantBak
	PerformTaskInput             = task.PerformTaskInput
	PublishTaskInput             = task.PublishTaskInput
	ReAssociatedSubtask          = task.ReAssociatedSubtask
	ReChestCollectionSrt         = task.ReChestCollectionSrt
	ReLabelList                  = task.ReLabelList
	ReLabelListOut               = task.ReLabelListOut
	RePublishTask                = task.RePublishTask
	RePublishTaskBak             = task.RePublishTaskBak
	ReSubtaskStyle               = task.ReSubtaskStyle
	ReTaskDetails                = task.ReTaskDetails
	ReTreasureTaskSrt            = task.ReTreasureTaskSrt
	SubtaskStyle                 = task.SubtaskStyle
	TaskDemand                   = task.TaskDemand
	TaskDemandBak                = task.TaskDemandBak
	TaskDetailsInput             = task.TaskDetailsInput
	TaskIDInquireInput           = task.TaskIDInquireInput
	TreasureTaskInput            = task.TreasureTaskInput
	TreasureTaskListInput        = task.TreasureTaskListInput
	TreasureTaskSrtInput         = task.TreasureTaskSrtInput
	TreasureTaskStage            = task.TreasureTaskStage
	TreasureTaskStageSeed        = task.TreasureTaskStageSeed
	UserIDInquireInput           = task.UserIDInquireInput
	UserLaunchTaskListInput      = task.UserLaunchTaskListInput
	VoluntarilyTaskScheduleInput = task.VoluntarilyTaskScheduleInput

	Task interface {
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
	}

	defaultTask struct {
		cli zrpc.Client
	}
)

func NewTask(cli zrpc.Client) Task {
	return &defaultTask{
		cli: cli,
	}
}

// 策展任务相关
func (m *defaultTask) CreateCuratorialTask(ctx context.Context, in *CreatePublishTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.CreateCuratorialTask(ctx, in, opts...)
}

func (m *defaultTask) QueryTaskList(ctx context.Context, in *PublishTaskInput, opts ...grpc.CallOption) (*RePublishTask, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryTaskList(ctx, in, opts...)
}

func (m *defaultTask) QueryTaskDetails(ctx context.Context, in *TaskDetailsInput, opts ...grpc.CallOption) (*ReTaskDetails, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryTaskDetails(ctx, in, opts...)
}

func (m *defaultTask) QueryUserLaunchTaskList(ctx context.Context, in *UserLaunchTaskListInput, opts ...grpc.CallOption) (*RePublishTask, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryUserLaunchTaskList(ctx, in, opts...)
}

func (m *defaultTask) CreateLabel(ctx context.Context, in *CreateLabelInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.CreateLabel(ctx, in, opts...)
}

func (m *defaultTask) DeleteLabel(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.DeleteLabel(ctx, in, opts...)
}

func (m *defaultTask) QueryLabelList(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReLabelListOut, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryLabelList(ctx, in, opts...)
}

func (m *defaultTask) PerformTask(ctx context.Context, in *PerformTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.PerformTask(ctx, in, opts...)
}

func (m *defaultTask) VoluntarilyTaskSchedule(ctx context.Context, in *VoluntarilyTaskScheduleInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.VoluntarilyTaskSchedule(ctx, in, opts...)
}

// 每日任务
func (m *defaultTask) AmendTreasureTask(ctx context.Context, in *TreasureTaskSrtInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.AmendTreasureTask(ctx, in, opts...)
}

func (m *defaultTask) ChangeTreasureTask(ctx context.Context, in *TreasureTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.ChangeTreasureTask(ctx, in, opts...)
}

func (m *defaultTask) QueryTreasureTaskList(ctx context.Context, in *TreasureTaskListInput, opts ...grpc.CallOption) (*ReTreasureTaskSrt, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryTreasureTaskList(ctx, in, opts...)
}

func (m *defaultTask) QuerySubtaskStyle(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReSubtaskStyle, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QuerySubtaskStyle(ctx, in, opts...)
}

func (m *defaultTask) AmendAssociatedSubtask(ctx context.Context, in *AssociatedSubtaskSrt, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.AmendAssociatedSubtask(ctx, in, opts...)
}

func (m *defaultTask) DeleteAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.DeleteAssociatedSubtask(ctx, in, opts...)
}

func (m *defaultTask) QueryAssociatedSubtask(ctx context.Context, in *TaskIDInquireInput, opts ...grpc.CallOption) (*ReAssociatedSubtask, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryAssociatedSubtask(ctx, in, opts...)
}

func (m *defaultTask) AmendChestCollection(ctx context.Context, in *AmendChestCollectionInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.AmendChestCollection(ctx, in, opts...)
}

func (m *defaultTask) QueryChestCollection(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*ReChestCollectionSrt, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.QueryChestCollection(ctx, in, opts...)
}

func (m *defaultTask) CreateUserPowerTask(ctx context.Context, in *CreateUserPowerTaskInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.CreateUserPowerTask(ctx, in, opts...)
}

func (m *defaultTask) CreateSubtaskStyle(ctx context.Context, in *UserIDInquireInput, opts ...grpc.CallOption) (*Mistake, error) {
	client := task.NewTaskClient(m.cli.Conn())
	return client.CreateSubtaskStyle(ctx, in, opts...)
}
