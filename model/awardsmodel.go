/*
 * @Author: licat
 * @Date: 2023-02-20 23:43:15
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 11:22:27
 * @Description: licat233@gmail.com
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AwardsModel = (*customAwardsModel)(nil)

type (
	// AwardsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAwardsModel.
	AwardsModel interface {
		awardsModel
		awardsmodel
		FindsByActivityId(ctx context.Context, activityId int64) ([]*Awards, error)
		CountAdd(ctx context.Context, num, id int64) error
	}

	customAwardsModel struct {
		*defaultAwardsModel
		*extendAwardsModel
	}
)

// NewAwardsModel returns a model for the database table.
func NewAwardsModel(conn sqlx.SqlConn) AwardsModel {
	m := newAwardsModel(conn)
	return &customAwardsModel{
		m,
		newExtendAwardsModelModel(m),
	}
}

func (m *customAwardsModel) FindsByActivityId(ctx context.Context, activityId int64) ([]*Awards, error) {
	var resp []*Awards
	query := fmt.Sprintf("select %s from %s where `activity_id` = ? ", awardsRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, activityId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customAwardsModel) CountAdd(ctx context.Context, num, id int64) error {
	query := fmt.Sprintf("update %s set count=count+? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, num, id)
	return err
}
