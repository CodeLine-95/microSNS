/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:56:40
 * @FilePath: /microSNS/app/frontend/internal/logic/common/login-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"context"
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
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
	var SnsUserM model.SnsUsers

	result := l.svcCtx.Engine.Table(SnsUserM.TableName()).Where("email = ?", req.Email).Find(&SnsUserM)
	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if result.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("该邮箱未注册")
	}

	AesDePass := Crypto.AesDecrypt(SnsUserM.Pass, l.svcCtx.Config.SecretKey)

	if AesDePass != req.Pass {
		return nil, Errorx.NewDefaultError("密码错误，请重试...")
	}
	// 生成token
	mapClaims := make(map[string]interface{})

	mapClaims["userId"] = SnsUserM.Id
	mapClaims["name"] = SnsUserM.Name
	mapClaims["email"] = SnsUserM.Email

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
