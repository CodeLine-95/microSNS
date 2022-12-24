// Code generated by goctl. DO NOT EDIT.
package types

type CommonResply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterReq struct {
	Name string `json:"name"` // 用户名
	Pass string `json:"pass"` // 用户密码，加密
}

type LoginReq struct {
	Name string `json:"name"` // 用户名
	Pass string `json:"pass"` // 用户密码
}

type CateListsReq struct {
	Page     int `json:"page,optional"`
	PageSize int `json:"pageSize,optional"`
}

type CateListsResply struct {
	CommonResply
	TotalCount int64           `json:"totalCount"`
	CurrCount  int             `json:"currCount"`
	Data       []CateListsItem `json:"data"`
}

type CateListsItem struct {
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
