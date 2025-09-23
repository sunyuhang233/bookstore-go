package main

import (
	"bookstore-go/config"
	"bookstore-go/global"
	"bookstore-go/web/router"
	"fmt"
	"net/http"
	"os"
)

func main() {
	config.InitConfig("../conf/config.yaml")
	global.InitMysql()
	global.InitRedis()
	r:=router.InitRouter()
	addr:=fmt.Sprintf("%s:%d","localhost",config.AppConfig.Server.Port)
	sever:= &http.Server{
		Addr: addr,
		Handler: r,
	}
	err:=sever.ListenAndServe()
	if err!=nil{
		fmt.Println("服务启动失败,err:",err)
		os.Exit(-1)
	}
	fmt.Println("服务启动成功,监听地址:",addr)
}
