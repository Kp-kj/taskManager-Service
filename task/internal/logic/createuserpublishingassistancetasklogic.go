package logic

import (
	"context"

	"taskManager-Service/internal/svc"
	"taskManager-Service/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserPublishingAssistanceTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserPublishingAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserPublishingAssistanceTaskLogic {
	return &CreateUserPublishingAssistanceTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserPublishingAssistanceTaskLogic) CreateUserPublishingAssistanceTask(in *task.CreateUserPublishingAssistanceTaskInput) (*task.Mistake, error) {
	// todo: add your logic here and delete this line

	return &task.Mistake{}, nil
}
