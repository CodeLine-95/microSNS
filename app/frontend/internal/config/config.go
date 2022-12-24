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
	}
	Redis struct {
		Addr string
		Pass string
	}

	Auth struct {
		Secret string
		Expire int
	}
}
