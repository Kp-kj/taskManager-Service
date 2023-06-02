package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ParticipantModel = (*customParticipantModel)(nil)

type (
	// ParticipantModel is an interface to be customized, add more methods here,
	// and implement the added methods in customParticipantModel.
	ParticipantModel interface {
		participantModel
	}

	customParticipantModel struct {
		*defaultParticipantModel
	}
)

// NewParticipantModel returns a model for the database table.
func NewParticipantModel(conn sqlx.SqlConn) ParticipantModel {
	return &customParticipantModel{
		defaultParticipantModel: newParticipantModel(conn),
	}
}
