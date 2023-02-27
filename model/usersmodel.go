/*
 * @Author: licat
 * @Date: 2023-02-20 23:43:15
 * @LastEditors: licat
 * @LastEditTime: 2023-02-21 11:29:50
 * @Description: licat233@gmail.com
 */
package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersmodel
		usersModel
		FindsByActivityId(ctx context.Context, activityId int64) ([]*Users, error)
		FindsByPassport(ctx context.Context, passport string) ([]*Users, error)
		FindsByPassportAndActivityId(ctx context.Context, passport string, activityId int64) (*Users, error)
		FindsByIdAndActivityId(ctx context.Context, userId, activityId int64) (*Users, error)
		CountAdd(ctx context.Context, num, id int64) error
	}

	customUsersModel struct {
		*extendUsersModel
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn) UsersModel {
	m := newUsersModel(conn)
	return &customUsersModel{
		newExtendUsersModelModel(m),
		m,
	}
}

func (m *customUsersModel) FindsByActivityId(ctx context.Context, activityId int64) ([]*Users, error) {
	var resp []*Users
	query := fmt.Sprintf("select %s from %s where `activity_id` = ?", usersRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUsersModel) FindsByPassport(ctx context.Context, passport string) ([]*Users, error) {
	var resp []*Users
	query := fmt.Sprintf("select %s from %s where `passport` = ?", usersRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customUsersModel) FindsByPassportAndActivityId(ctx context.Context, passport string, activityId int64) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `passport` = ? and `activity_id` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, passport, activityId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) FindsByIdAndActivityId(ctx context.Context, userId, activityId int64) (*Users, error) {
	var resp Users
	query := fmt.Sprintf("select %s from %s where `id` = ? and `activity_id` = ? limit 1", usersRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, activityId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) CountAdd(ctx context.Context, num, id int64) error {
	query := fmt.Sprintf("update %s set count=count+? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, num, id)
	return err
}
