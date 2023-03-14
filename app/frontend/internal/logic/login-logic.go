package logic

import (
	"context"
	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Crypto"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/jobhandsome/microSNS/pkg/jwt"
	"net/http"

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
