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
	chestCollectionsFieldNames          = builder.RawFieldNames(&ChestCollections{})
	chestCollectionsRows                = strings.Join(chestCollectionsFieldNames, ",")
	chestCollectionsRowsExpectAutoSet   = strings.Join(stringx.Remove(chestCollectionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chestCollectionsRowsWithPlaceHolder = strings.Join(stringx.Remove(chestCollectionsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	chestCollectionsModel interface {
		Insert(ctx context.Context, data *ChestCollections) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChestCollections, error)
		Update(ctx context.Context, data *ChestCollections) error
		Delete(ctx context.Context, id int64) error
		FindIndividualAllowance(ctx context.Context, userID string) (*ChestCollections, error)
		UpdateIndividualAllowance(ctx context.Context,userID string, amount int64) error
	}

	defaultChestCollectionsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ChestCollections struct {
		Id          int64         `db:"id"`
		CreatedAt   sql.NullTime  `db:"created_at"`
		UpdatedAt   sql.NullTime  `db:"updated_at"`
		DeletedAt   sql.NullTime  `db:"deleted_at"`
		UserId      string        `db:"user_id"`
		ChestAmount sql.NullInt64 `db:"chest_amount"`
	}
)

func newChestCollectionsModel(conn sqlx.SqlConn) *defaultChestCollectionsModel {
	return &defaultChestCollectionsModel{
		conn:  conn,
		table: "`chest_collections`",
	}
}

func (m *defaultChestCollectionsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultChestCollectionsModel) FindOne(ctx context.Context, id int64) (*ChestCollections, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chestCollectionsRows, m.table)
	var resp ChestCollections
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

func (m *defaultChestCollectionsModel) Insert(ctx context.Context, data *ChestCollections) (sql.Result, error) {
	chestCollections:= fmt.Sprintf("created_at, user_id, chest_amount")
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, chestCollections)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreatedAt, data.UserId, data.ChestAmount)
	return ret, err
}

func (m *defaultChestCollectionsModel) Update(ctx context.Context, data *ChestCollections) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chestCollectionsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.ChestAmount, data.Id)
	return err
}

// FindIndividualAllowance 查询个人宝箱领取度
func (m *defaultChestCollectionsModel) FindIndividualAllowance(ctx context.Context, userID string) (*ChestCollections, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? AND `created_at` like ?", chestCollectionsRows, m.table)
	var resp ChestCollections
	err := m.conn.QueryRowCtx(ctx, &resp, query, userID, time.Now().Format("2006-01-02")+"%")
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// UpdateIndividualAllowance 更新个人宝箱领取数
func (m *defaultChestCollectionsModel) UpdateIndividualAllowance(ctx context.Context,userID string, amount int64) error {
	query := fmt.Sprintf("update %s set `chest_amount` = chest_amount + ? where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, amount, userID)
	return err
}

func (m *defaultChestCollectionsModel) tableName() string {
	return m.table
}
