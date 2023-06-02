package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DailyTaskBakModel = (*customDailyTaskBakModel)(nil)

type (
	// DailyTaskBakModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDailyTaskBakModel.
	DailyTaskBakModel interface {
		dailyTaskBakModel
	}

	customDailyTaskBakModel struct {
		*defaultDailyTaskBakModel
	}
)

// NewDailyTaskBakModel returns a model for the database table.
func NewDailyTaskBakModel(conn sqlx.SqlConn) DailyTaskBakModel {
	return &customDailyTaskBakModel{
		defaultDailyTaskBakModel: newDailyTaskBakModel(conn),
	}
}
