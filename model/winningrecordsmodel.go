/*
 * @Author: licat
 * @Date: 2023-02-20 23:43:15
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 12:59:23
 * @Description: licat233@gmail.com
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WinningRecordsModel = (*customWinningRecordsModel)(nil)

type (
	// WinningRecordsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWinningRecordsModel.
	WinningRecordsModel interface {
		winningRecordsmodel
		winningRecordsModel
		FindsByUserIdAndActivityId(ctx context.Context, userId, activityId int64) ([]*WinningRecords, error)
	}

	customWinningRecordsModel struct {
		*extendWinningRecordsModel
		*defaultWinningRecordsModel
	}
)

// NewWinningRecordsModel returns a model for the database table.
func NewWinningRecordsModel(conn sqlx.SqlConn) WinningRecordsModel {
	m := newWinningRecordsModel(conn)
	return &customWinningRecordsModel{
		newExtendWinningRecordsModelModel(m),
		m,
	}
}

func (m *defaultWinningRecordsModel) FindsByUserIdAndActivityId(ctx context.Context, userId, activityId int64) ([]*WinningRecords, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `activity_id` = ?", winningRecordsRows, m.table)
	var resp []*WinningRecords
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId, activityId)
	return resp, err
}
