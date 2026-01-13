package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	// DB 避免业务代码直接引用gorm.
	DB = gorm.DB

	Locking    = clause.Locking
	OnConflict = clause.OnConflict

	Table      = clause.Table
	Column     = clause.Column
	Where      = clause.Where
	Expression = clause.Expression

	Set        = clause.Set
	Assignment = clause.Assignment
)

var (
	// Expr Update 条件表达式 避免业务代码直接引用gorm.
	Expr = gorm.Expr

	// ErrRecordNotFound mysql查找结果为空的错误字符串.
	ErrRecordNotFound = gorm.ErrRecordNotFound
	// ErrInvalidTransaction occurs when you are trying to `Commit` or `Rollback`.
	ErrInvalidTransaction = gorm.ErrInvalidTransaction
)

func Assignments(values map[string]any) clause.Set {
	return clause.Assignments(values)
}

func AssignmentColumns(values []string) clause.Set {
	return clause.AssignmentColumns(values)
}
