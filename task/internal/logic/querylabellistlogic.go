package logic

import (
	"context"
	"fmt"

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
func (l *QueryLabelListLogic) QueryLabelList(in *task.UserIDInquireInput) (*task.ReLabelListOut, error) {
	reLabelList, err := l.svcCtx.LabelModel.FindList(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var reLabel []*task.ReLabelList
	for _, item := range reLabelList {
		reLabelListSrt := task.ReLabelList{
			Id:      uint64(item.Id),
			Creator: item.Creator,
			Content: item.Content,
		}
		reLabel = append(reLabel, &reLabelListSrt)
	}
	reLabelSrt := &task.ReLabelListOut{
		ReLabelList: reLabel,
	}
	fmt.Printf("vxxxxx:%v\n", reLabelSrt)
	return reLabelSrt, nil
}
