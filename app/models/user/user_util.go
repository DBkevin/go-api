package user

import (
	"fmt"
	"go-api/pkg/database"
)

// IsValExist 判断 Email/phone 已被注册
func IsValExist(valnanme string, val string) bool {
	var count int64
	switch valnanme {
	case "email":
		fmt.Println("走email")
		database.DB.Model(User{}).Where("email =?", val).Count(&count)
	case "phone":
		fmt.Println("走phone")
		database.DB.Model(User{}).Where("phone =?", val).Count(&count)
	}

	return count > 0
}

// IsPhoneExist 判断手机号已被注册
