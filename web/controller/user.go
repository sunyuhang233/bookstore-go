package controller

import (
	"bookstore-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}

type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CaptchaID string `json:"captcha_id"`
	Image     string `json:"image"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err,
		})
	}
	service := service.NewUserService()
	err := service.Register(req.Username, req.Password, req.Email, req.Phone, req.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "注册失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
	})

}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err,
		})
	}
	user_service := service.NewUserService()
	captcha_service := service.NewCaptchaService()
	if !captcha_service.VerifyCaptcha(req.CaptchaID, req.Image) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}
	res, err := user_service.Login(req.Username, req.Password, req.Image, req.CaptchaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "登录失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    res,
	})
}
