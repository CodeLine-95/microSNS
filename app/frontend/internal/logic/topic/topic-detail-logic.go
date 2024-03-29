/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:59:32
 * @FilePath: /microSNS/app/frontend/internal/logic/topic/topic-detail-logic.go
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

	SelectFields := SnsTopicM.JoinSelectFields("cate_id,user_id,reply_id,view_count,like_count,comment_count")

	result := l.svcCtx.Engine.Table(SnsTopicM.TableName()).Select(SelectFields).Joins("left join "+SnsTopicUserM.TableName()+" on "+SnsTopicM.TableName()+".id = "+SnsTopicUserM.TableName()+".topics_id").Where(SnsTopicM.TableName()+".id = ?", req.TopicId).Find(&SnsTopicData)

	if result.Error != nil {
		return nil, Errorx.NewDefaultError("系统异常")
	}

	if result.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("该话题不存在或已删除")
	}

	return &types.TopicDetailResply{CommonResply: types.CommonResply{
		Code:    http.StatusOK,
		Message: "获取成功",
	}, Data: SnsTopicData}, nil
}
