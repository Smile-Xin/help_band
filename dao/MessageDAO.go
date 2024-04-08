package dao

import (
	"backend/model"
	"backend/utils/errmsg"
)

// AddMessage 添加留言
func AddMessage(message model.Message) (code uint) {
	err := db.Create(&message).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

// QueryMessage 查询留言
func QueryMessage(lid int) (code uint, messageList []model.Message) {
	err := db.Where("lid = ?", lid).Find(&messageList).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL, messageList
	}
	return errmsg.SUCCESS, messageList
}
