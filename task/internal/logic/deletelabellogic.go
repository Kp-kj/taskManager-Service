package logic

import (
	"context"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLabelLogic {
	return &DeleteLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLabelLogic) DeleteLabel(in *task.TaskIDInquireInput) (*task.Mistake, error) {
	// 删除标签
	err := l.svcCtx.LabelModel.Delete(l.ctx, int64(in.Id))
	if err != nil {
		return nil, err
	}
	return &task.Mistake{}, nil
}
