package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserPublishesHelperTaskModel = (*customUserPublishesHelperTaskModel)(nil)

type (
	// UserPublishesHelperTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPublishesHelperTaskModel.
	UserPublishesHelperTaskModel interface {
		userPublishesHelperTaskModel
	}

	customUserPublishesHelperTaskModel struct {
		*defaultUserPublishesHelperTaskModel
	}
)

// NewUserPublishesHelperTaskModel returns a model for the database table.
func NewUserPublishesHelperTaskModel(conn sqlx.SqlConn) UserPublishesHelperTaskModel {
	return &customUserPublishesHelperTaskModel{
		defaultUserPublishesHelperTaskModel: newUserPublishesHelperTaskModel(conn),
	}
}
