package logic

import (
	"context"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

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
		return &task.Mistake{Msg: err.Error()}, err
	}
	return &task.Mistake{Msg: "succeed"}, nil
}
