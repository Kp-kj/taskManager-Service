package logic

import (
	"context"
	"database/sql"
	"fmt"
	"task/internal/model"
	"task/internal/svc"
	"task/task"

	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendAssociatedSubtaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAmendAssociatedSubtaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendAssociatedSubtaskLogic {
	return &AmendAssociatedSubtaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AmendAssociatedSubtask 创建关联子任务+编辑关联子任务
func (l *AmendAssociatedSubtaskLogic) AmendAssociatedSubtask(in *task.AssociatedSubtaskSrt) (*task.Mistake, error) {
	// 判断字段信息
	if in.TaskStatus == 0 {
		if in.Article != "" || in.Link != "" || in.Label != "" {
			return &task.Mistake{Msg: "站内任务不需要额外信息"}, fmt.Errorf("站内任务不需要额外信息")
		}
	}
	// 查询宝箱样式ID
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindOne(l.ctx, int64(in.TreasureId))
	if err != nil && err.Error() != "sql: no rows in result set" {
		return &task.Mistake{Msg: err.Error()}, err
	}
	// 查询是否存在已关联任务FindAssociatedSubtask
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindAssociatedSubtask(l.ctx, int64(in.TreasureId), "id")
	if err != nil && err.Error() != "sql: no rows in result set" {
		return &task.Mistake{Msg: err.Error()}, err
	}
	// 获取子任务是否一致
	subtaskStyle, err := l.svcCtx.SubtaskStyleModel.FindOne(l.ctx, int64(in.TaskId))
	if err != nil && err.Error() != "sql: no rows in result set" {
		return &task.Mistake{Msg: err.Error()}, err
	}
	if subtaskStyle.TaskName.String != in.TaskName {
		return &task.Mistake{Msg: "任务样式不一致"}, fmt.Errorf("任务样式不一致")
	}
	//  判断是否已存在关联子任务,进行奖励计算
	var experience int
	var reward int
	if len(associatedSubtask) > 0 {
		// 循环增加奖励和经验
		for _, item := range associatedSubtask {
			experience = experience + int(item.Experience.Int64)
			reward = reward + int(item.Reward.Int64)
		}
		// 判断奖励是否超出
		if int(treasureTask.TaskReward) < reward+int(in.Reward) {
			return &task.Mistake{Msg: "奖励超出预算数量"}, fmt.Errorf("奖励超出预算数量")
		}
		// 判断经验是否超出
		if int(treasureTask.ExperienceReward) < experience+int(in.Experience) {
			return &task.Mistake{Msg: "经验超出预算数量"}, fmt.Errorf("经验超出预算数量")
		}
	}
	err = MakeDataChanges(l, in)
	if err != nil {
		return &task.Mistake{Msg: err.Error()}, err
	}
	return &task.Mistake{Msg: "succeed"}, nil
}

// MakeDataChanges 进行数据修改
func MakeDataChanges(l *AmendAssociatedSubtaskLogic, in *task.AssociatedSubtaskSrt) error {
	// 赋值数据
	associatedSubtaskSrt := model.AssociatedSubtask{
		CreatedAt:      sql.NullTime{Time: time.Now(), Valid: true},
		TaskId:         sql.NullInt64{Int64: int64(in.TaskId), Valid: true},
		TaskName:       sql.NullString{String: in.TaskName, Valid: true},
		TaskNameEng:    sql.NullString{String: in.TaskNameEng, Valid: true},
		TaskDetails:    sql.NullString{String: in.TaskDetails, Valid: true},
		TaskDetailsEng: sql.NullString{String: in.TaskDetailsEng, Valid: true},
		TaskStatus:     sql.NullInt64{Int64: int64(in.TaskStatus), Valid: true},
		Reward:         sql.NullInt64{Int64: int64(in.Reward), Valid: true},
		Experience:     sql.NullInt64{Int64: int64(in.Experience), Valid: true},
		Number:         sql.NullInt64{Int64: int64(in.Number), Valid: true},
		Article:        sql.NullString{String: in.Article, Valid: true},
		Link:           sql.NullString{String: in.Link, Valid: true},
		Label:          sql.NullString{String: in.Label, Valid: true},
		TreasureId:     int64(in.TreasureId),
	}
	// 判断关联子任务是否存在
	if in.AssociatedId == 0 {
		// 创建关联子任务
		_, err := l.svcCtx.AssociatedSubtaskModel.Insert(l.ctx, &associatedSubtaskSrt)
		if err != nil {
			return err
		}
	} else {
		// 编辑管理子任务
		associatedSubtaskSrt.Id = int64(in.AssociatedId)
		err := l.svcCtx.AssociatedSubtaskModel.Update(l.ctx, &associatedSubtaskSrt)
		if err != nil {
			return err
		}
	}
	return nil
}
