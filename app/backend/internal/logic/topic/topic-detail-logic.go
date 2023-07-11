package topic

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) TopicDetailLogic {
	return TopicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicDetailLogic) TopicDetail(req types.TopicItemReq) (resp *types.TopicDetailResply, err error) {
	// todo: add your logic here and delete this line

	return
}
