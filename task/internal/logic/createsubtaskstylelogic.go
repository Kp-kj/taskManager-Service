package logic

import (
	"context"
	"database/sql"
	"fmt"
	"taskManager-Service-main/task/internal/model"
	"time"

	"taskManager-Service-main/task/internal/svc"
	"taskManager-Service-main/task/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubtaskStyleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubtaskStyleLogic {
	return &CreateSubtaskStyleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateSubtaskStyle 创建子任务样式(初始化固定数据,初始化启动)
func (l *CreateSubtaskStyleLogic) CreateSubtaskStyle(in *task.UserIDInquireInput) (*task.Mistake, error) {
	subtaskStyle := []*model.SubtaskStyle{
		{TaskId: sql.NullInt64{Int64: 1}, TaskName: sql.NullString{String: "邀请好友助力"}, TaskNameEng: sql.NullString{String: "Invite friends to support"}, TaskDetails: sql.NullString{String: "将助力链接发送给好友，帮你点击助力，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Send links to them, once they support you will receive power"}, TaskStatus: sql.NullInt64{Int64: 1}},
		{TaskId: sql.NullInt64{Int64: 2}, TaskName: sql.NullString{String: "帮助好友助力"}, TaskNameEng: sql.NullString{String: "Support friends"}, TaskDetails: sql.NullString{String: "收到好友的助力链接，帮他们点击助力，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Receive links from friends, support them and you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 3}, TaskName: sql.NullString{String: "帮助好友砍价"}, TaskNameEng: sql.NullString{String: "Bargain for friends"}, TaskDetails: sql.NullString{String: "收到好友的砍价链接，帮他们点击砍价，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Receive links from friends, bargain for them and you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 4}, TaskName: sql.NullString{String: "发布策展任务"}, TaskNameEng: sql.NullString{String: "Publish curation"}, TaskDetails: sql.NullString{String: "在本平台发起一次策展，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Publish a curation, you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 5}, TaskName: sql.NullString{String: "完成策展任务"}, TaskNameEng: sql.NullString{String: "Complete curation missions"}, TaskDetails: sql.NullString{String: "完成平台内策展的相关任务，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Complete several missions from curation, you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 6}, TaskName: sql.NullString{String: "当日发布的策展任务参与者达到10人"}, TaskNameEng: sql.NullString{String: "Curation participants reach 10"}, TaskDetails: sql.NullString{String: "在本平台发起的单个策展参与人数达到10人，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Number of each curation participants reach 10, you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 7}, TaskName: sql.NullString{String: "获取矿机"}, TaskNameEng: sql.NullString{String: "Get MINER"}, TaskDetails: sql.NullString{String: "通过购买或砍价活动获得矿机时，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Purchase or get MINER by bargain, you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
		{TaskId: sql.NullInt64{Int64: 8}, TaskName: sql.NullString{String: "获取能量水"}, TaskNameEng: sql.NullString{String: "Get ENERGY WATER"}, TaskDetails: sql.NullString{String: "通过购买获得能量水时，你将获得算力"}, TaskDetailsEng: sql.NullString{String: "Purchase ENERGY WATER, you will receive power"}, TaskStatus: sql.NullInt64{Int64: 0}},
	}
	time.Sleep(5 * time.Second)
	// 判断是否已创建数据
	subtaskStyleSrt, err := l.svcCtx.SubtaskStyleModel.FindSubtaskStyleList(l.ctx)
	if err != nil {
		return nil, err
	}
	if len(subtaskStyleSrt) > 0 {
		return nil, fmt.Errorf("样式已创建")
	}
	// 不存在，进行创建
	for _, item := range subtaskStyle {
		_, err = l.svcCtx.SubtaskStyleModel.Insert(l.ctx, item)
		if err != nil {
			return nil, err
		}
	}
	return &task.Mistake{}, nil
}
