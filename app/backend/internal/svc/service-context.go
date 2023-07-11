/*
 * @Author: jobhandsome 1145938682@qq.com
 * @Date: 2023-07-10 23:18:02
 * @LastEditors: jobhandsome 1145938682@qq.com
 * @LastEditTime: 2023-07-11 09:43:33
 * @FilePath: /microSNS/app/backend/internal/svc/service-context.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package svc

import (
	"fmt"
	"net/url"

	"github.com/jobhandsome/microSNS/app/backend/internal/config"
	"github.com/jobhandsome/microSNS/common/define"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/LocalTime"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB
	RDB    *redis.Client
	T      *LocalTime.LocalTime
}

func NewServiceContext(c config.Config) *ServiceContext {
	DataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=%s",
		c.Mysql.User, c.Mysql.Pass, c.Mysql.Host, c.Mysql.Data, c.Mysql.Charset, url.QueryEscape(define.TimeZone))

	return &ServiceContext{
		Config: c,
		Engine: model.InitMysql(DataSource, c.Mysql.Debug),
		RDB:    model.InitRedis(c.Redis.Addr, c.Redis.Pass),
		T:      &LocalTime.LocalTime{},
	}
}
