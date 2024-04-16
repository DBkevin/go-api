// Package models 模型通用属性和方法
package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;comment:主键 uint64 自增;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatdAt  time.Time `gorm:"column:created_at;index;comment:创建时间戳,索引;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;comment:更新时间戳，索引;" json:"updated_at,omitempty"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index;comment:软删除，删除时间"`
}
