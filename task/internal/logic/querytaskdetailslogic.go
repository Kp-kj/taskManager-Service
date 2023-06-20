package logic

import (
	"context"
	"taskManager-Service-main/task/internal/model"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryTaskDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryTaskDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryTaskDetailsLogic {
	return &QueryTaskDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryTaskDetails 查询任务详情
func (l *QueryTaskDetailsLogic) QueryTaskDetails(in *task.TaskDetailsInput) (*task.ReTaskDetails, error) {
	// 查询策展任务详情
	publishTask, err := l.svcCtx.PublishTaskModel.FindOne(l.ctx, int64(in.TaskId))
	if err != nil {
		return nil, err
	}
	// 查询任务要求详情
	taskDemand, err := l.svcCtx.TaskDemandModel.FindList(l.ctx, publishTask.Id)
	if err != nil {
		return nil, err
	}
	// 赋值任务要求详情
	rePublishTask := givePublishTask(publishTask, taskDemand)

	// 获取推特详情（未完）+用户账户

	// 查询用户任务完成度
	taskDemandBak, err := QueryUserTaskCompletionDegree(l, in, taskDemand)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	// 获取任务参与者列表
	participant, err := l.svcCtx.ParticipantModel.FinParticipantList(l.ctx, in.TaskId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	// 赋值任务参与者列表
	reParticipant := giveParticipant(participant)
	// 赋值返回值
	reTaskDetails := task.ReTaskDetails{
		RePublishTaskSrt: &rePublishTask,
		Participant:      reParticipant,
		TaskDemandBak:    taskDemandBak,
	}
	return &reTaskDetails, nil
}

// QueryUserTaskCompletionDegree 查询用户任务完成度
func QueryUserTaskCompletionDegree(l *QueryTaskDetailsLogic, in *task.TaskDetailsInput, taskDemand []*model.TaskDemand) ([]*task.TaskDemandBak, error) {
	// 查询用户任务完成度
	var taskDemandBak []*task.TaskDemandBak
	for _, item := range taskDemand {
		treasureStage, err := l.svcCtx.TreasureStagesModel.FindSpecificTask(l.ctx, in.UserId, uint(in.TaskId), int(item.TaskName.Int64))
		if err != nil {
			return nil, err
		}
		taskDemandBak = append(taskDemandBak, &task.TaskDemandBak{
			TaskID:     uint64(item.TaskName.Int64),
			TaskName:   item.TaskName.Int64,
			TaskDemand: item.TaskDemand.String,
			Article:    item.Article.String,
			Done:       int32(treasureStage.TaskStatus.Int64),
		})
	}
	return taskDemandBak, nil
}

// givePublishTask 赋值策展任务详情数据
func givePublishTask(publishTask *model.PublishTask, taskDemand []*model.TaskDemand) (rePublishTask task.RePublishTaskBak) {
	var taskDemandSrt []*task.TaskDemand
	for _, item := range taskDemand {
		taskDemandSrt = append(taskDemandSrt, &task.TaskDemand{
			TaskId:     uint64(item.TaskId.Int64),
			TaskName:   int32(item.TaskName.Int64),
			TaskDemand: item.TaskDemand.String,
			Article:    item.Article.String,
		})
	}
	rePublishTask = task.RePublishTaskBak{
		TaskId:        uint64(publishTask.Id),
		CreatedAt:     publishTask.CreatedAt.Time.String(),
		Creator:       publishTask.Creator,
		CreatorName:   publishTask.Creator, // (未完，用户账户)
		CreatorNick:   publishTask.Creator, // (未完，用户昵称)
		CreatorAvatar: publishTask.Creator, // (未完，用户头像)
		Status:        int32(publishTask.Status),
		TweetDetails:  publishTask.TweetAddress,
		TweetPicture:  publishTask.TweetAddress, // (未完，推特内容)
		Label:         publishTask.Label,
		AwardBudget:   publishTask.AwardBudget,
		MaxUser:       publishTask.MaxUser,
		AwardAmount:   publishTask.AwardAmount,
		EndTime:       publishTask.EndTime.String(),
		Accomplish:    publishTask.Accomplish,
		TaskDemand:    taskDemandSrt,
	}
	return
}

// giveParticipant 赋值任务参与者列表
func giveParticipant(participant []*model.Participant) (reParticipant []*task.ParticipantBak) {
	for _, item := range participant {
		reParticipant = append(reParticipant, &task.ParticipantBak{
			UserId:      item.UserId,
			UserName:    item.UserName,
			NickName:    item.NickName.String,
			Avatar:      item.Avatar.String,
			AwardAmount: item.AwardAmount,
			TaskID:      uint64(item.TaskId),
			Status:      int32(item.Status),
		})
	}
	return
}
