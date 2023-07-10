/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:57:20
 * @FilePath: /microSNS/app/frontend/internal/logic/common/send-email-code-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"context"
	"net/http"
	"time"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/common/helper"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"

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
	var SnsUserM model.SnsUsers
	// 验证邮箱
	if ok := helper.CheckUseremail(req.Email); !ok {
		return nil, Errorx.NewDefaultError("邮箱格式错误")
	}
	result := l.svcCtx.Engine.Table(SnsUserM.TableName()).Where("email = ?", req.Email).Find(&SnsUserM)
	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if result.RowsAffected > 0 {
		return nil, Errorx.NewDefaultError("该邮箱已注册")
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
