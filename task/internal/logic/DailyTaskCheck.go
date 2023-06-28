package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"task/internal/model"
	"task/internal/svc"
	"task/task"
)

type DailyTaskCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDailyTaskCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DailyTaskCheckLogic {
	return &DailyTaskCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VerifyAssociatedSubtask 验证每日任务是否完成(站外)(未完成)
// @param userId string 用户ID
// @param taskId int    任务ID
func VerifyAssociatedSubtask(userId string, taskId int) error {
	return nil
}

// verifyInviteFriendsSupport 每日任务—校验邀请好友助力任务(未完)（缺少算力增加）
// @param userId string
func verifyInviteFriendsSupport(l *CreateUserPowerTaskLogic, in *task.CreateUserPowerTaskInput, userId, address string) (bool, error) {
	// 查询今日是否发布用户助力任务
	userPublishesHelperTask, err := QueryUserPublishesHelperTask(l, in, userId)
	if err != nil {
		return false, err
	}
	if userPublishesHelperTask.Id == 0 {
		return false, fmt.Errorf("用户发布任务不存在")
	}
	// 推特信息验证
	// 验证是否有人助力任务
	return false, nil
}

// verifySupportFriends 每日任务—校验帮助好友助力任务(未完)（缺少算力增加）
func verifySupportFriends(l *CreateCuratorialTaskLogic, userId string) (bool, error) {
	// 获取今日是否有这个任务
	taskData, err := QueryTaskDay(l, "帮助好友助力")
	if err != nil {
		return false, err
	}
	// 判断是否存任务
	if taskData.Id == 0 {
		return false, err
	}
	// 获取用户帮助助力次数
	userPowerTaskAmount, err := l.svcCtx.UserPowerTaskModel.FindNumberHelpings(l.ctx, userId)
	if err != nil {
		return false, err
	}
	// 判断任务次数是否超出
	if taskData.Number.Int64 > userPowerTaskAmount {
		return false, nil
	}
	// 每日任务—更新日常任务完成表
	exist, err := amendDailyTask(l, userId, taskData)
	if !exist || err != nil {
		return false, nil
	}
	return false, nil
}

// verifyBargainForFriends 每日任务—校验帮助好友砍价任务(未完)（缺少算力增加）
func verifyBargainForFriends(serId string) (bool, error) {
	return false, nil
}

// verifyPublishCuration  每日任务—校验发布策展任务(待验证)（缺少算力增加）
// @param userId    string 用户ID
// @param subtaskId uint   子任务ID
func verifyPublishCuration(l *CreateCuratorialTaskLogic, userId string, subtaskId string) (bool, error) {
	// 查询今日宝箱任务
	associatedSubtask, err := QueryTaskDay(l, subtaskId)
	if err != nil {
		return false, err
	}
	// 判断今日任务是否存在
	if associatedSubtask.Id == 0 {
		return false, err
	}
	// 查询今日是否发布策展任务数量
	publishTask, err := l.svcCtx.PublishTaskModel.FindTaskCount(l.ctx, userId)
	if err != nil {
		return false, nil
	}
	// 判断今日数量是否已经超出
	if publishTask > associatedSubtask.Number.Int64 {
		return false, nil
	}
	// 更新数据
	exist, err := amendDailyTask(l, userId, associatedSubtask)
	if !exist || err != nil {
		return false, nil
	}
	return true, nil
}

// verifyCompleteCurationMissions  每日任务—校验完成策展任务(未完)（三个调用，1-关注(内部)，2-喜欢(内部)，3-必要任务(验证)）(待验证)（缺少算力增加）
func verifyCompleteCurationMissions(l *CreateCuratorialTaskLogic, userId, subtaskId string, taskId int64) (bool, error) {
	// 查询今日宝箱任务
	associatedSubtask, err := QueryTaskDay(l, subtaskId)
	if err != nil {
		return false, err
	}
	// 判断今日任务是否存在
	if associatedSubtask.Id == 0 {
		return false, err
	}
	// 获取任务要求
	taskDemand, err := l.svcCtx.TaskDemandModel.FindList(l.ctx, taskId)
	if err != nil {
		return false, err
	}
	// 获取用户参与者任务要求完成度
	treasureStage, err := l.svcCtx.TreasureStagesModel.FindPersonalMissionCompletionDetails(l.ctx, userId, taskId)
	if err != nil {
		return false, err
	}
	// 判断是否都完成
	for _, item1 := range taskDemand {
		for _, item2 := range treasureStage {
			if item1.TaskName == item2.TaskName {
				switch item2.TaskStatus.Int64 {
				case 0:
					break
				case 1:
					return false, nil
				case 2:
					return false, nil
				}
			}
		}
	}
	// 更新数据
	exist, err := amendDailyTask(l, userId, associatedSubtask)
	if !exist || err != nil {
		return false, nil
	}
	return true, nil
}

// verifyCurationParticipantsReach 每日任务—校验当日发布的策展任务参与者达到10人任务(未完)（三个调用，1-关注(内部)，2-喜欢(内部)，3-必要任务(验证)）(待验证)（缺少算力增加）
func verifyCurationParticipantsReach(l *CreateCuratorialTaskLogic, userId, subtaskId string, taskId int64) (bool, error) {
	// 查询今日宝箱任务
	associatedSubtask, err := QueryTaskDay(l, subtaskId)
	if err != nil {
		return false, err
	}
	// 判断今日任务是否存在
	if associatedSubtask.Id == 0 {
		return false, err
	}
	// 获取任务参与者数量
	participant, err := l.svcCtx.ParticipantModel.FindTaskParticipant(l.ctx, userId, taskId)
	if err != nil {
		return false, err
	}
	if len(participant) < 10 {
		return false, nil
	}
	// 更新数据
	exist, err := amendDailyTask(l, userId, associatedSubtask)
	if !exist || err != nil {
		return false, nil
	}
	return true, nil
}

// verifyGetMINER  每日任务—校验获取矿机任务(未完)
func verifyGetMINER(serId string) (bool, error) {
	return false, nil
}

// verifyGetENERGYWATER  每日任务—校验获取能量水任务(未完)
func verifyGetENERGYWATER(serId string) (bool, error) {
	return false, nil
}

// QueryTaskDay 每日任务—查询今日的宝箱任务，是否存在本任务
// @param subtaskId uint 子任务ID
func QueryTaskDay(l *CreateCuratorialTaskLogic, subtaskName string) (*model.AssociatedSubtask, error) {
	// 查询今日宝箱信息
	treasureTask, err := l.svcCtx.TreasureTaskModel.FindTreasureQuantity(l.ctx)
	if err != nil {
		return nil, err
	}
	// 获取相应的子任务信息
	associatedSubtask, err := l.svcCtx.AssociatedSubtaskModel.FindSubtaskInformation(l.ctx, subtaskName, treasureTask.Id)
	if err != nil {
		return nil, err
	}
	return associatedSubtask, nil
}

// amendDailyTask 每日任务—更新日常任务完成表
func amendDailyTask(l *CreateCuratorialTaskLogic, userId string, associatedSubtask *model.AssociatedSubtask) (bool, error) {
	// 判断日常任务完是否已存在
	dailyTask, err := l.svcCtx.DailyTaskModel.FindCompletionTask(l.ctx, userId, associatedSubtask.Id)
	if err != nil {
		return false, nil
	}
	// 判断今日任务是否存在
	if dailyTask.Id > 0 {
		// 任务存在进行更新
		err = l.svcCtx.DailyTaskModel.Update(l.ctx, userId, associatedSubtask.Id, dailyTask.Complete.Int64+100/associatedSubtask.Number.Int64, dailyTask.Experience.Int64+associatedSubtask.Experience.Int64/associatedSubtask.Number.Int64)
		if err != nil {
			return false, nil
		}
	} else {
		// 任务不存在进行创建
		dailyTaskSrt := &model.DailyTask{
			UserId:     userId,
			TaskId:     associatedSubtask.Id,
			Complete:   sql.NullInt64{Int64: 100 / associatedSubtask.Number.Int64, Valid: true},
			Experience: sql.NullInt64{Int64: associatedSubtask.Experience.Int64 / associatedSubtask.Number.Int64, Valid: true},
		}
		_, err = l.svcCtx.DailyTaskModel.Insert(l.ctx, dailyTaskSrt)
		if err != nil {
			return false, nil
		}
	}
	return true, nil
}
