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
	activityFieldNames          = builder.RawFieldNames(&Activity{})
	activityRows                = strings.Join(activityFieldNames, ",")
	activityRowsExpectAutoSet   = strings.Join(stringx.Remove(activityFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	activityRowsWithPlaceHolder = strings.Join(stringx.Remove(activityFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	activityModel interface {
		Insert(ctx context.Context, data *Activity) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Activity, error)
		Update(ctx context.Context, data *Activity) error
		Delete(ctx context.Context, id int64) error
	}

	defaultActivityModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Activity struct {
		Id        int64     `db:"id"`
		Uuid      string    `db:"uuid"`
		Name      string    `db:"name"`
		Status    int64     `db:"status"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func newActivityModel(conn sqlx.SqlConn) *defaultActivityModel {
	return &defaultActivityModel{
		conn:  conn,
		table: "`activity`",
	}
}

func (m *defaultActivityModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultActivityModel) FindOne(ctx context.Context, id int64) (*Activity, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", activityRows, m.table)
	var resp Activity
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

func (m *defaultActivityModel) Insert(ctx context.Context, data *Activity) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, activityRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status)
	return ret, err
}

func (m *defaultActivityModel) Update(ctx context.Context, data *Activity) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, activityRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Uuid, data.Name, data.Status, data.Id)
	return err
}

func (m *defaultActivityModel) tableName() string {
	return m.table
}
