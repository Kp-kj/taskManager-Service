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
	associatedSubtaskFieldNames          = builder.RawFieldNames(&AssociatedSubtask{})
	associatedSubtaskRows                = strings.Join(associatedSubtaskFieldNames, ",")
	associatedSubtaskRowsExpectAutoSet   = strings.Join(stringx.Remove(associatedSubtaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	associatedSubtaskRowsWithPlaceHolder = strings.Join(stringx.Remove(associatedSubtaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	associatedSubtaskModel interface {
		Insert(ctx context.Context, data *AssociatedSubtask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AssociatedSubtask, error)
		Update(ctx context.Context, data *AssociatedSubtask) error
		Delete(ctx context.Context, id int64) error
		FindAssociatedSubtask(ctx context.Context, id int64,genre string) ([]*AssociatedSubtask, error)
		FindSubtaskInformation(ctx context.Context,subtaskName string,treasureId int64) (*AssociatedSubtask, error)
	}

	defaultAssociatedSubtaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AssociatedSubtask struct {
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
		Reward         sql.NullInt64  `db:"reward"`
		Experience     sql.NullInt64  `db:"experience"`
		Number         sql.NullInt64  `db:"number"`
		Article        sql.NullString `db:"article"`
		Link           sql.NullString `db:"link"`
		Label          sql.NullString `db:"label"`
		TreasureId     int64          `db:"treasure_id"`
	}
)

func newAssociatedSubtaskModel(conn sqlx.SqlConn) *defaultAssociatedSubtaskModel {
	return &defaultAssociatedSubtaskModel{
		conn:  conn,
		table: "`associated_subtask`",
	}
}

func (m *defaultAssociatedSubtaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssociatedSubtaskModel) FindOne(ctx context.Context, id int64) (*AssociatedSubtask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", associatedSubtaskRows, m.table)
	var resp AssociatedSubtask
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
// FindAssociatedSubtask 查询是否存在已关联子任务
func (m *defaultAssociatedSubtaskModel) FindAssociatedSubtask(ctx context.Context, id int64,genre string) ([]*AssociatedSubtask, error) {
	query := fmt.Sprintf("select %s from %s where %s = ?", associatedSubtaskRows, m.table,genre)
	var resp []*AssociatedSubtask
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// 获取相应的子任务信息
func (m *defaultAssociatedSubtaskModel) FindSubtaskInformation(ctx context.Context,subtaskName string,treasureId int64) (*AssociatedSubtask, error) {
	query := fmt.Sprintf("select %s from %s where `task_name` = ? AND `treasure_id` = ? limit 1", associatedSubtaskRows, m.table)
	var resp AssociatedSubtask
	err := m.conn.QueryRowCtx(ctx, &resp, query, subtaskName,treasureId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAssociatedSubtaskModel) Insert(ctx context.Context, data *AssociatedSubtask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, associatedSubtaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.TaskId, data.TaskName, data.TaskNameEng, data.TaskDetails, data.TaskDetailsEng, data.TaskStatus, data.Reward, data.Experience, data.Number, data.Article, data.Link, data.Label, data.TreasureId)
	return ret, err
}

func (m *defaultAssociatedSubtaskModel) Update(ctx context.Context, data *AssociatedSubtask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, associatedSubtaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.TaskId, data.TaskName, data.TaskNameEng, data.TaskDetails, data.TaskDetailsEng, data.TaskStatus, data.Reward, data.Experience, data.Number, data.Article, data.Link, data.Label, data.TreasureId, data.Id)
	return err
}

func (m *defaultAssociatedSubtaskModel) tableName() string {
	return m.table
}
