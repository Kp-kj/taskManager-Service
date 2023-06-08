package logic

import (
	"context"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLabelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLabelListLogic {
	return &QueryLabelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryLabelListLogic) QueryLabelList(in *task.UserIDInquireInput) (*task.ReLabelList, error) {
	_, err := l.svcCtx.LabelModel.FindList(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &task.ReLabelList{}, nil
}
