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
	taskDemandFieldNames          = builder.RawFieldNames(&TaskDemand{})
	taskDemandRows                = strings.Join(taskDemandFieldNames, ",")
	taskDemandRowsExpectAutoSet   = strings.Join(stringx.Remove(taskDemandFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	taskDemandRowsWithPlaceHolder = strings.Join(stringx.Remove(taskDemandFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	taskDemandModel interface {
		Insert(ctx context.Context, data *TaskDemand) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TaskDemand, error)
		Update(ctx context.Context, data *TaskDemand) error
		Delete(ctx context.Context, id int64) error
		FindList(ctx context.Context,id int64) ([]*TaskDemand, error)

		FindTaskDemand(ctx context.Context, taskId int64,taskName string) ([]*TaskDemand, error)
	}

	defaultTaskDemandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TaskDemand struct {
		Id         int64          `db:"id"`
		CreatedAt  sql.NullTime   `db:"created_at"`
		UpdatedAt  sql.NullTime   `db:"updated_at"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
		TaskId     sql.NullInt64  `db:"task_id"`
		TaskName   sql.NullInt64  `db:"task_name"`
		TaskDemand sql.NullString `db:"task_demand"`
		Article    sql.NullString `db:"article"`
	}
)

func newTaskDemandModel(conn sqlx.SqlConn) *defaultTaskDemandModel {
	return &defaultTaskDemandModel{
		conn:  conn,
		table: "`task_demand`",
	}
}

func (m *defaultTaskDemandModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTaskDemandModel) FindOne(ctx context.Context, id int64) (*TaskDemand, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskDemandRows, m.table)
	var resp TaskDemand
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
// FindList 查询任务要求列表
func (m *defaultTaskDemandModel) FindList(ctx context.Context, id int64) ([]*TaskDemand, error) {
	query := fmt.Sprintf("select %s from %s where `task_id` = ? ", taskDemandRows, m.table)
	var resp []*TaskDemand
	err := m.conn.QueryRows(&resp, query, id)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskDemandModel) Insert(ctx context.Context, data *TaskDemand) (sql.Result, error) {
	taskDemand:="`created_at`,`task_id`,`task_name`,`task_demand`,`article`"	
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, taskDemand)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreatedAt, data.TaskId, data.TaskName, data.TaskDemand, data.Article)
	return ret, err
}

func (m *defaultTaskDemandModel) Update(ctx context.Context, data *TaskDemand) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskDemandRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.TaskId, data.TaskName, data.TaskDemand, data.Article, data.Id)
	return err
}

// FindList 查询任务要求
func (m *defaultTaskDemandModel) FindTaskDemand(ctx context.Context, taskId int64,taskName string) ([]*TaskDemand, error) {
	query := fmt.Sprintf("select %s from %s where `task_id` = ? AND `task_name` = ?", taskDemandRows, m.table)
	var resp []*TaskDemand
	err := m.conn.QueryRows(&resp, query, taskId,taskName)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskDemandModel) tableName() string {
	return m.table
}
