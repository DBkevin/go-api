// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "go-api/app/http/controllers/api/v1"
	"go-api/app/models/user"
	"go-api/app/requests"
	"go-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsValExist("phone", request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsValExist("email", request.Email),
	})
}
