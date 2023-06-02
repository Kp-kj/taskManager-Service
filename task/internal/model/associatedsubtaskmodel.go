package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AssociatedSubtaskModel = (*customAssociatedSubtaskModel)(nil)

type (
	// AssociatedSubtaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssociatedSubtaskModel.
	AssociatedSubtaskModel interface {
		associatedSubtaskModel
	}

	customAssociatedSubtaskModel struct {
		*defaultAssociatedSubtaskModel
	}
)

// NewAssociatedSubtaskModel returns a model for the database table.
func NewAssociatedSubtaskModel(conn sqlx.SqlConn) AssociatedSubtaskModel {
	return &customAssociatedSubtaskModel{
		defaultAssociatedSubtaskModel: newAssociatedSubtaskModel(conn),
	}
}
