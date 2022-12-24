/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/24 18:07
 */

package Paginate

import (
	"github.com/jobhandsome/microSNS/common/define"
	"gorm.io/gorm"
)

// Paginate
// 统一分页封装
// return *gorm.DB
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if pageSize == 0 {
			pageSize = define.PageSize
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
