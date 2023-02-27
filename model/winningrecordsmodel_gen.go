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
	winningRecordsFieldNames          = builder.RawFieldNames(&WinningRecords{})
	winningRecordsRows                = strings.Join(winningRecordsFieldNames, ",")
	winningRecordsRowsExpectAutoSet   = strings.Join(stringx.Remove(winningRecordsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	winningRecordsRowsWithPlaceHolder = strings.Join(stringx.Remove(winningRecordsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	winningRecordsModel interface {
		Insert(ctx context.Context, data *WinningRecords) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WinningRecords, error)
		Update(ctx context.Context, data *WinningRecords) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWinningRecordsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	WinningRecords struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		AwardId    int64     `db:"award_id"`
		ActivityId int64     `db:"activity_id"`
		Ip         string    `db:"ip"`
		Platform   string    `db:"platform"`
		CreatedAt  time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
	}
)

func newWinningRecordsModel(conn sqlx.SqlConn) *defaultWinningRecordsModel {
	return &defaultWinningRecordsModel{
		conn:  conn,
		table: "`winning_records`",
	}
}

func (m *defaultWinningRecordsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultWinningRecordsModel) FindOne(ctx context.Context, id int64) (*WinningRecords, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", winningRecordsRows, m.table)
	var resp WinningRecords
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

func (m *defaultWinningRecordsModel) Insert(ctx context.Context, data *WinningRecords) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, winningRecordsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.AwardId, data.ActivityId, data.Ip, data.Platform)
	return ret, err
}

func (m *defaultWinningRecordsModel) Update(ctx context.Context, data *WinningRecords) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, winningRecordsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.AwardId, data.ActivityId, data.Ip, data.Platform, data.Id)
	return err
}

func (m *defaultWinningRecordsModel) tableName() string {
	return m.table
}
