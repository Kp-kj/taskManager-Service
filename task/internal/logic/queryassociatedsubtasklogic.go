package logic

import (
	"context"
	"task/internal/svc"
	"task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAssociatedSubtaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAssociatedSubtaskLogic {
	return &QueryAssociatedSubtaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryAssociatedSubtask 查询关联子任务
func (l *QueryAssociatedSubtaskLogic) QueryAssociatedSubtask(in *task.TaskIDInquireInput) (*task.ReAssociatedSubtask, error) {
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindAssociatedSubtask(l.ctx, int64(in.Id), "treasure_id")
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	var reAssociatedSubtask []*task.AssociatedSubtask
	for _, item := range associatedSubtask {
		associatedSubtask := task.AssociatedSubtask{
			TaskId:         uint64(item.TaskId.Int64),
			TaskName:       item.TaskName.String,
			TaskNameEng:    item.TaskNameEng.String,
			TaskDetails:    item.TaskDetails.String,
			TaskDetailsEng: item.TaskDetailsEng.String,
			TaskStatus:     item.TaskStatus.Int64,
			Reward:         item.Reward.Int64,
			Experience:     item.Experience.Int64,
			Number:         item.Number.Int64,
			Article:        item.Article.String,
			Link:           item.Link.String,
			Label:          item.Label.String,
			TreasureId:     uint64(item.TreasureId),
		}
		reAssociatedSubtask = append(reAssociatedSubtask, &associatedSubtask)
	}
	ReAssociatedSubtask := task.ReAssociatedSubtask{AssociatedSubtask: reAssociatedSubtask}
	return &ReAssociatedSubtask, nil
}
