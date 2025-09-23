package router

import (
	"bookstore-go/web/controller"

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
	{
		// 用户相关路由组
		user := v1.Group("/user")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
		}
	}

	captcha := v1.Group("/captcha")
	{
		captcha.GET("/generate", controller.GenerateCaptcha)
	}

	return r
}
