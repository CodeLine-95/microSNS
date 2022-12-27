package logic

import (
	"context"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"github.com/jobhandsome/microSNS/pkg/Paginate"
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"

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

	var SnsTopicM model.SnsTopics
	var SnsTopicUserM model.SnsTopicsUsers
	var SnsTopicData []types.TopicItem

	var TotalCount int64 = 0

	Where := "1=1 and " + SnsTopicM.TableName() + ".is_delete = 0"

	l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SnsTopicM.TableName() + ".*," + SnsTopicUserM.TableName() + ".cate_id," + SnsTopicUserM.TableName() + ".user_id").Joins("left join " + SnsTopicUserM.TableName() + " on " + SnsTopicM.TableName() + ".id = " + SnsTopicUserM.TableName() + ".topics_id").Count(&TotalCount)

	result := l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SnsTopicM.TableName() + ".*," + SnsTopicUserM.TableName() + ".cate_id," + SnsTopicUserM.TableName() + ".user_id").Joins("left join " + SnsTopicUserM.TableName() + " on " + SnsTopicM.TableName() + ".id = " + SnsTopicUserM.TableName() + ".topics_id").Where(Where).Scopes(Paginate.Paginate(req.Page, req.PageSize)).Order(SnsTopicM.TableName() + ".id desc").Find(&SnsTopicData)

	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	return &types.TopicListsResply{types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
	}, TotalCount, len(SnsTopicData), SnsTopicData}, nil
}
