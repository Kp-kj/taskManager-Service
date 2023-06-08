package logic

import (
	"context"
	"database/sql"
	"fmt"
	"taskManager-Service-main/task/internal/model"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

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

func (l *AmendAssociatedSubtaskLogic) AmendAssociatedSubtask(in *task.AssociatedSubtaskSrt) (*task.Mistake, error) {
	// 判断字段信息
	if in.TaskStatus == 0 {
		if in.Article != "" || in.Link != "" || in.Label != "" {
			return &task.Mistake{Err: fmt.Sprintf("站内任务不需要额外信息")}, fmt.Errorf("站内任务不需要额外信息")
		}
	}
	// 查询宝箱样式ID
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindOne(l.ctx, int64(in.TreasureId))
	if err != nil {
		return nil, err
	}
	// 查询是否存在已关联任务FindAssociatedSubtask
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindAssociatedSubtask(l.ctx, int64(in.TreasureId), "id")
	if err != nil {
		return nil, err
	}
	// 获取子任务是否一致
	subtaskStyle, err := l.svcCtx.SubtaskStyleModel.FindOne(l.ctx, int64(in.TaskId))
	if err != nil {
		return nil, err
	}
	if subtaskStyle.TaskName.String != in.TaskName {
		return &task.Mistake{Err: fmt.Sprintf("任务样式不一致")}, fmt.Errorf("任务样式不一致")
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
			return &task.Mistake{Err: fmt.Sprintf("奖励超出预算数量")}, fmt.Errorf("奖励超出预算数量")
		}
		// 判断经验是否超出
		if int(treasureTask.ExperienceReward) < experience+int(in.Experience) {
			return &task.Mistake{Err: fmt.Sprintf("经验超出预算数量")}, fmt.Errorf("经验超出预算数量")
		}
	}
	err = MakeDataChanges(l, in)
	if err != nil {
		return nil, err
	}
	return &task.Mistake{}, nil
}

// MakeDataChanges 进行数据修改
func MakeDataChanges(l *AmendAssociatedSubtaskLogic, in *task.AssociatedSubtaskSrt) error {
	// 赋值数据
	associatedSubtaskSrt := model.AssociatedSubtask{
		TaskId:         sql.NullInt64{Int64: int64(in.TaskId)},
		TaskName:       sql.NullString{String: in.TaskName},
		TaskNameEng:    sql.NullString{String: in.TaskNameEng},
		TaskDetails:    sql.NullString{String: in.TaskDetails},
		TaskDetailsEng: sql.NullString{String: in.TaskDetailsEng},
		TaskStatus:     sql.NullInt64{Int64: int64(in.TaskStatus)},
		Reward:         sql.NullInt64{Int64: int64(in.Reward)},
		Experience:     sql.NullInt64{Int64: int64(in.Experience)},
		Number:         sql.NullInt64{Int64: int64(in.Number)},
		Article:        sql.NullString{String: in.Article},
		Link:           sql.NullString{String: in.Link},
		Label:          sql.NullString{String: in.Label},
		TreasureId:     int64(in.TreasureId),
	}
	// 判断关联子任务是否存在
	if in.TaskId == 0 {
		// 创建关联子任务
		_, err := l.svcCtx.AssociatedSubtaskModel.Insert(l.ctx, &associatedSubtaskSrt)
		if err != nil {
			return err
		}
	} else {
		err := l.svcCtx.AssociatedSubtaskModel.Update(l.ctx, &associatedSubtaskSrt)
		if err != nil {
			return err
		}
	}
	return nil
}
