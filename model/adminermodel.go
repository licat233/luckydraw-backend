/*
 * @Author: licat
 * @Date: 2023-02-22 14:23:59
 * @LastEditors: licat
 * @LastEditTime: 2023-02-22 14:35:40
 * @Description: licat233@gmail.com
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminerModel = (*customAdminerModel)(nil)

type (
	// AdminerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminerModel.
	AdminerModel interface {
		adminermodel
		adminerModel
		FindByUsername(ctx context.Context, username string) (*Adminer, error)
	}

	customAdminerModel struct {
		*extendAdminerModel
		*defaultAdminerModel
	}
)

// NewAdminerModel returns a model for the database table.
func NewAdminerModel(conn sqlx.SqlConn) AdminerModel {
	m := newAdminerModel(conn)
	return &customAdminerModel{
		newExtendAdminerModelModel(m),
		m,
	}
}

func (m *customAdminerModel) FindByUsername(ctx context.Context, username string) (*Adminer, error) {
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", adminerRows, m.table)
	var resp Adminer
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
