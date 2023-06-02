package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TreasureTaskModel = (*customTreasureTaskModel)(nil)

type (
	// TreasureTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTreasureTaskModel.
	TreasureTaskModel interface {
		treasureTaskModel
	}

	customTreasureTaskModel struct {
		*defaultTreasureTaskModel
	}
)

// NewTreasureTaskModel returns a model for the database table.
func NewTreasureTaskModel(conn sqlx.SqlConn) TreasureTaskModel {
	return &customTreasureTaskModel{
		defaultTreasureTaskModel: newTreasureTaskModel(conn),
	}
}
