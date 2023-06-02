package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChestCollectionsModel = (*customChestCollectionsModel)(nil)

type (
	// ChestCollectionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChestCollectionsModel.
	ChestCollectionsModel interface {
		chestCollectionsModel
	}

	customChestCollectionsModel struct {
		*defaultChestCollectionsModel
	}
)

// NewChestCollectionsModel returns a model for the database table.
func NewChestCollectionsModel(conn sqlx.SqlConn) ChestCollectionsModel {
	return &customChestCollectionsModel{
		defaultChestCollectionsModel: newChestCollectionsModel(conn),
	}
}
