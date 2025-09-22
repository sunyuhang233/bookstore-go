package global

import (
	"bookstore-go/config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbClient *gorm.DB
var RedisClient *redis.Client

func InitMysql() {
	mysqlConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}
	DbClient = client
	log.Println("数据库连接成功")
}

func InitRedis() {
	redisConfig := config.AppConfig.Redis
	client := redis.NewClient(&redis.Options{
		DB:       redisConfig.DB,
		Password: redisConfig.Password,
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
	})
	str, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatalln("连接Redis失败:", err)
	}
	log.Println("Redis返回:", str)
	RedisClient = client
	log.Println("Redis连接成功")
}
