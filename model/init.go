/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/24 22:02
 */

package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql(DataSource string) *gorm.DB {
	engine, err := gorm.Open(mysql.Open(DataSource))
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

func InitRedis(Addr, Pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Pass, // no password set
		DB:       0,    // use default DB
	})
}
