/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/24 18:49
 */

package model

// SnsAccount 管理员账户
type SnsAccount struct {
	Id        uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"column:name;type:varchar(255);comment:账户名;NOT NULL" json:"name"`
	Nicks     string `gorm:"column:nicks;type:varchar(255);comment:昵称;NOT NULL" json:"nicks"`
	Pass      string `gorm:"column:pass;type:varchar(255);comment:密码;NOT NULL" json:"pass"`
	Ip        string `gorm:"column:ip;type:varchar(255);comment:IP地址;NOT NULL" json:"ip"`
	State     uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:状态：0-冻结/1-正常;NOT NULL" json:"state"`
	IsDelete  uint   `gorm:"column:is_delete;type:tinyint(4) unsigned;default:0;comment:软删除：0未删；1删除;NOT NULL" json:"is_delete"`
	CreatedAt string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at;type:varchar(255);comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SnsAccount) TableName() string {
	return "sns_account"
}

// SnsCate 分类表
type SnsCate struct {
	Id        uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"column:name;type:varchar(20);comment:名称;NOT NULL" json:"name"`
	Alias     string `gorm:"column:alias;type:varchar(100);comment:别名;NOT NULL" json:"alias"`
	Desc      string `gorm:"column:desc;type:varchar(255);comment:介绍;NOT NULL" json:"desc"`
	Pid       uint   `gorm:"column:pid;type:tinyint(4) unsigned;default:0;comment:上级节点;NOT NULL" json:"pid"`
	Level     uint   `gorm:"column:level;type:tinyint(4) unsigned;default:0;comment:层级：0一级(默认);NOT NULL" json:"level"`
	Icons     string `gorm:"column:icons;type:varchar(255);comment:图标;NOT NULL" json:"icons"`
	Sorts     uint   `gorm:"column:sorts;type:int(11) unsigned;default:0;comment:排序;NOT NULL" json:"sorts"`
	State     uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:状态：0-关闭/1-开启;NOT NULL" json:"state"`
	IsDelete  uint   `gorm:"column:is_delete;type:tinyint(4) unsigned;default:0;comment:软删除;NOT NULL" json:"is_delete"`
	CreatedAt string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at;type:varchar(255);comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SnsCate) TableName() string {
	return "sns_cate"
}

// SnsCheckins 用户签到
type SnsCheckins struct {
	Id             uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId         uint   `gorm:"column:user_id;type:int(11) unsigned;comment:用户ID;NOT NULL" json:"user_id"`
	CumulativeDays uint64 `gorm:"column:cumulative_days;type:bigint(20) unsigned;default:0;comment:累积签到(天);NOT NULL" json:"cumulative_days"`
	ContinuityDays uint64 `gorm:"column:continuity_days;type:bigint(20) unsigned;default:0;comment:连续签到(天);NOT NULL" json:"continuity_days"`
	LastTime       string `gorm:"column:last_time;type:varchar(255);comment:最后签到时间;NOT NULL" json:"last_time"`
	CreatedAt      string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt      string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
}

func (m *SnsCheckins) TableName() string {
	return "sns_checkins"
}

// SnsFollows 用户关注
type SnsFollows struct {
	Id        uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    uint   `gorm:"column:user_id;type:int(11) unsigned;comment:用户ID;NOT NULL" json:"user_id"`
	TargetId  uint   `gorm:"column:target_id;type:int(11) unsigned;comment:被关注用户ID;NOT NULL" json:"target_id"`
	State     uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:状态：1-关注/0-取消;NOT NULL" json:"state"`
	CreatedAt string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
}

func (m *SnsFollows) TableName() string {
	return "sns_follows"
}

// SnsIntegralLogs 积分日志
type SnsIntegralLogs struct {
	Id        uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    uint   `gorm:"column:user_id;type:int(11) unsigned;comment:用户ID;NOT NULL" json:"user_id"`
	Rewards   uint64 `gorm:"column:rewards;type:bigint(20) unsigned;default:0;comment:奖励积分;NOT NULL" json:"rewards"`
	Mode      string `gorm:"column:mode;type:varchar(20);comment:获取方式;NOT NULL" json:"mode"`
	CreatedAt string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
}

func (m *SnsIntegralLogs) TableName() string {
	return "sns_integral_logs"
}

// SnsTopics 话题
type SnsTopics struct {
	Id          uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title       string `gorm:"column:title;type:varchar(255);comment:标题;NOT NULL" json:"title"`
	Tags        string `gorm:"column:tags;type:varchar(255);comment:标签;NOT NULL" json:"tags"`
	State       uint   `gorm:"column:state;type:tinyint(4) unsigned;default:0;comment:状态：0-草稿/1-发布;NOT NULL" json:"state"`
	Type        uint   `gorm:"column:type;type:tinyint(4) unsigned;default:0;comment:类型：0-默认/1-精华/2-置顶;NOT NULL" json:"type"`
	Content     string `gorm:"column:content;type:longtext;comment:内容;NOT NULL" json:"content"`
	MdContent   string `gorm:"column:md_content;type:longtext;comment:markdown内容;NOT NULL" json:"md_content"`
	IsDelete    uint   `gorm:"column:is_delete;type:tinyint(4) unsigned;default:0;comment:软删除;NOT NULL" json:"is_delete"`
	LastReplyAt string `gorm:"column:last_reply_at;type:varchar(255);comment:最后回复时间;NOT NULL" json:"last_reply_at"`
	CreatedAt   string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt   string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt   string `gorm:"column:deleted_at;type:varchar(255);comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SnsTopics) TableName() string {
	return "sns_topics"
}

// SnsTopicsUsers 话题-用户
type SnsTopicsUsers struct {
	Id           uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	TopicsId     uint   `gorm:"column:topics_id;type:int(11) unsigned;comment:话题ID;NOT NULL" json:"topics_id"`
	CateId       uint   `gorm:"column:cate_id;type:int(11) unsigned;comment:分类ID;NOT NULL" json:"cate_id"`
	UserId       uint   `gorm:"column:user_id;type:int(11) unsigned;comment:用户ID;NOT NULL" json:"user_id"`
	ReplyId      uint   `gorm:"column:reply_id;type:int(11) unsigned;comment:最后回复者ID;NOT NULL" json:"reply_id"`
	ViewCount    uint64 `gorm:"column:view_count;type:bigint(20) unsigned;default:0;comment:浏览统计;NOT NULL" json:"view_count"`
	LikeCount    uint64 `gorm:"column:like_count;type:bigint(20) unsigned;default:0;comment:喜欢统计;NOT NULL" json:"like_count"`
	CommentCount uint64 `gorm:"column:comment_count;type:bigint(20) unsigned;default:0;comment:评论统计;NOT NULL" json:"comment_count"`
	CreatedAt    string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt    string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
}

func (m *SnsTopicsUsers) TableName() string {
	return "sns_topics_users"
}

// SnsUsers 网站用户
type SnsUsers struct {
	Id        uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string `gorm:"column:name;type:varchar(50);comment:用户名;NOT NULL" json:"name"`
	Pass      string `gorm:"column:pass;type:varchar(255);comment:密码;NOT NULL" json:"pass"`
	Gender    uint   `gorm:"column:gender;type:tinyint(4) unsigned;default:0;comment:性别：0-未知/1-男/2-女/3-保密;NOT NULL" json:"gender"`
	City      string `gorm:"column:city;type:varchar(50);comment:城市;NOT NULL" json:"city"`
	Email     string `gorm:"column:email;type:varchar(100);comment:邮箱/用户名;NOT NULL" json:"email"`
	Phone     string `gorm:"column:phone;type:varchar(50);comment:手机/用户名;NOT NULL" json:"phone"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);comment:头像;NOT NULL" json:"avatar"`
	Site      string `gorm:"column:site;type:varchar(255);comment:网站;NOT NULL" json:"site"`
	Job       string `gorm:"column:job;type:varchar(50);comment:职业;NOT NULL" json:"job"`
	Desc      string `gorm:"column:desc;type:varchar(255);comment:简介;NOT NULL" json:"desc"`
	Integral  uint64 `gorm:"column:integral;type:bigint(20) unsigned;default:0;comment:个人积分;NOT NULL" json:"integral"`
	State     uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:状态：1-正常/0-禁用;NOT NULL" json:"state"`
	IsDelete  uint   `gorm:"column:is_delete;type:tinyint(4) unsigned;default:0;comment:软删除;NOT NULL" json:"is_delete"`
	Level     uint   `gorm:"column:level;type:tinyint(4) unsigned;default:1;comment:等级：1-20;NOT NULL" json:"level"`
	CreatedAt string `gorm:"column:created_at;type:varchar(255);comment:创建时间;NOT NULL" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at;type:varchar(255);comment:更新时间;NOT NULL" json:"updated_at"`
	DeletedAt string `gorm:"column:deleted_at;type:varchar(255);comment:删除时间;NOT NULL" json:"deleted_at"`
}

func (m *SnsUsers) TableName() string {
	return "sns_users"
}
