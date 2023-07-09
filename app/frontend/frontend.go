package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jobhandsome/microSNS/app/frontend/internal/config"
	"github.com/jobhandsome/microSNS/app/frontend/internal/handler"
	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/frontend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.WriteJson(w, http.StatusUnauthorized, Errorx.NewCodeError(http.StatusUnauthorized, "Authorization 不能为空"))
	}))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *Errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, e
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
