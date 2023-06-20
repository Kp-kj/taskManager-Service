// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	treasureTaskStageFieldNames          = builder.RawFieldNames(&TreasureTaskStage{})
	treasureTaskStageRows                = strings.Join(treasureTaskStageFieldNames, ",")
	treasureTaskStageRowsExpectAutoSet   = strings.Join(stringx.Remove(treasureTaskStageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	treasureTaskStageRowsWithPlaceHolder = strings.Join(stringx.Remove(treasureTaskStageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	treasureTaskStageModel interface {
		Insert(ctx context.Context, data *TreasureTaskStage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TreasureTaskStage, error)
		Update(ctx context.Context, data *TreasureTaskStage) error
		Delete(ctx context.Context, id int64) error
		UpdateTreasureTaskStage(ctx context.Context, treasure, treasureSequence, stageExperience, stageReward int64) error
		FindTreasureInformation(ctx context.Context, id int64,genre string) ([]*TreasureTaskStage, error)
	}

	defaultTreasureTaskStageModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TreasureTaskStage struct {
		Id               int64 `db:"id"`
		Treasure         int64 `db:"treasure"`
		TreasureSequence int64 `db:"treasure_sequence"`
		StageExperience  int64 `db:"stage_experience"`
		StageReward      int64 `db:"stage_reward"`
	}
)

func newTreasureTaskStageModel(conn sqlx.SqlConn) *defaultTreasureTaskStageModel {
	return &defaultTreasureTaskStageModel{
		conn:  conn,
		table: "`treasure_task_stage`",
	}
}

func (m *defaultTreasureTaskStageModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTreasureTaskStageModel) FindOne(ctx context.Context, id int64) (*TreasureTaskStage, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", treasureTaskStageRows, m.table)
	var resp TreasureTaskStage
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTreasureTaskStageModel) Insert(ctx context.Context, data *TreasureTaskStage) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, treasureTaskStageRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Treasure, data.TreasureSequence, data.StageExperience, data.StageReward)
	return ret, err
}

func (m *defaultTreasureTaskStageModel) Update(ctx context.Context, data *TreasureTaskStage) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, treasureTaskStageRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Treasure, data.TreasureSequence, data.StageExperience, data.StageReward, data.Id)
	return err
}

// 编辑宝箱样式
func (m *defaultTreasureTaskStageModel) UpdateTreasureTaskStage(ctx context.Context,treasure,treasureSequence int64,stageExperience,stageReward int64) error {
	query := fmt.Sprintf("update %s set `stage_experience` = ?,`stage_reward` = ? where `treasure` = ? AND `treasure_sequence` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query,stageExperience,stageReward,treasure,treasureSequence)
	return err
}

func (m *defaultTreasureTaskStageModel) FindTreasureInformation(ctx context.Context, id int64,genre string) ([]*TreasureTaskStage, error) {
	query := fmt.Sprintf("select %s from %s where `%s` = ?", treasureTaskStageRows, m.table,genre)
	var resp []*TreasureTaskStage
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultTreasureTaskStageModel) tableName() string {
	return m.table
}
