package svc

import (
	"taskManager-Service/internal/config"
	"taskManager-Service/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                       config.Config
	PublishTaskModel             model.PublishTaskModel             // 创建策展任务
	TaskDemandModel              model.TaskDemandModel              // 创建任务详情
	LabelModel                   model.LabelModel                   // 创建标签
	TreasureStagesModel          model.TreasureStagesModel          // 参与者任务要求完成度
	ParticipantModel             model.ParticipantModel             // 获取任务参与者列表
	AssociatedSubtaskModel       model.AssociatedSubtaskModel       // 关联子任务
	ChestCollectionsModel        model.ChestCollectionsModel        // 宝箱领取度
	DailyTaskBakModel            model.DailyTaskBakModel            // 日常任务完成表(冷表,存储30天)
	DailyTaskModel               model.DailyTaskModel               // 日常任务完成表
	SubtaskStyleModel            model.SubtaskStyleModel            // 子任务样式表
	TaskTemplateModel            model.TaskTemplateModel            // 任务要求模板
	TreasureTaskModel            model.TreasureTaskModel            // 宝箱样式
	UserPowerTaskModel           model.UserPowerTaskModel           // 用户帮助助力任务表
	UserPublishesHelperTaskModel model.UserPublishesHelperTaskModel // 用户发布助力任务表
	TreasureTaskStageModel       model.TreasureTaskStageModel       // 宝箱阶段
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                       c,
		PublishTaskModel:             model.NewPublishTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		TaskDemandModel:              model.NewTaskDemandModel(sqlx.NewMysql(c.DB.DataSource)),
		LabelModel:                   model.NewLabelModel(sqlx.NewMysql(c.DB.DataSource)),
		TreasureStagesModel:          model.NewTreasureStagesModel(sqlx.NewMysql(c.DB.DataSource)),
		ParticipantModel:             model.NewParticipantModel(sqlx.NewMysql(c.DB.DataSource)),
		AssociatedSubtaskModel:       model.NewAssociatedSubtaskModel(sqlx.NewMysql(c.DB.DataSource)),
		ChestCollectionsModel:        model.NewChestCollectionsModel(sqlx.NewMysql(c.DB.DataSource)),
		DailyTaskBakModel:            model.NewDailyTaskBakModel(sqlx.NewMysql(c.DB.DataSource)),
		DailyTaskModel:               model.NewDailyTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		SubtaskStyleModel:            model.NewSubtaskStyleModel(sqlx.NewMysql(c.DB.DataSource)),
		TaskTemplateModel:            model.NewTaskTemplateModel(sqlx.NewMysql(c.DB.DataSource)),
		TreasureTaskModel:            model.NewTreasureTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		UserPowerTaskModel:           model.NewUserPowerTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		UserPublishesHelperTaskModel: model.NewUserPublishesHelperTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		TreasureTaskStageModel:       model.NewTreasureTaskStageModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
