// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package role

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// 角色表
type Role struct {
	ID int32
	// 角色名称
	Rolename string
	// 状态
	Status int32
	// 创建时间
	Created pgtype.Timestamp
}