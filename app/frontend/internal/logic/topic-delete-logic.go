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

	var SnsTopic model.SnsTopics

	result := l.svcCtx.Engine.Table(SnsTopic.TableName()).Where("id = ?", req.TopicId).Updates(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": l.svcCtx.T.String(),
	})

	if result.Error != nil {
		return nil, Errorx.NewDefaultError(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, Errorx.NewDefaultError("删除失败")
	}
	return &types.CommonResply{
		Code:    http.StatusOK,
		Message: "删除成功",
	}, nil
}
