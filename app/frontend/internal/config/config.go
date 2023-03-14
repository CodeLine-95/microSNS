package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	SecretKey string
	Mysql     struct {
		Host    string
		User    string
		Pass    string
		Data    string
		Charset string
		Debug   bool
	}
	Redis struct {
		Addr string
		Pass string
	}

	Auth struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
		AccessExpire int64  // 有效期，单位：秒
	}
}
