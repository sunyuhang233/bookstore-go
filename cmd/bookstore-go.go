package main

import (
	"bookstore-go/config"
	"bookstore-go/global"
)

func main() {
	config.InitConfig("../conf/config.yaml")
	global.InitMysql()
	global.InitRedis()
}
