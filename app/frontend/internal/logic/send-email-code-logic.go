package logic

import (
	"context"
	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/common/helper"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendEmailCodeLogic {
	return SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req types.SenEmailCodeReq) (resp *types.CommonResply, err error) {

	// 验证邮箱
	if ok := helper.CheckUseremail(req.Email); !ok {
		return nil, Errorx.NewDefaultError("邮箱格式错误")
	}
	// 生成邮箱验证码
	emailCode := helper.RandCode()
	// 发送到邮箱
	sendErr := helper.MailSendCode(req.Email, emailCode)
	if sendErr != nil {
		return nil, Errorx.NewDefaultError(sendErr.Error())
	}
	// 存储邮箱验证码到redis
	setErr := l.svcCtx.RDB.Set(l.ctx, req.Email, emailCode, 30*time.Minute).Err()
	if setErr != nil {
		return nil, Errorx.NewDefaultError(setErr.Error())
	}

	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "发送成功",
	}, nil
}
