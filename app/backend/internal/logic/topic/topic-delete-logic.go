package topic

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopicDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) TopicDeleteLogic {
	return TopicDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicDeleteLogic) TopicDelete(req types.TopicItemReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
