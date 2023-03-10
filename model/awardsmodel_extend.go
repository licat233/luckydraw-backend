// Code generated by sql2rpc. DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

type (
	awardsmodel interface {
		FindList(ctx context.Context, pageSize, page int64, keyword string, awards *Awards) (resp []*Awards, total int64, err error)
	}
	extendAwardsModel struct {
		*defaultAwardsModel
	}
)

func newExtendAwardsModelModel(defaultAwardsModel *defaultAwardsModel) *extendAwardsModel {
	return &extendAwardsModel{
		defaultAwardsModel,
	}
}

func (m *extendAwardsModel) FindList(ctx context.Context, pageSize, page int64, keyword string, awards *Awards) (resp []*Awards, total int64, err error) {
	hasName := false
	sq := squirrel.Select(awardsRows).From(m.table)
	if awards != nil {
		if awards.Id > 0 {
			sq = sq.Where("id = ?", awards.Id)
		}
		if awards.ActivityId > 0 {
			sq = sq.Where("activity_id = ?", awards.ActivityId)
		}
		if awards.Uuid != "" {
			sq = sq.Where("uuid = ?", awards.Uuid)
		}
		if awards.Grade != "" {
			sq = sq.Where("grade = ?", awards.Grade)
		}
		if awards.Name != "" {
			sq = sq.Where("name = ?", awards.Name)
			hasName = true
		}
		if awards.Image != "" {
			sq = sq.Where("image = ?", awards.Image)
		}
		if awards.Prob >= 0 {
			sq = sq.Where("prob = ?", awards.Prob)
		}
		if awards.Quantity >= 0 {
			sq = sq.Where("quantity = ?", awards.Quantity)
		}
		if awards.Count >= 0 {
			sq = sq.Where("count = ?", awards.Count)
		}
		if awards.IsWin >= 0 {
			sq = sq.Where("is_win = ?", awards.IsWin)
		}
	}
	if keyword != "" && !hasName {
		sq = sq.Where("name LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	if pageSize > 0 && page > 0 {
		sqCount := sq.RemoveLimit().RemoveOffset()
		sq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))
		queryCount, agrsCount, e := sqCount.ToSql()
		if e != nil {
			err = e
			return
		}
		queryCount = strings.ReplaceAll(queryCount, awardsRows, "COUNT(*)")
		if err = m.conn.QueryRowCtx(ctx, &total, queryCount, agrsCount...); err != nil {
			return
		}
	}
	query, agrs, err := sq.ToSql()
	if err != nil {
		return
	}
	resp = make([]*Awards, 0)
	if err = m.conn.QueryRowsCtx(ctx, &resp, query, agrs...); err != nil {
		return
	}
	return
}
