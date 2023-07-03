package logic

import (
	"context"
	"fmt"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAssistanceTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAssistanceTaskLogic {
	return &QueryAssistanceTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryAssistanceTask 查询用户发布助力信息
func (l *QueryAssistanceTaskLogic) QueryAssistanceTask(in *task.UserIDInquireInput) (*task.UserPublishingAssistanceTask, error) {
	data, err := l.svcCtx.UserPublishesHelperTaskModel.FindUserPublishesHelperTask(l.ctx, in.UserId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if data == nil {
		return &task.UserPublishingAssistanceTask{}, fmt.Errorf("未查询到用户发布助力信息")
	}
	return &task.UserPublishingAssistanceTask{
		ID:        data.Id,
		CreatedAt: data.CreatedAt.Time.String(),
		UserId:    data.UserId,
		UserName:  data.UserName.String,
		Avatar:    data.Avatar.String,
		Article:   data.Article.String,
		Link:      data.Link.String,
		Label:     data.Label.String,
	}, nil
}
