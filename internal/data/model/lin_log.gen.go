// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLinLog = "lin_log"

// LinLog mapped from table <lin_log>
type LinLog struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Message    string         `gorm:"column:message" json:"message"`
	UserID     int32          `gorm:"column:user_id;not null" json:"user_id"`
	Username   string         `gorm:"column:username" json:"username"`
	StatusCode int64          `gorm:"column:status_code" json:"status_code"`
	Method     string         `gorm:"column:method" json:"method"`
	Path       string         `gorm:"column:path" json:"path"`
	Permission string         `gorm:"column:permission" json:"permission"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName LinLog's table name
func (*LinLog) TableName() string {
	return TableNameLinLog
}
