package logic

import (
	"context"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
