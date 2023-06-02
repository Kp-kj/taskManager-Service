package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TreasureStagesModel = (*customTreasureStagesModel)(nil)

type (
	// TreasureStagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTreasureStagesModel.
	TreasureStagesModel interface {
		treasureStagesModel
	}

	customTreasureStagesModel struct {
		*defaultTreasureStagesModel
	}
)

// NewTreasureStagesModel returns a model for the database table.
func NewTreasureStagesModel(conn sqlx.SqlConn) TreasureStagesModel {
	return &customTreasureStagesModel{
		defaultTreasureStagesModel: newTreasureStagesModel(conn),
	}
}
