// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	chatFieldNames          = builder.RawFieldNames(&Chat{})
	chatRows                = strings.Join(chatFieldNames, ",")
	chatRowsExpectAutoSet   = strings.Join(stringx.Remove(chatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatRowsWithPlaceHolder = strings.Join(stringx.Remove(chatFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheChatIdPrefix = "cache:chat:id:"
)

type (
	chatModel interface {
		Insert(ctx context.Context, data *Chat) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Chat, error)
		Update(ctx context.Context, data *Chat) error
		Delete(ctx context.Context, id int64) error
	}

	defaultChatModel struct {
		sqlc.CachedConn
		table string
	}

	Chat struct {
		Id         int64     `db:"id"`
		User       string    `db:"user"`        // 用户标识
		AgentId    int64     `db:"agent_id"`    // 应用ID
		ReqContent string    `db:"req_content"` // 用户发送内容
		ResContent string    `db:"res_content"` // openai响应内容
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		UpdatedAt  time.Time `db:"updated_at"`  // 更新时间
	}
)

func newChatModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultChatModel {
	return &defaultChatModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`chat`",
	}
}

func (m *defaultChatModel) Delete(ctx context.Context, id int64) error {
	chatIdKey := fmt.Sprintf("%s%v", cacheChatIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, chatIdKey)
	return err
}

func (m *defaultChatModel) FindOne(ctx context.Context, id int64) (*Chat, error) {
	chatIdKey := fmt.Sprintf("%s%v", cacheChatIdPrefix, id)
	var resp Chat
	err := m.QueryRowCtx(ctx, &resp, chatIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatModel) Insert(ctx context.Context, data *Chat) (sql.Result, error) {
	chatIdKey := fmt.Sprintf("%s%v", cacheChatIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, chatRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.User, data.AgentId, data.ReqContent, data.ResContent)
	}, chatIdKey)
	return ret, err
}

func (m *defaultChatModel) Update(ctx context.Context, data *Chat) error {
	chatIdKey := fmt.Sprintf("%s%v", cacheChatIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.User, data.AgentId, data.ReqContent, data.ResContent, data.Id)
	}, chatIdKey)
	return err
}

func (m *defaultChatModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheChatIdPrefix, primary)
}

func (m *defaultChatModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultChatModel) tableName() string {
	return m.table
}
