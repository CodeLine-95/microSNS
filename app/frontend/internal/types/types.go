// Code generated by goctl. DO NOT EDIT.
package types

type CommonResply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterReq struct {
	Email       string `json:"email"`        // 用户邮箱
	Pass        string `json:"pass"`         // 用户密码，加密
	ConfirmPass string `json:"confirm_pass"` // 确认密码
	Code        string `json:"code"`         // 邮箱验证码
}

type LoginReq struct {
	Email string `json:"email"` // 用户邮箱
	Pass  string `json:"pass"`  // 用户密码
}

type SenEmailCodeReq struct {
	Email string `json:"email"`
}

type UserCheckinsReq struct {
}

type UserIntegralReq struct {
	Page     int `json:"page,optional"`
	PageSize int `json:"pageSize,optional"`
}

type UserIntegralResply struct {
	CommonResply
	TotalCount int64              `json:"totalCount"`
	CurrCount  int                `json:"currCount"`
	Data       []UserIntegralItem `json:"data"`
}

type UserIntegralItem struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Rewards   uint64 `json:"rewards"`
	Mode      string `json:"mode"`
	CreatedAt string `json:"created_at"`
}

type CateListsReq struct {
	Page     int `json:"page,optional"`
	PageSize int `json:"pageSize,optional"`
}

type CateListsResply struct {
	CommonResply
	TotalCount int64      `json:"totalCount"`
	CurrCount  int        `json:"currCount"`
	Data       []CateItem `json:"data"`
}

type CateItem struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`       // 名称
	Alias     string `json:"alias"`      // 别名
	Desc      string `json:"desc"`       // 介绍
	Pid       uint8  `json:"pid"`        // 上级节点
	Level     uint8  `json:"level"`      // 层级：0一级(默认)
	Icons     string `json:"icons"`      // 图标
	Sorts     uint32 `json:"sorts"`      // 排序
	State     uint8  `json:"state"`      // 状态：0-关闭/1-开启
	IsDelete  uint8  `json:"is_delete"`  // 软删除
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
}

type TopicCreateReq struct {
	CateId    uint32 `json:"cate_id"`    // 所属板块ID
	Title     string `json:"title"`      // 标题
	Tags      string `json:"tags"`       // 标签
	Content   string `json:"content"`    // 内容
	MdContent string `json:"md_content"` // markdown内容
}

type TopicListsReq struct {
	Page     int `json:"page,optional"`
	PageSize int `json:"pageSize,optional"`
}

type TopicListsResply struct {
	CommonResply
	TotalCount int64       `json:"totalCount"`
	CurrCount  int         `json:"currCount"`
	Data       []TopicItem `json:"data"`
}

type TopicDetailResply struct {
	CommonResply
	Data TopicItem `json:"data"`
}

type TopicItemReq struct {
	TopicId uint32 `json:"topic_id"`
}

type TopicItem struct {
	Id           uint32 `json:"id"`
	CateId       uint   `json:"cate_id"`       // 版块ID
	UserId       uint   `json:"user_id"`       // 用户ID
	Title        string `json:"title"`         // 标题
	Tags         string `json:"tags"`          // 标签
	State        uint8  `json:"state"`         // 状态：0-草稿/1-发布
	Type         uint8  `json:"type"`          // 类型：0-默认/1-精华/2-置顶
	Content      string `json:"content"`       // 内容
	MdContent    string `json:"md_content"`    // markdown内容
	IsDelete     uint8  `json:"is_delete"`     // 软删除
	LastReplyAt  string `json:"last_reply_at"` // 最后回复时间
	ReplyId      string `json:"reply_id"`      // 最后回复者ID
	ViewCount    uint64 `json:"view_count"`    // 浏览量
	LikeCount    uint64 `json:"like_count"`    // 点赞量
	CommentCount uint64 `json:"comment_count"` // 评论量
	CreatedAt    string `json:"created_at"`    // 创建时间
	UpdatedAt    string `json:"updated_at"`    // 更新时间
}
