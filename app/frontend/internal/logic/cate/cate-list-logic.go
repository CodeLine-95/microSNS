/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 20:00:07
 * @FilePath: /microSNS/app/frontend/internal/logic/cate/cate-list-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cate

import (
	"context"
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/jobhandsome/microSNS/pkg/Paginate"

	"github.com/zeromicro/go-zero/core/logx"
)

type CateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CateListLogic {
	return CateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CateListLogic) CateList(req types.CateListsReq) (resp *types.CateListsResply, err error) {
	var SnsCateM model.SnsCate

	var SnsCateData []types.CateItem

	var TotalCount int64 = 0

	Where := "1=1 and is_delete = 0"

	l.svcCtx.Engine.Table(SnsCateM.TableName()).Count(&TotalCount)

	result := l.svcCtx.Engine.Table(SnsCateM.TableName()).Where(Where).Scopes(Paginate.Paginate(req.Page, req.PageSize)).Order("id desc").Find(&SnsCateData)

	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	return &types.CateListsResply{CommonResply: types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
	}, TotalCount: TotalCount, CurrCount: len(SnsCateData), Data: SnsCateData}, nil
}
