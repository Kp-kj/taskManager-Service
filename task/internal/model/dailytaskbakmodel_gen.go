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
	dailyTaskBakFieldNames          = builder.RawFieldNames(&DailyTaskBak{})
	dailyTaskBakRows                = strings.Join(dailyTaskBakFieldNames, ",")
	dailyTaskBakRowsExpectAutoSet   = strings.Join(stringx.Remove(dailyTaskBakFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dailyTaskBakRowsWithPlaceHolder = strings.Join(stringx.Remove(dailyTaskBakFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dailyTaskBakModel interface {
		Insert(ctx context.Context, data *DailyTaskBak) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DailyTaskBak, error)
		Update(ctx context.Context, data *DailyTaskBak) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDailyTaskBakModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DailyTaskBak struct {
		Id         int64          `db:"id"`
		CreatedAt  sql.NullTime   `db:"created_at"`
		UpdatedAt  sql.NullTime   `db:"updated_at"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
		UserId     sql.NullString `db:"user_id"`
		TaskId     sql.NullInt64  `db:"task_id"`
		Complete   sql.NullInt64  `db:"complete"`
		Experience sql.NullInt64  `db:"experience"`
	}
)

func newDailyTaskBakModel(conn sqlx.SqlConn) *defaultDailyTaskBakModel {
	return &defaultDailyTaskBakModel{
		conn:  conn,
		table: "`daily_task_bak`",
	}
}

func (m *defaultDailyTaskBakModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDailyTaskBakModel) FindOne(ctx context.Context, id int64) (*DailyTaskBak, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dailyTaskBakRows, m.table)
	var resp DailyTaskBak
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

func (m *defaultDailyTaskBakModel) Insert(ctx context.Context, data *DailyTaskBak) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, dailyTaskBakRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.TaskId, data.Complete, data.Experience)
	return ret, err
}

func (m *defaultDailyTaskBakModel) Update(ctx context.Context, data *DailyTaskBak) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dailyTaskBakRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.TaskId, data.Complete, data.Experience, data.Id)
	return err
}

func (m *defaultDailyTaskBakModel) tableName() string {
	return m.table
}
