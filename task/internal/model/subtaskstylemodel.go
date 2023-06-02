package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SubtaskStyleModel = (*customSubtaskStyleModel)(nil)

type (
	// SubtaskStyleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubtaskStyleModel.
	SubtaskStyleModel interface {
		subtaskStyleModel
	}

	customSubtaskStyleModel struct {
		*defaultSubtaskStyleModel
	}
)

// NewSubtaskStyleModel returns a model for the database table.
func NewSubtaskStyleModel(conn sqlx.SqlConn) SubtaskStyleModel {
	return &customSubtaskStyleModel{
		defaultSubtaskStyleModel: newSubtaskStyleModel(conn),
	}
}
