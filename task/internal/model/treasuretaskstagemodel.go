package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TreasureTaskStageModel = (*customTreasureTaskStageModel)(nil)

type (
	// TreasureTaskStageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTreasureTaskStageModel.
	TreasureTaskStageModel interface {
		treasureTaskStageModel
	}

	customTreasureTaskStageModel struct {
		*defaultTreasureTaskStageModel
	}
)

// NewTreasureTaskStageModel returns a model for the database table.
func NewTreasureTaskStageModel(conn sqlx.SqlConn) TreasureTaskStageModel {
	return &customTreasureTaskStageModel{
		defaultTreasureTaskStageModel: newTreasureTaskStageModel(conn),
	}
}
