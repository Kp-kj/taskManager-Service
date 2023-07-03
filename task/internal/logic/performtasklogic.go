package logic

import (
	"context"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type PerformTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPerformTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PerformTaskLogic {
	return &PerformTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PerformTask 判断是否完成任务
func (l *PerformTaskLogic) PerformTask(in *task.PerformTaskInput) (*task.Mistake, error) {
	// 获取策展任务详情
	taskDemand, err := l.svcCtx.TaskDemandModel.FindList(l.ctx, int64(in.TaskId))
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	if len(taskDemand) == 0 {
		return &task.Mistake{Msg: "策展任务要求不存在"}, nil
	}
	// 获取个人任务完成详情
	treasureStage, err := l.svcCtx.TreasureStagesModel.FindPersonalMissionCompletionDetails(l.ctx, in.UserId, int64(in.TaskId))
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	if len(treasureStage) == 0 {
		return &task.Mistake{Msg: "任务未完成"}, nil
	}
	// 判断是否都完成
	for _, item1 := range taskDemand {
		for _, item2 := range treasureStage {
			if item1.TaskName == item2.TaskName {
				switch item2.TaskStatus.Int64 {
				case 0:
					break
				case 1:
					return nil, nil
				case 2:
					return nil, nil
				}
			}
		}
	}
	// 更新个人任务状态
	err = l.svcCtx.ParticipantModel.UpdateParticipant(l.ctx, in.UserId, in.TaskId)
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	// 更新完成人物
	err = l.svcCtx.PublishTaskModel.UpdateNumberCompleters(l.ctx, in.TaskId)
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	return &task.Mistake{Msg: "任务已完成"}, nil
}
