package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PublishTaskModel = (*customPublishTaskModel)(nil)

type (
	// PublishTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPublishTaskModel.
	PublishTaskModel interface {
		publishTaskModel
	}

	customPublishTaskModel struct {
		*defaultPublishTaskModel
	}
)

// NewPublishTaskModel returns a model for the database table.
func NewPublishTaskModel(conn sqlx.SqlConn) PublishTaskModel {
	return &customPublishTaskModel{
		defaultPublishTaskModel: newPublishTaskModel(conn),
	}
}
