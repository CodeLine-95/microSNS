package section

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SectionListLogic {
	return SectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SectionListLogic) SectionList(req types.SectionListsReq) (resp *types.SectionListsResply, err error) {
	// todo: add your logic here and delete this line

	return
}
