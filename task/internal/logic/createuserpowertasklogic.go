package logic

import (
	"context"
	"fmt"
	"taskManager-Service-main/task/internal/model"
	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserPowerTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserPowerTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserPowerTaskLogic {
	return &CreateUserPowerTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUserPowerTask 创建用户助力信息
func (l *CreateUserPowerTaskLogic) CreateUserPowerTask(in *task.CreateUserPowerTaskInput) (*task.Mistake, error) {
	// 查询是否存在助力任务
	userPublishesHelperTask, err := QueryUserPublishesHelperTask(l, in, in.PublishesUserId)
	if err != nil {
		return nil, err
	}
	if userPublishesHelperTask.Id == 0 {
		return nil, fmt.Errorf("用户发布助力信息不存在")
	}
	//赋值数据
	userPowerTask := model.UserPowerTask{
		PublishesUserId: in.PublishesUserId,
		HelperUserId:    in.HelperUserId,
	}
	// 创建用户助力信息
	_, err = l.svcCtx.UserPowerTaskModel.Insert(l.ctx, &userPowerTask)
	if err != nil {
		return nil, err
	}
	return &task.Mistake{}, nil
}

// QueryUserPublishesHelperTask 查询用户发布助力信息
// @param userId string 用户ID
func QueryUserPublishesHelperTask(l *CreateUserPowerTaskLogic, in *task.CreateUserPowerTaskInput, userId string) (*model.UserPublishesHelperTask, error) {
	// 查询用户助力信息
	userPublishesHelperTask, err := l.svcCtx.UserPublishesHelperTaskModel.FindUserPublishesHelperTask(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return userPublishesHelperTask, nil
}
