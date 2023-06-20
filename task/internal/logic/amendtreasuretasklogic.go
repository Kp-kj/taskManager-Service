package logic

import (
	"context"
	"fmt"
	"strings"
	"taskManager-Service-main/task/internal/model"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendTreasureTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAmendTreasureTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendTreasureTaskLogic {
	return &AmendTreasureTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AmendTreasureTask 创建宝箱样式+编辑宝箱样式
func (l *AmendTreasureTaskLogic) AmendTreasureTask(in *task.TreasureTaskSrtInput) (*task.Mistake, error) {
	// 宝箱数据校验
	if in.Name == "" || strings.Count(in.Name, "") >= 21 {
		return &task.Mistake{Msg: "任务名称长度不合规"}, fmt.Errorf("任务名称长度不合规")
	}
	if int(in.RewardQuantity) != len(in.TreasureTaskStage) {
		return &task.Mistake{Msg: "宝箱阶段奖励数量不合规"}, fmt.Errorf("宝箱阶段奖励数量不合规")
	}
	// 赋值宝箱样式
	treasureTaskSrt := model.TreasureTask{
		Id:               int64(in.Id),
		Name:             in.Name,
		DemandIntegral:   in.DemandIntegral,
		TaskReward:       in.TaskReward,
		ExperienceReward: in.ExperienceReward,
		RewardQuantity:   in.RewardQuantity,
	}
	// 判断是创建还是编辑
	if in.Id > 0 {
		// 编辑
		err := EditTheChestStyle(l, in, treasureTaskSrt)
		if err != nil {
			return &task.Mistake{Msg: err.Error()}, err
		}
	} else {
		// 创建
		err := CreateTheChestStyle(l, in, treasureTaskSrt)
		if err != nil {
			return &task.Mistake{Msg: err.Error()}, err
		}
	}
	return &task.Mistake{Msg: "succeed"}, nil
}

// EditTheChestStyle 编辑宝箱样式
func EditTheChestStyle(l *AmendTreasureTaskLogic, in *task.TreasureTaskSrtInput, treasureTaskSrt model.TreasureTask) error {
	// 编辑宝箱样式
	err := l.svcCtx.TreasureTaskModel.Update(l.ctx, &treasureTaskSrt)
	if err != nil {
		return err
	}
	// 循环编辑宝箱样式
	for _, item := range in.TreasureTaskStage {
		err := l.svcCtx.TreasureTaskStageModel.UpdateTreasureTaskStage(l.ctx, item.Treasure, item.TreasureSequence, item.StageExperience, item.StageReward)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateTheChestStyle 创建宝箱样式
func CreateTheChestStyle(l *AmendTreasureTaskLogic, in *task.TreasureTaskSrtInput, treasureTaskSrt model.TreasureTask) error {
	// 添加宝箱样式
	_, err := l.svcCtx.TreasureTaskModel.Insert(l.ctx, &treasureTaskSrt)
	if err != nil {
		return err
	}
	// 获取刚刚添加的数据（后期改为UUID,未完）
	treasureTask, _ := l.svcCtx.TreasureTaskModel.FindNewlyAddedData(l.ctx, treasureTaskSrt)
	// 创建宝箱阶段
	for _, item := range in.TreasureTaskStage {
		treasureTaskStage := model.TreasureTaskStage{
			Treasure:         treasureTask.Id,
			TreasureSequence: int64(item.TreasureSequence),
			StageExperience:  int64(item.StageExperience),
			StageReward:      int64(item.StageReward),
		}
		_, err := l.svcCtx.TreasureTaskStageModel.Insert(l.ctx, &treasureTaskStage)
		if err != nil {
			return err
		}
	}
	return nil
}
