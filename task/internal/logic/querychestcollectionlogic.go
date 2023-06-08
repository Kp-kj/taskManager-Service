package logic

import (
	"context"
	"taskManager-Service-main/task/internal/model"
	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryChestCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryChestCollectionLogic {
	return &QueryChestCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryChestCollection 获取用户宝箱进度（未验证）
func (l *QueryChestCollectionLogic) QueryChestCollection(in *task.UserIDInquireInput) (*task.ReChestCollectionSrt, error) {
	// 获取宝箱样式
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindTreasureQuantity(l.ctx)
	if err != nil {
		return nil, err
	}
	// 获取宝箱关联子任务
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindAssociatedSubtask(l.ctx, treasureTask.Id, "treasure_id")
	if err != nil {
		return nil, err
	}
	// 获取宝箱领取度
	chestCollection, err := l.svcCtx.ChestCollectionsModel.FindIndividualAllowance(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	// 获取宝箱阶段信息
	treasureTaskStage, err := l.svcCtx.TreasureTaskStageModel.FindTreasureInformation(l.ctx, treasureTask.Id, "treasure")
	if err != nil {
		return nil, err
	}
	treasureTaskStageSrt, associatedSubtaskSrt, err := UserChestAssignmentInformation(l, in, treasureTaskStage, associatedSubtask, treasureTask)
	if err != nil {
		return nil, err
	}
	// 赋值返回结构体
	chestCollectionSrt := &task.ReChestCollectionSrt{
		SerId:             in.UserId,
		DemandIntegral:    treasureTask.DemandIntegral,
		ChestAmount:       chestCollection.ChestAmount.Int64,
		RewardQuantity:    treasureTask.RewardQuantity,
		TreasureTaskStage: treasureTaskStageSrt,
		AssociatedSubtask: associatedSubtaskSrt,
	}
	return chestCollectionSrt, nil
}

// UserChestAssignmentInformation 用户宝箱赋值信息
func UserChestAssignmentInformation(l *QueryChestCollectionLogic, in *task.UserIDInquireInput, treasureTaskStage []*model.TreasureTaskStage, associatedSubtask []*model.AssociatedSubtask, treasureTask *model.TreasureTask) ([]*task.TreasureTaskStageSeed, []*task.AssociatedSubtaskSeed, error) {
	var treasureTaskStageSrt []*task.TreasureTaskStageSeed
	// 计算宝箱位置数量
	for _, item := range treasureTaskStage {
		// 赋值宝箱详细信息
		treasureTaskStageSrt = append(treasureTaskStageSrt, &task.TreasureTaskStageSeed{
			Treasure:         item.Treasure,
			TreasureSequence: item.TreasureSequence,
			StageExperience:  item.StageExperience,
			TreasureRatio:    float64(item.StageExperience) / float64(treasureTask.DemandIntegral),
			StageReward:      item.StageReward,
		})
	}
	var associatedSubtaskSrt []*task.AssociatedSubtaskSeed
	for _, item := range associatedSubtask {
		// 获取日常任务完成信息
		complete, err := l.svcCtx.DailyTaskModel.FindCompletionTask(l.ctx, in.UserId, item.TaskId.Int64)
		if err != nil {
			return nil, nil, err
		}
		// 判断任务是否完成
		if complete.Complete.Int64 == 0 {
			continue
		}
		// 赋值返回日常任务信息
		associatedSubtaskSrt = append(associatedSubtaskSrt, &task.AssociatedSubtaskSeed{
			TreasureId:     item.TreasureId,
			TaskId:         item.TaskId.Int64,
			TaskName:       item.TaskName.String,
			TaskNameEng:    item.TaskNameEng.String,
			TaskDetails:    item.TaskDetails.String,
			TaskDetailsEng: item.TaskDetailsEng.String,
			TaskStatus:     item.TaskStatus.Int64,
			Reward:         item.Reward.Int64,
			Experience:     item.Experience.Int64,
			Number:         item.Number.Int64,
			Article:        item.Article.String,
			Link:           item.Link.String,
			Label:          item.Label.String,
			Complete:       complete.Complete.Int64 / (100 / item.Number.Int64),
		})
	}
	return treasureTaskStageSrt, associatedSubtaskSrt, nil
}
