package svc

import (
	"fmt"
	"github.com/jobhandsome/microSNS/app/frontend/internal/config"
	"github.com/jobhandsome/microSNS/common/define"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/LocalTime"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/url"
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
