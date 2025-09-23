package controller

import (
	"bookstore-go/service"

	"github.com/gin-gonic/gin"
)

func GenerateCaptcha(c *gin.Context) {
	service := service.NewCaptchaService()
	res, err := service.GenerateCaptcha()
	if err != nil {
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "验证码生成失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "验证码生成成功",
		"data":    res,
	})
}
