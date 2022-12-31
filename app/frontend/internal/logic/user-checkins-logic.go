package logic

import (
	"context"
	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/common/helper"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserCheckinsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCheckinsLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserCheckinsLogic {
	return UserCheckinsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCheckinsLogic) UserCheckins(req types.UserCheckinsReq) (resp *types.CommonResply, err error) {

	UserId := 1

	var SnsCheckins model.SnsCheckins

	var SnsUsersM model.SnsUsers
	var SnsUserRow int64
	resultU := l.svcCtx.Engine.Table(SnsUsersM.TableName()).Where("id = ?", UserId).Count(&SnsUserRow)
	if resultU.Error != nil {
		return nil, Errorx.NewDefaultError(resultU.Error.Error())
	}
	if SnsUserRow == 0 {
		return nil, Errorx.NewDefaultError("该用户不存在，请检查用户......")
	}

	resultC := l.svcCtx.Engine.Table(SnsCheckins.TableName()).Where("user_id = ?", UserId).Limit(1).Order("id desc").Find(&SnsCheckins)
	if resultC.Error != nil {
		return nil, Errorx.NewDefaultError(resultC.Error.Error())
	}

	// 如果是第一次签到 | 断签
	SnsCheckinsC := model.SnsCheckins{
		UserId:         uint(UserId),
		CumulativeDays: SnsCheckins.CumulativeDays + 1,
		ContinuityDays: SnsCheckins.ContinuityDays + 1,
		LastTime:       l.svcCtx.T.String(),
		CreatedAt:      l.svcCtx.T.String(),
		UpdatedAt:      l.svcCtx.T.String(),
	}

	// 如果之前签到过
	if resultC.RowsAffected > 0 {
		LastTime := l.svcCtx.T.DateTimeToStamp(SnsCheckins.LastTime)
		LastDayTime := l.svcCtx.T.DateTimeToStamp(helper.SubStr(SnsCheckins.LastTime, 0, 10) + " 23:59:59")
		NowTime := l.svcCtx.T.DateTimeToStamp(l.svcCtx.T.String())
		BreakTime := l.svcCtx.T.DateTimeToStamp(l.svcCtx.T.SpecifiedDate(0, 0, 1) + " 23:59:59")

		if NowTime < LastDayTime {
			return nil, Errorx.NewDefaultError("今天已签到")
		}

		if LastTime < BreakTime {
			SnsCheckinsC.ContinuityDays = SnsCheckins.CumulativeDays + 1
		}
	}

	resultS := l.svcCtx.Engine.Create(&SnsCheckinsC)
	if resultS.Error != nil {
		return nil, Errorx.NewDefaultError(resultS.Error.Error())
	}

	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "签到成功",
	}, nil
}
