/*
 * @Author: jobhandsome 1145938682@qq.com
 * @Date: 2023-07-10 23:18:02
 * @LastEditors: jobhandsome 1145938682@qq.com
 * @LastEditTime: 2023-07-11 09:44:07
 * @FilePath: /microSNS/app/backend/internal/logic/common/login-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"context"
	"net/http"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Crypto"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/jobhandsome/microSNS/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResply, err error) {
	var SnsAccountM model.SnsAccount

	result := l.svcCtx.Engine.Table(SnsAccountM.TableName()).Where("name = ?", req.UserName).Find(&SnsAccountM)
	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if result.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("该账户不存在...")
	}

	AesDePass := Crypto.AesDecrypt(SnsAccountM.Pass, l.svcCtx.Config.SecretKey)

	if AesDePass != req.PassWord {
		return nil, Errorx.NewDefaultError("密码错误，请重试...")
	}
	// 生成token
	mapClaims := make(map[string]interface{})

	mapClaims["userId"] = SnsAccountM.Id
	mapClaims["name"] = SnsAccountM.Name

	token, tokenErr := jwt.GenerateToken(mapClaims, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if tokenErr != nil {
		return nil, Errorx.NewDefaultError(tokenErr.Error())
	}

	return &types.LoginResply{
		CommonResply: types.CommonResply{
			Code:    http.StatusOK,
			Message: "登录成功",
		},
		Token: token,
	}, nil
}
