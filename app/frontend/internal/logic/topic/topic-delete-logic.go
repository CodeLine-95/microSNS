/*
 * @Author: qiaoshuai p-qiaoshuai@xiaomi.com
 * @Date: 2023-07-10 19:54:20
 * @LastEditors: qiaoshuai p-qiaoshuai@xiaomi.com
 * @LastEditTime: 2023-07-10 19:59:42
 * @FilePath: /microSNS/app/frontend/internal/logic/topic/topic-delete-logic.go
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
	var SnsTopicM model.SnsTopics

	result := l.svcCtx.Engine.Table(SnsTopicM.TableName()).Where("id = ?", req.TopicId).Updates(map[string]interface{}{
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
