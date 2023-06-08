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
	treasureTaskFieldNames          = builder.RawFieldNames(&TreasureTask{})
	treasureTaskRows                = strings.Join(treasureTaskFieldNames, ",")
	treasureTaskRowsExpectAutoSet   = strings.Join(stringx.Remove(treasureTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	treasureTaskRowsWithPlaceHolder = strings.Join(stringx.Remove(treasureTaskFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	treasureTaskModel interface {
		Insert(ctx context.Context, data *TreasureTask) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TreasureTask, error)
		Update(ctx context.Context, data *TreasureTask) error
		Delete(ctx context.Context, id int64) error
		UpdateOldStyle(ctx context.Context, data uint64) error
		UpdateNewStyle(ctx context.Context, data uint64) error
		FindTreasureQuantity(ctx context.Context) (*TreasureTask, error)
		FindNameAndQuantity(ctx context.Context,reward int64,name string,currPage, maxNum int64) ([]*TreasureTask, error)
		FindNameAndQuantityCount(ctx context.Context,reward int64,name string) (int64, error)
		FindAll(ctx context.Context,currPage, maxNum int64) ([]*TreasureTask, error)
		FindAllCount(ctx context.Context) (int64, error)
		FindName(ctx context.Context,name string,currPage, maxNum int64) ([]*TreasureTask, error)
		FindNameCount(ctx context.Context,name string) (int64, error)
		FindQuantity(ctx context.Context,reward int64,currPage, maxNum int64) ([]*TreasureTask, error)
		FindQuantityCount(ctx context.Context,reward int64) (int64, error)

	}

	defaultTreasureTaskModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TreasureTask struct {
		Id               int64        `db:"id"`
		CreatedAt        sql.NullTime `db:"created_at"`
		UpdatedAt        sql.NullTime `db:"updated_at"`
		DeletedAt        sql.NullTime `db:"deleted_at"`
		Name             string       `db:"name"`
		DemandIntegral   int64        `db:"demand_integral"`
		TaskReward       int64        `db:"task_reward"`
		ExperienceReward int64        `db:"experience_reward"`
		RewardQuantity   int64        `db:"reward_quantity"`
		Employ           int64        `db:"employ"`
	}
)

func newTreasureTaskModel(conn sqlx.SqlConn) *defaultTreasureTaskModel {
	return &defaultTreasureTaskModel{
		conn:  conn,
		table: "`treasure_task`",
	}
}

func (m *defaultTreasureTaskModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTreasureTaskModel) FindOne(ctx context.Context, id int64) (*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", treasureTaskRows, m.table)
	var resp TreasureTask
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

func (m *defaultTreasureTaskModel) Insert(ctx context.Context, data *TreasureTask) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, treasureTaskRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.DemandIntegral, data.TaskReward, data.ExperienceReward, data.RewardQuantity, data.Employ)
	return ret, err
}

func (m *defaultTreasureTaskModel) Update(ctx context.Context, data *TreasureTask) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, treasureTaskRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.DemandIntegral, data.TaskReward, data.ExperienceReward, data.RewardQuantity, data.Id)
	return err
}

// UpdateOldStyle 下架旧样式
func (m *defaultTreasureTaskModel) UpdateOldStyle(ctx context.Context, data uint64) error {
	query := fmt.Sprintf("update %s set %s where `employ` = 1", m.table, "employ = 0")
	_, err := m.conn.ExecCtx(ctx, query)
	return err
}

// UpdateOldStyle 上架新样式
func (m *defaultTreasureTaskModel) UpdateNewStyle(ctx context.Context, data uint64) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, "employ = 1")
	_, err := m.conn.ExecCtx(ctx, query,data)
	return err
}

// 获取今日宝箱信息
func (m *defaultTreasureTaskModel) FindTreasureQuantity(ctx context.Context) (*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s where `employ` = 1 limit 1", treasureTaskRows, m.table)
	var resp TreasureTask
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 名称+奖励个数搜索
func (m *defaultTreasureTaskModel) FindNameAndQuantity(ctx context.Context,reward int64,name string,currPage, maxNum int64) ([]*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s where `reward_quantity` = ? AND `name` like ? limit %s offset %s order by id desc", treasureTaskRows, m.table,currPage,maxNum)
	var resp []*TreasureTask
	err := m.conn.QueryRowCtx(ctx, &resp, query,reward, "%"+name+"%")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// 名称+奖励个数搜索数量
func (m *defaultTreasureTaskModel) FindNameAndQuantityCount(ctx context.Context,reward int64,name string) (int64, error) {
	query := fmt.Sprintf("select COUNT(*) from %s where `reward_quantity` = ? AND `name` like ? ", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query,reward, "%"+name+"%")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// FindAll 查看全部
func (m *defaultTreasureTaskModel) FindAll(ctx context.Context,currPage, maxNum int64) ([]*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s limit %s offset %s order by id desc", treasureTaskRows, m.table,currPage,maxNum)
	var resp []*TreasureTask
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// FindAllCount 查看全部数量
func (m *defaultTreasureTaskModel) FindAllCount(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select COUNT(*) from %s", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// FindName 按名称查询
func (m *defaultTreasureTaskModel) FindName(ctx context.Context,name string,currPage, maxNum int64) ([]*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s where `name` like ? limit %s offset %s order by id desc", treasureTaskRows, m.table,currPage,maxNum)
	var resp []*TreasureTask
	err := m.conn.QueryRowCtx(ctx, &resp, query,"%"+name+"%")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// FindName 按名称查询数量
func (m *defaultTreasureTaskModel) FindNameCount(ctx context.Context,name string) (int64, error) {
	query := fmt.Sprintf("select COUNT(*) from %s where `name` like ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query,"%"+name+"%")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// FindQuantity 按奖励个数搜索
func (m *defaultTreasureTaskModel) FindQuantity(ctx context.Context,reward int64,currPage, maxNum int64) ([]*TreasureTask, error) {
	query := fmt.Sprintf("select %s from %s where `reward_quantity` = ? limit %s offset %s order by id desc", treasureTaskRows, m.table,currPage,maxNum)
	var resp []*TreasureTask
	err := m.conn.QueryRowCtx(ctx, &resp, query,reward)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// FindQuantityCount 按奖励个数搜索数量
func (m *defaultTreasureTaskModel) FindQuantityCount(ctx context.Context,reward int64) (int64, error) {
	query := fmt.Sprintf("select COUNT(*) from %s where `reward_quantity` = ?", m.table)
	var resp int64
	err := m.conn.QueryRowCtx(ctx, &resp, query,reward)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
func (m *defaultTreasureTaskModel) tableName() string {
	return m.table
}
