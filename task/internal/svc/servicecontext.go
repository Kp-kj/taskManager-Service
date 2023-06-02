package svc

import (
	"task/internal/config"
	"task/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	PublishTaskModel    model.PublishTaskModel    // 创建策展任务
	TaskDemandModel     model.TaskDemandModel     // 创建任务详情
	LabelModel          model.LabelModel          // 创建标签
	TreasureStagesModel model.TreasureStagesModel // 参与者任务要求完成度
	ParticipantModel    model.ParticipantModel    // 获取任务参与者列表
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		PublishTaskModel:    model.NewPublishTaskModel(sqlx.NewMysql(c.DB.DataSource)),
		TaskDemandModel:     model.NewTaskDemandModel(sqlx.NewMysql(c.DB.DataSource)),
		LabelModel:          model.NewLabelModel(sqlx.NewMysql(c.DB.DataSource)),
		TreasureStagesModel: model.NewTreasureStagesModel(sqlx.NewMysql(c.DB.DataSource)),
		ParticipantModel:    model.NewParticipantModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
