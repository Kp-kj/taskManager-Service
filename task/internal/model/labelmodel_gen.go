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
	labelFieldNames          = builder.RawFieldNames(&Label{})
	labelRows                = strings.Join(labelFieldNames, ",")
	labelRowsExpectAutoSet   = strings.Join(stringx.Remove(labelFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	labelRowsWithPlaceHolder = strings.Join(stringx.Remove(labelFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	labelModel interface {
		Insert(ctx context.Context, data *Label) (sql.Result, error)
		FindOne(ctx context.Context, content,creator string) (*Label, error)
		Update(ctx context.Context, data *Label) error
		Delete(ctx context.Context, id int64) error
		FindLabelAmount(ctx context.Context,creator string) (int64, error)
		FindList(ctx context.Context,creator string) ([]*Label, error)
	}

	defaultLabelModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Label struct {
		Id      int64  `db:"id"`
		Creator string `db:"creator"`
		Content string `db:"content"`
	}
)

func newLabelModel(conn sqlx.SqlConn) *defaultLabelModel {
	return &defaultLabelModel{
		conn:  conn,
		table: "`label`",
	}
}

func (m *defaultLabelModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLabelModel) FindOne(ctx context.Context, content,creator string) (*Label, error) {
	query := fmt.Sprintf("select %s from %s where `content` = ? AND `creator` = ? limit 1", labelRows, m.table)
	var resp Label
	err := m.conn.QueryRowCtx(ctx, &resp, query, content,creator)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// 获取标签列表
func (m *defaultLabelModel) FindList(ctx context.Context, creator string) ([]*Label, error) {
	query := fmt.Sprintf("select %s from %s where `creator` = ?", labelRows, m.table)
	var resp []*Label
	err := m.conn.QueryRowsCtx(ctx,&resp, query, creator)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// FindLabelAmount 获取标签数量
func (m *defaultLabelModel) FindLabelAmount(ctx context.Context, creator string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE creator = ?",  m.table)
	var amount int64
	err := m.conn.QueryRow(&amount, query, creator)
	switch err {
	case nil:
		return amount, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *defaultLabelModel) Insert(ctx context.Context, data *Label) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, labelRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Creator, data.Content)
	return ret, err
}

func (m *defaultLabelModel) Update(ctx context.Context, data *Label) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, labelRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Creator, data.Content, data.Id)
	return err
}

func (m *defaultLabelModel) tableName() string {
	return m.table
}
