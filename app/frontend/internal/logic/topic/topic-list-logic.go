/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:59:22
 * @FilePath: /microSNS/app/frontend/internal/logic/topic/topic-list-logic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package topic

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

	SelectFields := SnsTopicM.JoinSelectFields("cate_id,user_id,reply_id,view_count,like_count,comment_count")

	l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SelectFields).Joins("left join " + SnsTopicUserM.TableName() + " on " + SnsTopicM.TableName() + ".id = " + SnsTopicUserM.TableName() + ".topics_id").Count(&TotalCount)

	result := l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SelectFields).Joins("left join " + SnsTopicUserM.TableName() + " on " + SnsTopicM.TableName() + ".id = " + SnsTopicUserM.TableName() + ".topics_id").Where(Where).Scopes(Paginate.Paginate(req.Page, req.PageSize)).Order(SnsTopicM.TableName() + ".id desc").Find(&SnsTopicData)

	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	return &types.TopicListsResply{CommonResply: types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
	}, TotalCount: TotalCount, CurrCount: len(SnsTopicData), Data: SnsTopicData}, nil
}
