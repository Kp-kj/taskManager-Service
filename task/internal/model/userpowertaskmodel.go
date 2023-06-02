package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserPowerTaskModel = (*customUserPowerTaskModel)(nil)

type (
	// UserPowerTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPowerTaskModel.
	UserPowerTaskModel interface {
		userPowerTaskModel
	}

	customUserPowerTaskModel struct {
		*defaultUserPowerTaskModel
	}
)

// NewUserPowerTaskModel returns a model for the database table.
func NewUserPowerTaskModel(conn sqlx.SqlConn) UserPowerTaskModel {
	return &customUserPowerTaskModel{
		defaultUserPowerTaskModel: newUserPowerTaskModel(conn),
	}
}
