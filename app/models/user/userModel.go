// Package user 存放用户 Model 相关逻辑
package user

import "go-api/app/models"

// User用户模型
type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty" gorm:"comment:用户名称;"`
	Email    string `json:"-" gorm:"comment:用户邮箱唯一;index:uniqueIndex;not null;"`
	Phone    string `json:"-" gorm:"comment:用户手机唯一;index:uniqueIndex;not null;"`
	Password string `json:"-" gorm:"not null;comment:用户密码;"`
	models.CommonTimestampsField
	models.SoftDeletes
}
