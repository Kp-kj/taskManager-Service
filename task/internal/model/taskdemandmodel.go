package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TaskDemandModel = (*customTaskDemandModel)(nil)

type (
	// TaskDemandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskDemandModel.
	TaskDemandModel interface {
		taskDemandModel
	}

	customTaskDemandModel struct {
		*defaultTaskDemandModel
	}
)

// NewTaskDemandModel returns a model for the database table.
func NewTaskDemandModel(conn sqlx.SqlConn) TaskDemandModel {
	return &customTaskDemandModel{
		defaultTaskDemandModel: newTaskDemandModel(conn),
	}
}
