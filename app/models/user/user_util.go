package user

import "go-api/pkg/database"

// IsValExist 判断 Email/phone 已被注册
func IsValExist(val string) bool {
	var count int64
	switch val {
	case "email":
		database.DB.Model(User{}).Where("email =?", val).Count(&count)
	case "phone":
		database.DB.Model(User{}).Where("email =?", val).Count(&count)
	}

	return count > 0
}

// IsPhoneExist 判断手机号已被注册
