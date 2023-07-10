/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:58:58
 * @FilePath: /microSNS/app/frontend/internal/logic/user/user-integral-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"context"
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIntegralLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserIntegralLogic {
	return UserIntegralLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserIntegralLogic) UserIntegral(req types.UserIntegralReq) (resp *types.UserIntegralResply, err error) {
	UserId := 1

	var SnsUserIntegral model.SnsIntegralLogs
	var SnsUserIntegralData []types.UserIntegralItem
	var SnsUserIntegralCount int64

	page := req.Page
	pageSize := req.PageSize

	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	l.svcCtx.Engine.Table(SnsUserIntegral.TableName()).Where("user_id = ?", UserId).Count(&SnsUserIntegralCount)

	resultS := l.svcCtx.Engine.Table(SnsUserIntegral.TableName()).Where("user_id = ?", UserId).Order("id desc").Offset(offset).Limit(pageSize).Find(&SnsUserIntegralData)
	if resultS.Error != nil {
		return nil, Errorx.NewDefaultError(resultS.Error.Error())
	}

	return &types.UserIntegralResply{
		CommonResply: types.CommonResply{
			Code:    http.StatusOK,
			Message: "获取成功",
		},
		TotalCount: SnsUserIntegralCount,
		CurrCount:  len(SnsUserIntegralData),
		Data:       SnsUserIntegralData,
	}, nil
}
