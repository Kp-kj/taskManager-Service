package logic

import (
	"context"
	"strings"
	"taskManager-Service-main/task/internal/model"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLabelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLabelLogic {
	return &CreateLabelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateLabel 创建标签
func (l *CreateLabelLogic) CreateLabel(in *task.CreateLabelInput) (*task.Mistake, error) {
	// 获取现有标签数量
	amount, err := l.svcCtx.LabelModel.FindLabelAmount(l.ctx, in.UserId)
	if err != nil {
		return &task.Mistake{Msg: "数据未查询到"}, err
	}
	// 判断标签数量
	if amount > 20 {
		return &task.Mistake{Msg: "标签数量大于20条"}, nil
	}
	// 判断标签字数
	if in.Label == "" || strings.Count(in.Label, "") >= 10 {
		return &task.Mistake{Msg: "标签字数不合规"}, nil
	}
	// 创建标签
	label := &model.Label{
		Creator: in.UserId,
		Content: in.Label,
	}
	_, err = l.svcCtx.LabelModel.Insert(l.ctx, label)
	if err != nil {
		return &task.Mistake{}, err
	}
	return &task.Mistake{Msg: "succeed"}, nil
}
