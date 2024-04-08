package dao

import (
	"backend/model"
	"backend/utils/errmsg"
)

func ExistLetter(userAId uint, userBId uint) bool {
	var letter model.Letter
	uid := model.GetUid(userAId, userBId)
	db.Where("uid = ?", uid).First(&model.Letter{})
	if letter.ID > 0 {
		return true
	}
	return false
}

// AddLetter 添加信件
func AddLetter(letter model.Letter) (code uint) {
	err := db.Create(&letter).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

// QueryLetterByUid 查询信件
func QueryLetterByUid(uid int) (code uint, letter model.Letter) {
	err := db.Where("uid = ?", uid).Find(&letter).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL, letter
	}
	return errmsg.SUCCESS, letter
}

// QueryLetterByTwoUserId 查询信件
func QueryLetterByTwoUserId(userAId int, userBId int) (code uint, letter model.Letter) {
	err := db.Where("userA_id = ? and userB_id = ?", userAId, userBId).Or("userA_id = ? and userB_id = ?", userBId, userAId).Find(&letter).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL, letter
	}
	return errmsg.SUCCESS, letter
}

// QueryLetterByUserId 查询信件
func QueryLetterByUserId(userId int) (letterList []model.Letter, code uint) {
	err := db.Where("userA_id = ?", userId).Or("userB_id = ?", userId).Find(&letterList).Error
	if err != nil {
		return letterList, errmsg.DATABASE_WRITE_FAIL
	}
	return letterList, errmsg.SUCCESS
}
