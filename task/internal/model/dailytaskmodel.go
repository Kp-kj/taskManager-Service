package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DailyTaskModel = (*customDailyTaskModel)(nil)

type (
	// DailyTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDailyTaskModel.
	DailyTaskModel interface {
		dailyTaskModel
	}

	customDailyTaskModel struct {
		*defaultDailyTaskModel
	}
)

// NewDailyTaskModel returns a model for the database table.
func NewDailyTaskModel(conn sqlx.SqlConn) DailyTaskModel {
	return &customDailyTaskModel{
		defaultDailyTaskModel: newDailyTaskModel(conn),
	}
}
