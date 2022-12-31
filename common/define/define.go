/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/24 18:02
 */

package define

var (
	// MailPassword 邮件发送者密码
	MailPassword = ""

	// PageSize 分页的默认参数
	PageSize = 10

	// CodeLength 验证码长度
	CodeLength = 6

	// CodeExpire 验证码过期时间（s）
	CodeExpire = 300

	// TimeZone 默认时区
	TimeZone = "Asia/Shanghai"

	// DatetimeFormat 时间格式化参数
	DatetimeFormat = "2006-01-02 15:04:05"

	// DateDayFormat 时间格式化参数
	DateDayFormat = "2006-01-02"

	// TokenExpire Token 过期时间
	TokenExpire = 3600

	// RefreshTokenExpire Token 刷新过期时间
	RefreshTokenExpire = 7200
)
