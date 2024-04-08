package dao

import (
	. "backend/model"
	"backend/utils"
	"backend/utils/errmsg"
	"fmt"
)

// ExistUser User名字查重
func ExistUser(userName string) bool {
	var user User
	result := db.Where("user_name = ?", userName).Find(&user)
	if result.Error != nil {
		fmt.Printf("find user false %s", result.Error)
	}
	//fmt.Printf("ExistUser name %s", user.UserName)
	if user.ID > 0 {
		return true
	} else {
		return false
	}
}

// QueryUserByName 用name查user，非模糊查询
func QueryUserByName(name string) (user User, code uint) {
	// 计算用户评分
	db.Model(&User{}).Where("user_name = ?", name).Update("score",
		db.Table("task_comment").Select("AVG(score)").Where(map[string]interface{}{"receiver_name": name}).Not(map[string]interface{}{"score": 0}),
	)
	code = errmsg.SUCCESS
	result := db.Preload("Profile").Where("user_name = ?", name).Find(&user)
	if result.Error != nil {
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	//判断是否有该用户
	if result.RowsAffected < 1 {
		code = errmsg.INEXISTENCE_USER
		return
	}
	return user, code
}

// GetAllUser 获取用户列表
func GetAllUser() (users []User, code uint) {
	err := db.Find(&users)
	if err.Error != nil {
		fmt.Printf("get all user fail %s", err.Error)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

// AddUser 添加用户
func AddUser(user *User) (code uint) {
	if ExistUser(user.UserName) {
		code = errmsg.EXIST_USER
	} else {
		var profile Profile
		profile.Name = user.UserName
		err = db.Create(&profile).Error
		if err != nil {
			fmt.Printf("creat profile fail %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		user.Pid = profile.ID
		err := db.Create(&user).Error
		if err != nil {
			fmt.Printf("creat user fail %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}

		code = errmsg.SUCCESS
	}

	return code
}

// EditUser 修改用户
func EditUser(user *User) (code uint) {
	// 检查是否存在用户
	if !ExistUser(user.UserName) {
		fmt.Printf("用户修改失败，不存在用户")
		code = errmsg.INEXISTENCE_USER
		return
	}
	// 使用gorm修改用户
	result := db.Model(&user).Updates(map[string]interface{}{
		"password": user.Password,
		"role":     user.Role,
		"score":    user.Score,
	})
	if result.Error != nil {
		fmt.Printf("user updates fail:%s", result.Error)
		code = errmsg.DATABASE_WRITE_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

// DeleteUser 删除用户
func DeleteUser(name string) (code uint) {
	// 检查是否存在用户
	if !ExistUser(name) {
		fmt.Printf("用户删除失败，不存在用户")
		code = errmsg.INEXISTENCE_USER
		return
	}
	// 使用gorm删除用户
	result := db.Delete(&User{}, "user_name", name)
	if result.Error != nil {
		fmt.Printf("delete user fial:%s", result.Error)
		code = errmsg.DATABASE_WRITE_FAIL
	} else {
		code = errmsg.SUCCESS
	}
	return
}

// Login 登录
func Login(user *User) (u User, code uint) {
	if !ExistUser(user.UserName) {
		code = errmsg.INEXISTENCE_USER
		return
	}
	result := db.Where("user_name = ?", user.UserName).First(&u)
	// 数据库错误
	if result.Error != nil {
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	// 用户不存在
	if u.ID == 0 {
		code = errmsg.INEXISTENCE_USER
		return
	}

	// 密码加密
	Password, _ := utils.ScryptPW(user.Password)
	// 密码错误
	//fmt.Printf("u.Password=%s, Password=%s", u.Password, Password)
	if u.Password != Password {
		code = errmsg.ERROR_PASSWORD
		return
	}
	// 清空密码
	u.Password = ""

	code = errmsg.SUCCESS
	return
}

// ExamineRole 检查用户权限
func ExamineRole(name string, role uint) bool {
	var user User
	db.Where("user_name = ?", name).First(&user)
	if user.ID == 0 {
		return false
	}
	if user.Role >= role {
		return true
	}
	return false
}
