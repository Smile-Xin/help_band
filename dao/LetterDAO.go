package dao

import (
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
)

func ExistLetter(userAId uint, userBId uint) bool {
	var letter model.Letter
	uid := model.GetUid(userAId, userBId)
	db.Where("uid = ?", uid).First(&letter)
	fmt.Println("letter.ID", letter.ID)
	if letter.ID == 0 {
		return false
	}
	return true
}

// AddLetter 添加信件
func AddLetter(letter model.Letter) (code uint) {
	fmt.Println("AddLetter2", letter)
	if ExistLetter(letter.UserA_id, letter.UserB_id) {
		return errmsg.LETTER_EXIST
	}
	if letter.UserA_id == letter.UserB_id {
		return errmsg.LETTER_USER_SAME
	}
	// 初始化letter的UserA_name和UserB_name
	letter.UserA_name = GetUserName(letter.UserA_id)
	letter.UserB_name = GetUserName(letter.UserB_id)
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
func QueryLetterByTwoUserId(userAId uint, userBId uint) (code uint, letter model.Letter) {
	uid := model.GetUid(userAId, userBId)
	if !ExistLetter(userAId, userBId) {
		return errmsg.LETTER_NOT_EXIST, letter
	}
	err := db.Where("uid = ?", uid).Find(&letter).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL, letter
	}
	return errmsg.SUCCESS, letter
}

// QueryLetterByUserId 查询信件
func QueryLetterByUserId(userId int) (letterList []model.Letter, code uint) {
	err := db.Where("userA_id = ? || userB_id = ?", userId, userId).Find(&letterList).Error
	if err != nil {
		return letterList, errmsg.DATABASE_WRITE_FAIL
	}
	return letterList, errmsg.SUCCESS
}

// QueryLetterById 查询信件
func QueryLetterById(id int) (code uint, letter model.Letter) {
	err := db.Where("id = ?", id).First(&letter).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL, letter
	}
	return errmsg.SUCCESS, letter
}
