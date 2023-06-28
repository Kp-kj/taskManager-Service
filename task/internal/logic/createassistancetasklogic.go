package logic

import (
	"context"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAssistanceTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAssistanceTaskLogic {
	return &CreateAssistanceTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateAssistanceTask 创建用户发布助力任务
func (l *CreateAssistanceTaskLogic) CreateAssistanceTask(in *task.CreateUserPublishingAssistanceTaskInput) (*task.Mistake, error) {
	// 获取用户发布助力任务
	exist, err := l.svcCtx.UserPublishesHelperTaskModel.FindUserPublishesHelperTask(l.ctx, in.UserId)
	if err != nil || err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if exist != nil {
		return &task.Mistake{Msg: "今日用户发布助力任务已存在"}, nil
	}
	// 获取用户信息
	//
	return &task.Mistake{}, nil
}
