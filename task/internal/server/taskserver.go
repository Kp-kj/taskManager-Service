// Code generated by goctl. DO NOT EDIT.
// Source: task.proto

package server

import (
	"context"

	"taskManager-Service-main/task/internal/logic"
	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"
)

type TaskServer struct {
	svcCtx *svc.ServiceContext
	task.UnimplementedTaskServer
}

func NewTaskServer(svcCtx *svc.ServiceContext) *TaskServer {
	return &TaskServer{
		svcCtx: svcCtx,
	}
}

// 策展任务相关
func (s *TaskServer) CreateCuratorialTask(ctx context.Context, in *task.CreatePublishTaskInput) (*task.Mistake, error) {
	l := logic.NewCreateCuratorialTaskLogic(ctx, s.svcCtx)
	return l.CreateCuratorialTask(in)
}

func (s *TaskServer) QueryTaskList(ctx context.Context, in *task.PublishTaskInput) (*task.RePublishTask, error) {
	l := logic.NewQueryTaskListLogic(ctx, s.svcCtx)
	return l.QueryTaskList(in)
}

func (s *TaskServer) QueryTaskDetails(ctx context.Context, in *task.TaskDetailsInput) (*task.ReTaskDetails, error) {
	l := logic.NewQueryTaskDetailsLogic(ctx, s.svcCtx)
	return l.QueryTaskDetails(in)
}

func (s *TaskServer) QueryUserLaunchTaskList(ctx context.Context, in *task.UserLaunchTaskListInput) (*task.RePublishTask, error) {
	l := logic.NewQueryUserLaunchTaskListLogic(ctx, s.svcCtx)
	return l.QueryUserLaunchTaskList(in)
}

func (s *TaskServer) CreateLabel(ctx context.Context, in *task.CreateLabelInput) (*task.Mistake, error) {
	l := logic.NewCreateLabelLogic(ctx, s.svcCtx)
	return l.CreateLabel(in)
}

func (s *TaskServer) DeleteLabel(ctx context.Context, in *task.TaskIDInquireInput) (*task.Mistake, error) {
	l := logic.NewDeleteLabelLogic(ctx, s.svcCtx)
	return l.DeleteLabel(in)
}

func (s *TaskServer) QueryLabelList(ctx context.Context, in *task.UserIDInquireInput) (*task.ReLabelList, error) {
	l := logic.NewQueryLabelListLogic(ctx, s.svcCtx)
	return l.QueryLabelList(in)
}

func (s *TaskServer) PerformTask(ctx context.Context, in *task.PerformTaskInput) (*task.Mistake, error) {
	l := logic.NewPerformTaskLogic(ctx, s.svcCtx)
	return l.PerformTask(in)
}

func (s *TaskServer) VoluntarilyTaskSchedule(ctx context.Context, in *task.VoluntarilyTaskScheduleInput) (*task.Mistake, error) {
	l := logic.NewVoluntarilyTaskScheduleLogic(ctx, s.svcCtx)
	return l.VoluntarilyTaskSchedule(in)
}

// 每日任务
func (s *TaskServer) AmendTreasureTask(ctx context.Context, in *task.TreasureTaskSrtInput) (*task.Mistake, error) {
	l := logic.NewAmendTreasureTaskLogic(ctx, s.svcCtx)
	return l.AmendTreasureTask(in)
}

func (s *TaskServer) ChangeTreasureTask(ctx context.Context, in *task.TreasureTaskInput) (*task.Mistake, error) {
	l := logic.NewChangeTreasureTaskLogic(ctx, s.svcCtx)
	return l.ChangeTreasureTask(in)
}

func (s *TaskServer) QueryTreasureTaskList(ctx context.Context, in *task.TreasureTaskListInput) (*task.ReTreasureTaskSrt, error) {
	l := logic.NewQueryTreasureTaskListLogic(ctx, s.svcCtx)
	return l.QueryTreasureTaskList(in)
}

func (s *TaskServer) QuerySubtaskStyle(ctx context.Context, in *task.TaskIDInquireInput) (*task.ReSubtaskStyle, error) {
	l := logic.NewQuerySubtaskStyleLogic(ctx, s.svcCtx)
	return l.QuerySubtaskStyle(in)
}

func (s *TaskServer) AmendAssociatedSubtask(ctx context.Context, in *task.AssociatedSubtaskSrt) (*task.Mistake, error) {
	l := logic.NewAmendAssociatedSubtaskLogic(ctx, s.svcCtx)
	return l.AmendAssociatedSubtask(in)
}

func (s *TaskServer) DeleteAssociatedSubtask(ctx context.Context, in *task.TaskIDInquireInput) (*task.Mistake, error) {
	l := logic.NewDeleteAssociatedSubtaskLogic(ctx, s.svcCtx)
	return l.DeleteAssociatedSubtask(in)
}

func (s *TaskServer) QueryAssociatedSubtask(ctx context.Context, in *task.TaskIDInquireInput) (*task.ReAssociatedSubtask, error) {
	l := logic.NewQueryAssociatedSubtaskLogic(ctx, s.svcCtx)
	return l.QueryAssociatedSubtask(in)
}

func (s *TaskServer) AmendChestCollection(ctx context.Context, in *task.AmendChestCollectionInput) (*task.Mistake, error) {
	l := logic.NewAmendChestCollectionLogic(ctx, s.svcCtx)
	return l.AmendChestCollection(in)
}

func (s *TaskServer) QueryChestCollection(ctx context.Context, in *task.UserIDInquireInput) (*task.ReChestCollectionSrt, error) {
	l := logic.NewQueryChestCollectionLogic(ctx, s.svcCtx)
	return l.QueryChestCollection(in)
}

func (s *TaskServer) CreateUserPowerTask(ctx context.Context, in *task.CreateUserPowerTaskInput) (*task.Mistake, error) {
	l := logic.NewCreateUserPowerTaskLogic(ctx, s.svcCtx)
	return l.CreateUserPowerTask(in)
}

func (s *TaskServer) CreateSubtaskStyle(ctx context.Context, in *task.UserIDInquireInput) (*task.Mistake, error) {
	l := logic.NewCreateSubtaskStyleLogic(ctx, s.svcCtx)
	return l.CreateSubtaskStyle(in)
}
