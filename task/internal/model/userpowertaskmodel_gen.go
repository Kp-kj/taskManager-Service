// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userPowerTaskFieldNames          = builder.RawFieldNames(&UserPowerTask{})
	userPowerTaskRows                = strings.Join(userPowerTaskFieldNames, ",")
	userPowerTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(userPowerTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userPowerTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(userPowerTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	userPowerTaskModel interface {
		Insert(ctx context.Context, data *UserPowerTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserPowerTask, error)
		Update(ctx context.Context, data *UserPowerTask) error
		Delete(ctx context.Context, id int64) error
		FindNumberHelpings(ctx context.Context, userId string) (int64, error)
	}

	defaultUserPowerTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserPowerTask struct {
		Id              int64        `db:"id"`
		CreatedAt       sql.NullTime `db:"created_at"`
		UpdatedAt       sql.NullTime `db:"updated_at"`
		DeletedAt       sql.NullTime `db:"deleted_at"`
		PublishesUserId string       `db:"publishes_user_id"`
		HelperUserId    string       `db:"helper_user_id"`
	}
)

func newUserPowerTaskModel(conn sqlx.SqlConn) *defaultUserPowerTaskModel {
	return &defaultUserPowerTaskModel{
		conn:  conn,
		table: "`user_power_task`",
	}
}

func (m *defaultUserPowerTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserPowerTaskModel) FindOne(ctx context.Context, id int64) (*UserPowerTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userPowerTaskRows, m.table)
	var resp UserPowerTask
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
// FindNumberHelpings 获取用户帮助助力次数
func (m *defaultUserPowerTaskModel) FindNumberHelpings(ctx context.Context, userId string) (int64, error) {
	query := fmt.Sprintf("select COUNT(*) from %s where `helper_user_id` = ? AND `created_at` like ? ", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, time.Now().Format("2006-01-02")+"%")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultUserPowerTaskModel) Insert(ctx context.Context, data *UserPowerTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userPowerTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.PublishesUserId, data.HelperUserId)
	return ret, err
}

func (m *defaultUserPowerTaskModel) Update(ctx context.Context, data *UserPowerTask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userPowerTaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.PublishesUserId, data.HelperUserId, data.Id)
	return err
}

func (m *defaultUserPowerTaskModel) tableName() string {
	return m.table
}