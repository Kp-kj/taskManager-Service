package logic

import (
	"context"
	"task/internal/svc"
	"task/task"

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

// QueryLabelList 获取标签列表
func (l *QueryLabelListLogic) QueryLabelList(in *task.LabelListInput) (*task.ReLabelList, error) {
	_, err := l.svcCtx.LabelModel.FindList(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &task.ReLabelList{}, nil
}
