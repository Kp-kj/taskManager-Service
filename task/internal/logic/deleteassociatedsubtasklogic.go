package logic

import (
	"context"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAssociatedSubtaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAssociatedSubtaskLogic {
	return &DeleteAssociatedSubtaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteAssociatedSubtask 删除管理子任务
func (l *DeleteAssociatedSubtaskLogic) DeleteAssociatedSubtask(in *task.TaskIDInquireInput) (*task.Mistake, error) {
	err := l.svcCtx.AssociatedSubtaskModel.Delete(l.ctx, int64(in.Id))
	if err != nil {
		return nil, err
	}
	return &task.Mistake{}, nil
}
