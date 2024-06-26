// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	User       *user
	UserApply  *userApply
	UserFriend *userFriend
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	User = &Q.User
	UserApply = &Q.UserApply
	UserFriend = &Q.UserFriend
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		User:       newUser(db, opts...),
		UserApply:  newUserApply(db, opts...),
		UserFriend: newUserFriend(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	User       user
	UserApply  userApply
	UserFriend userFriend
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		User:       q.User.clone(db),
		UserApply:  q.UserApply.clone(db),
		UserFriend: q.UserFriend.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		User:       q.User.replaceDB(db),
		UserApply:  q.UserApply.replaceDB(db),
		UserFriend: q.UserFriend.replaceDB(db),
	}
}

type queryCtx struct {
	User       IUserDo
	UserApply  IUserApplyDo
	UserFriend IUserFriendDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		User:       q.User.WithContext(ctx),
		UserApply:  q.UserApply.WithContext(ctx),
		UserFriend: q.UserFriend.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
