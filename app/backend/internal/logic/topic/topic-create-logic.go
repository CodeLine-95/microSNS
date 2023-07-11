package topic

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) TopicCreateLogic {
	return TopicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicCreateLogic) TopicCreate(req types.TopicCreateReq) (resp *types.CommonResply, err error) {
	// todo: add your logic here and delete this line

	return
}
