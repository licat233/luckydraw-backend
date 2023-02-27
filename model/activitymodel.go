/*
 * @Author: licat
 * @Date: 2023-02-20 23:43:15
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 11:06:14
 * @Description: licat233@gmail.com
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ActivityModel = (*customActivityModel)(nil)

type (
	// ActivityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityModel.
	ActivityModel interface {
		activitymodel
		activityModel
		FindOneByUuid(ctx context.Context, uuid string) (*Activity, error)
		FindAll(ctx context.Context) ([]*Activity, error)
	}

	customActivityModel struct {
		*extendActivityModel
		*defaultActivityModel
	}
)

// NewActivityModel returns a model for the database table.
func NewActivityModel(conn sqlx.SqlConn) ActivityModel {
	m := newActivityModel(conn)
	return &customActivityModel{
		newExtendActivityModelModel(m),
		m,
	}
}

func (m *customActivityModel) FindOneByUuid(ctx context.Context, uuid string) (*Activity, error) {
	query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", activityRows, m.table)
	var resp Activity
	err := m.conn.QueryRowCtx(ctx, &resp, query, uuid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customActivityModel) FindAll(ctx context.Context) ([]*Activity, error) {
	query := fmt.Sprintf("select %s from %s", activityRows, m.table)
	var resp []*Activity
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	return resp, err
}
