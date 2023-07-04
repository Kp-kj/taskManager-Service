package logic

import (
	"context"
	"database/sql"
	"fmt"
	"taskManager-Service/internal/model"
	"time"

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
	fmt.Print("11111111111111111111111\n")
	exist, err := l.svcCtx.UserPublishesHelperTaskModel.FindUserPublishesHelperTask(l.ctx, in.UserId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if exist != nil {
		return &task.Mistake{Msg: "今日用户发布助力任务已存在"}, nil
	}
	// 获取今日任务-助力信息
	fmt.Print("222222222222222222\n")
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindTreasureQuantity(l.ctx)
	if err != nil {
		return nil, err
	}
	fmt.Print("3333333333333333333333\n")
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindSubtaskInformation(l.ctx, "邀请好友助力", treasureTask.Id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if associatedSubtask == nil {
		return &task.Mistake{Msg: "今日任务不存在邀请好友助力信息"}, nil
	}
	// 创建用户用户发布助力任务
	userPublishes := &model.UserPublishesHelperTask{
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UserId:    in.UserId,
		UserName:  sql.NullString{String: in.UserName, Valid: true},
		Avatar:    sql.NullString{String: in.Avatar, Valid: true},
		Article:   associatedSubtask.Article,
		Link:      associatedSubtask.Link,
		Label:     associatedSubtask.Label,
	}
	fmt.Print("44444444444444444444444\n")
	_, err = l.svcCtx.UserPublishesHelperTaskModel.Insert(l.ctx, userPublishes)
	if err != nil {
		return nil, err
	}
	return &task.Mistake{Msg: "succeed"}, nil
}
