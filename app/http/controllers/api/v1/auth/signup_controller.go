// Package auth 处理用户身份认证相关逻辑
package auth

import (
	"fmt"
	v1 "go-api/app/http/controllers/api/v1"
	"go-api/app/models/user"
	"go-api/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	errs := requests.ValidateSignupPhoneExist(&request, c)
	//errs返回长度等于零既通过，大于0既有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	//检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsValExist("phone", request.Phone),
	})
}
