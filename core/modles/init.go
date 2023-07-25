package modles

import (
	"github.com/go-redis/redis/v8"
	"go_disk/core/internal/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Init(c config.Config) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", c.Mysql.DataSource)
	if err != nil {
		log.Printf("xorm new Engine failed...,error:%v", err)
		return nil
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
