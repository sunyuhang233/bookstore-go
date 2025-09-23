package controller

import (
	"bookstore-go/service"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	CaptchaService *service.CaptchaService
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{
		CaptchaService: service.NewCaptchaService(),
	}
}

func (captcha *CaptchaController) GenerateCaptcha(c *gin.Context) {
	res, err := captcha.CaptchaService.GenerateCaptcha()
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
