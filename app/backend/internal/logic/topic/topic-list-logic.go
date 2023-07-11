package topic

import (
	"context"

	"github.com/jobhandsome/microSNS/app/backend/internal/svc"
	"github.com/jobhandsome/microSNS/app/backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) TopicListLogic {
	return TopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TopicListLogic) TopicList(req types.TopicListsReq) (resp *types.TopicListsResply, err error) {
	// todo: add your logic here and delete this line

	return
}
