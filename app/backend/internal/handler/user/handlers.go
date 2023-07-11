package handler

import (
	"net/http"

	"github.com/jobhandsome/microSNS/app/backend/internal/logic/user"
	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func EditPassHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditPassReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditPassLogic(r.Context(), ctx)
		resp, err := l.EditPass(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
