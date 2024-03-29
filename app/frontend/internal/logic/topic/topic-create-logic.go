/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:59:53
 * @FilePath: /microSNS/app/frontend/internal/logic/topic/topic-create-logic.go
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
	SnsTopic := model.SnsTopics{
		Title:     req.Title,
		Tags:      req.Tags,
		State:     0,
		Type:      0,
		Content:   req.Content,
		MdContent: req.MdContent,
		IsDelete:  0,
		CreatedAt: l.svcCtx.T.String(),
		UpdatedAt: l.svcCtx.T.String(),
	}

	transaction := l.svcCtx.Engine.Begin()

	if transaction.Error != nil {
		transaction.Rollback()
		return nil, Errorx.NewDefaultError(transaction.Error.Error())
	}

	if SnsTopicErr := transaction.Create(&SnsTopic).Error; SnsTopicErr != nil {
		transaction.Rollback()
		return nil, Errorx.NewDefaultError(SnsTopicErr.Error())
	}

	SnsTopicsUsers := model.SnsTopicsUsers{
		TopicsId:  SnsTopic.Id,
		CateId:    uint(req.CateId),
		UserId:    0,
		ReplyId:   0,
		CreatedAt: l.svcCtx.T.String(),
		UpdatedAt: l.svcCtx.T.String(),
	}

	if SnsTopicsUsersErr := transaction.Create(&SnsTopicsUsers).Error; SnsTopicsUsersErr != nil {
		transaction.Rollback()
		return nil, Errorx.NewDefaultError(SnsTopicsUsersErr.Error())
	}

	if CommitErr := transaction.Commit().Error; CommitErr != nil {
		transaction.Rollback()
		return nil, Errorx.NewDefaultError(CommitErr.Error())
	}

	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "发布成功",
	}, nil
}
