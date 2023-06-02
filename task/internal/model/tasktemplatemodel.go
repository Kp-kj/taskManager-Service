package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TaskTemplateModel = (*customTaskTemplateModel)(nil)

type (
	// TaskTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskTemplateModel.
	TaskTemplateModel interface {
		taskTemplateModel
	}

	customTaskTemplateModel struct {
		*defaultTaskTemplateModel
	}
)

// NewTaskTemplateModel returns a model for the database table.
func NewTaskTemplateModel(conn sqlx.SqlConn) TaskTemplateModel {
	return &customTaskTemplateModel{
		defaultTaskTemplateModel: newTaskTemplateModel(conn),
	}
}
