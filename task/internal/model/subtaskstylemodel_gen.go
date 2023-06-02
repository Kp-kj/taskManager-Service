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
	subtaskStyleFieldNames          = builder.RawFieldNames(&SubtaskStyle{})
	subtaskStyleRows                = strings.Join(subtaskStyleFieldNames, ",")
	subtaskStyleRowsExpectAutoSet   = strings.Join(stringx.Remove(subtaskStyleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	subtaskStyleRowsWithPlaceHolder = strings.Join(stringx.Remove(subtaskStyleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	subtaskStyleModel interface {
		Insert(ctx context.Context, data *SubtaskStyle) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SubtaskStyle, error)
		Update(ctx context.Context, data *SubtaskStyle) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSubtaskStyleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SubtaskStyle struct {
		Id             int64          `db:"id"`
		CreatedAt      sql.NullTime   `db:"created_at"`
		UpdatedAt      sql.NullTime   `db:"updated_at"`
		DeletedAt      sql.NullTime   `db:"deleted_at"`
		TaskId         sql.NullInt64  `db:"task_id"`
		TaskName       sql.NullString `db:"task_name"`
		TaskNameEng    sql.NullString `db:"task_name_eng"`
		TaskDetails    sql.NullString `db:"task_details"`
		TaskDetailsEng sql.NullString `db:"task_details_eng"`
		TaskStatus     sql.NullInt64  `db:"task_status"`
	}
)

func newSubtaskStyleModel(conn sqlx.SqlConn) *defaultSubtaskStyleModel {
	return &defaultSubtaskStyleModel{
		conn:  conn,
		table: "`subtask_style`",
	}
}

func (m *defaultSubtaskStyleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSubtaskStyleModel) FindOne(ctx context.Context, id int64) (*SubtaskStyle, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", subtaskStyleRows, m.table)
	var resp SubtaskStyle
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

func (m *defaultSubtaskStyleModel) Insert(ctx context.Context, data *SubtaskStyle) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, subtaskStyleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.TaskId, data.TaskName, data.TaskNameEng, data.TaskDetails, data.TaskDetailsEng, data.TaskStatus)
	return ret, err
}

func (m *defaultSubtaskStyleModel) Update(ctx context.Context, data *SubtaskStyle) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, subtaskStyleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.TaskId, data.TaskName, data.TaskNameEng, data.TaskDetails, data.TaskDetailsEng, data.TaskStatus, data.Id)
	return err
}

func (m *defaultSubtaskStyleModel) tableName() string {
	return m.table
}
