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

type VoluntarilyTaskScheduleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVoluntarilyTaskScheduleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VoluntarilyTaskScheduleLogic {
	return &VoluntarilyTaskScheduleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// VoluntarilyTaskSchedule 完成站内策展任务(喜欢+关注)(未完成)
// @param userId  string 用户UUID
// @param taskId  uint   任务ID
// @param genre   int    任务类型，1-喜欢，2-关注
func (l *VoluntarilyTaskScheduleLogic) VoluntarilyTaskSchedule(in *task.VoluntarilyTaskScheduleInput) (*task.Mistake, error) {
	// 判断任务类型
	var demand int
	switch in.Genre {
	case 1:
		// 调用第三方，实现推特喜欢
		demand = 4
	case 2:
		// 调用第三方，实现推特关注
		demand = 5
	default:
		return nil, fmt.Errorf("任务类型不合规")
	}
	// 更新参与者列表
	err := UpdateParticipant(l, in.TaskId, in.UserId, demand)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// UpdateParticipant 更新参与者列表
// @param taskId uint   任务ID
// @param userID string 用户ID
// @param demand int    任务要求，1-转推，2-引用，3-回复,4-喜欢,5-关注
func UpdateParticipant(l *VoluntarilyTaskScheduleLogic, taskId int64, userId string, demand int) error {
	// 判断此任务参与者是否存在
	participant, err := l.svcCtx.ParticipantModel.FindListParticipants(l.ctx, userId, taskId)
	if err != nil {
		return err
	}
	// 判断参与用户列表是否存在
	if participant.Id > 0 {
		// 创建参与者任务要求度
		err := CreateTreasureStage(l, taskId, userId, demand)
		if err != nil {
			return err
		}
	} else {
		// 查询用户ID的信息（用户表未完成）
		userInfo, err := QueryUser(userId)
		if err != nil {
			return err
		}
		// 创建数据
		participantCre := model.Participant{
			UserId:      userId,
			UserName:    userInfo,                         // 未完
			NickName:    sql.NullString{String: userInfo}, // 未完
			Avatar:      sql.NullString{String: userInfo}, // 未完
			AwardAmount: participant.AwardAmount,
			TaskId:      taskId,
			Status:      1,
		}
		_, err = l.svcCtx.ParticipantModel.Insert(l.ctx, &participantCre)
		if err != nil {
			return err
		}
		// 创建任务完成度
		err = CreateTreasureStage(l, taskId, userId, demand)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateTreasureStage 创建参与者任务要求完成度
func CreateTreasureStage(l *VoluntarilyTaskScheduleLogic, taskId int64, userId string, taskName int) error {
	treasureStage := model.TreasureStages{
		UserId:     userId,
		TaskId:     taskId,
		TaskName:   sql.NullInt64{Int64: int64(taskName)},
		TaskStatus: sql.NullInt64{Int64: 0},
	}
	_, err := l.svcCtx.TreasureStagesModel.Insert(l.ctx, &treasureStage)
	if err != nil {
		return err
	}
	return nil
}
