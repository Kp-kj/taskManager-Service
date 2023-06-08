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
	dailyTaskFieldNames          = builder.RawFieldNames(&DailyTask{})
	dailyTaskRows                = strings.Join(dailyTaskFieldNames, ",")
	dailyTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(dailyTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	dailyTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(dailyTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	dailyTaskModel interface {
		Insert(ctx context.Context, data *DailyTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DailyTask, error)
		Update(ctx context.Context,userId string,taskId int64, complete ,experience int64) error
		Delete(ctx context.Context, id int64) error
		FindCompletionTask(ctx context.Context, userId string,taskId int64) (*DailyTask, error)
	}

	defaultDailyTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DailyTask struct {
		Id         int64         `db:"id"`
		CreatedAt  sql.NullTime  `db:"created_at"`
		UpdatedAt  sql.NullTime  `db:"updated_at"`
		DeletedAt  sql.NullTime  `db:"deleted_at"`
		UserId     string        `db:"user_id"`
		TaskId     int64         `db:"task_id"`
		Complete   sql.NullInt64 `db:"complete"`
		Experience sql.NullInt64 `db:"experience"`
	}
)

func newDailyTaskModel(conn sqlx.SqlConn) *defaultDailyTaskModel {
	return &defaultDailyTaskModel{
		conn:  conn,
		table: "`daily_task`",
	}
}

func (m *defaultDailyTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDailyTaskModel) FindOne(ctx context.Context, id int64) (*DailyTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dailyTaskRows, m.table)
	var resp DailyTask
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

// 获取日常任务完成信息
func (m *defaultDailyTaskModel) FindCompletionTask(ctx context.Context, userId string,taskId int64) (*DailyTask, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? AND `task_id` = ?", dailyTaskRows, m.table)
	var resp DailyTask
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId,taskId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDailyTaskModel) Insert(ctx context.Context, data *DailyTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, dailyTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.TaskId, data.Complete, data.Experience)
	return ret, err
}

func (m *defaultDailyTaskModel) Update(ctx context.Context,userId string,taskId int64, complete ,experience int64) error {
	query := fmt.Sprintf("update %s set `complete` = ?,`experience` = ? where `userId` = ? AND `taskId` = ? ", m.table)
	_, err := m.conn.ExecCtx(ctx, query,complete,experience,userId,taskId,)
	return err
}

func (m *defaultDailyTaskModel) tableName() string {
	return m.table
}
