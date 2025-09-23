package router

import (
	"bookstore-go/web/controller"
	"bookstore-go/web/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
	})
	v1 := r.Group("/api/v1")
	userController := controller.NewUserController()
	captchaController := v1.Group("/captcha")
	{
		// 用户相关路由组
		user := v1.Group("/user")

		{
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
		}
	}

	controller := controller.NewCaptchaController()
	{
		captchaController.GET("/generate", controller.GenerateCaptcha)
	}

	auth := v1.Group("/auth")
	auth.Use(middleware.AdminAuthMiddleware())
	{
		auth.GET("/info", userController.UserInfo)
	}

	return r
}
