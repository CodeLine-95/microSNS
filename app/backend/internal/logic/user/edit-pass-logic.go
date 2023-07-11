package user

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPassLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditPassLogic {
	return EditPassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditPassLogic) EditPass(req types.EditPassReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
