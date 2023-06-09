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
	treasureStagesFieldNames          = builder.RawFieldNames(&TreasureStages{})
	treasureStagesRows                = strings.Join(treasureStagesFieldNames, ",")
	treasureStagesRowsExpectAutoSet   = strings.Join(stringx.Remove(treasureStagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	treasureStagesRowsWithPlaceHolder = strings.Join(stringx.Remove(treasureStagesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	treasureStagesModel interface {
		Insert(ctx context.Context, data *TreasureStages) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TreasureStages, error)
		Update(ctx context.Context, data *TreasureStages) error
		Delete(ctx context.Context, id int64) error
		FindSpecificTask(ctx context.Context,userId string,taskId uint,taskName int) (*TreasureStages, error)
		FindPersonalMissionCompletionDetails(ctx context.Context, userId string,taskId int64) ([]*TreasureStages, error)
	}

	defaultTreasureStagesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TreasureStages struct {
		Id         int64         `db:"id"`
		CreatedAt  sql.NullTime  `db:"created_at"`
		UpdatedAt  sql.NullTime  `db:"updated_at"`
		DeletedAt  sql.NullTime  `db:"deleted_at"`
		UserId     string        `db:"user_id"`
		TaskId     int64         `db:"task_id"`
		TaskName   sql.NullInt64 `db:"task_name"`
		TaskStatus sql.NullInt64 `db:"task_status"`
	}
)

func newTreasureStagesModel(conn sqlx.SqlConn) *defaultTreasureStagesModel {
	return &defaultTreasureStagesModel{
		conn:  conn,
		table: "`treasure_stages`",
	}
}

func (m *defaultTreasureStagesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTreasureStagesModel) FindOne(ctx context.Context, id int64) (*TreasureStages, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", treasureStagesRows, m.table)
	var resp TreasureStages
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

// FindSpecificTask 查询特定任务
func (m *defaultTreasureStagesModel) FindSpecificTask(ctx context.Context, userId string,taskId uint,taskName int) (*TreasureStages, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? AND `task_id` = ? AND `task_name` = ?", treasureStagesRows, m.table)
	var resp TreasureStages
	err := m.conn.QueryRow(&resp, query, userId,taskId,taskName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTreasureStagesModel) Insert(ctx context.Context, data *TreasureStages) (sql.Result, error) {
	treasureStages := fmt.Sprintf("created_at, user_id, task_id, task_name,task_status")
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, treasureStages)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreatedAt, data.UserId, data.TaskId, data.TaskName, data.TaskStatus)
	return ret, err
}

func (m *defaultTreasureStagesModel) Update(ctx context.Context, data *TreasureStages) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, treasureStagesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.TaskId, data.TaskName, data.TaskStatus, data.Id)
	return err
}

func (m *defaultTreasureStagesModel) tableName() string {
	return m.table
}


// FindPersonalMissionCompletionDetails 获取个人任务完成详情
func (m *defaultTreasureStagesModel) FindPersonalMissionCompletionDetails(ctx context.Context, userId string,taskId int64) ([]*TreasureStages, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? AND `task_id` = ?", treasureStagesRows, m.table)
	var resp []*TreasureStages
	err := m.conn.QueryRows(&resp, query, userId,taskId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
