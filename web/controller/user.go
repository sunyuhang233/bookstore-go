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

type UserController struct {
	UserService    *service.UserService
	CaptchaService *service.CaptchaService
}

func NewUserController() *UserController {
	return &UserController{
		UserService:    service.NewUserService(),
		CaptchaService: service.NewCaptchaService(),
	}
}

func (u *UserController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err,
		})
	}
	err := u.UserService.Register(req.Username, req.Password, req.Email, req.Phone, req.ConfirmPassword)
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

func (u *UserController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数错误",
			"error":   err,
		})
	}
	if !u.CaptchaService.VerifyCaptcha(req.CaptchaID, req.Image) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "验证码错误",
		})
		return
	}
	res, err := u.UserService.Login(req.Username, req.Password, req.Image, req.CaptchaID)
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

func (u *UserController) UserInfo(c *gin.Context) {
	userId, exists := c.Get("admin_user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	user, err := u.UserService.GetUserInfo(userId.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "获取用户信息失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取用户信息成功",
		"data":    user,
	})
}
