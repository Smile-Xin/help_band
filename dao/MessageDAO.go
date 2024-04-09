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

// QueryMessageByTwoUserId 查询留言
func QueryMessageByTwoUserId(userAId uint, userBId uint) (code uint, messageList []model.Message) {
	// 由俩个用户id得到获取uid
	uid := model.GetUid(userAId, userBId)
	// 判断是否存在这封信
	if !ExistLetter(userAId, userBId) {
		code = errmsg.LETTER_NOT_EXIST
		return
	}
	// 通过uid查询Message列表
	err := db.Where("lid = ?", db.Select("id").Where("uid = ?", uid).Table("letter")).Find(&messageList).Error
	if err != nil {
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}
