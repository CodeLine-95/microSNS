package logic

import (
	"context"
	"github.com/jobhandsome/microSNS/model"
	"github.com/jobhandsome/microSNS/pkg/Errorx"
	"net/http"

	"github.com/jobhandsome/microSNS/app/frontend/internal/svc"
	"github.com/jobhandsome/microSNS/app/frontend/internal/types"

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

	var SnsTopicM model.SnsTopics
	var SnsTopicUserM model.SnsTopicsUsers
	var SnsTopicData types.TopicItem

	result := l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SnsTopicM.TableName()+".*,"+SnsTopicUserM.TableName()+".cate_id,"+SnsTopicUserM.TableName()+".user_id").Joins("left join "+SnsTopicUserM.TableName()+" on "+SnsTopicM.TableName()+".id = "+SnsTopicUserM.TableName()+".topics_id").Where(SnsTopicM.TableName()+".id = ?", req.TopicId).Find(&SnsTopicData)

	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if result.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("该话题不存在或已删除")
	}

	return &types.TopicDetailResply{types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
	}, SnsTopicData}, nil
}
