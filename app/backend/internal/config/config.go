/*
 * @Author: jobhandsome 1145938682@qq.com
 * @Date: 2023-07-10 23:18:02
 * @LastEditors: jobhandsome 1145938682@qq.com
 * @LastEditTime: 2023-07-11 09:41:33
 * @FilePath: /microSNS/app/backend/internal/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
