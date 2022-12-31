/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/12/27 14:35
 */

package model

import "strings"

func (m *SnsTopics) JoinSelectFields(fieldStr string) string {

	var SnsTopicUsers SnsTopicsUsers

	returnStr := m.TableName() + ".*"
	if len(fieldStr) == 0 {
		return returnStr
	}

	fields := strings.Split(fieldStr, ",")

	returnStr = returnStr + ","

	for _, s := range fields {
		returnStr += SnsTopicUsers.TableName() + "." + s + ","
	}

	return strings.TrimRight(returnStr, ",")
}
