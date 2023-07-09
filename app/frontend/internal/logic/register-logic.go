package logic

import (
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"

	"github.com/jobhandsome/microSNS/common/helper"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Crypto"
	"github.com/jobhandsome/microSNS/pkg/Errorx"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.CommonResply, err error) {

	var SnsUserM model.SnsUsers

	var RowsAffected int64 = 0

	if ok := helper.CheckUseremail(req.Email); !ok {
		return nil, Errorx.NewDefaultError("邮箱格式错误，请更改后重试...")
	}

	// 验证密码格式：最小6位，最大16位
	if ok := helper.CheckPassword(req.Pass); !ok {
		return nil, Errorx.NewDefaultError("密码格式错误，请更改后重试...")
	}

	result := l.svcCtx.Engine.Table(SnsUserM.TableName()).Where("email = ?", req.Email).Count(&RowsAffected)
	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if RowsAffected > 0 {
		return nil, Errorx.NewDefaultError("该邮箱已注册")
	}

	// 确认密码验证
	if req.Pass != req.ConfirmPass {
		return nil, Errorx.NewDefaultError("确认密码错误")
	}

	// 获取发送的验证码
	sendCode, sendErr := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if sendErr != nil && sendErr != redis.Nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}
	if len(sendCode) == 0 {
		return nil, Errorx.NewDefaultError("该验证码已过期")
	}

	// 邮箱验证码验证
	if sendCode != req.Code {
		return nil, Errorx.NewDefaultError("验证码错误")
	}

	AesEnPass := Crypto.AesEncrypt(req.Pass, l.svcCtx.Config.SecretKey)

	saveRes := l.svcCtx.Engine.Create(&model.SnsUsers{
		Name:      "sns_" + helper.RandomString(10),
		Pass:      AesEnPass,
		Email:     req.Email,
		State:     0,
		IsDelete:  0,
		CreatedAt: l.svcCtx.T.String(),
		UpdatedAt: l.svcCtx.T.String(),
	})

	if saveRes.Error != nil || saveRes.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("注册失败")
	}

	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "注册成功",
	}, nil
}
