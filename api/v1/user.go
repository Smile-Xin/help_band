package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
)

// QueryUserByName 查询用户
func QueryUserByName(c *gin.Context) {
	name := c.Param("name")
	user, code := dao.QueryUserByName(name)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    user,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetAllUser 获取全部用户
func GetAllUser(c *gin.Context) {
	// 调用dao层获取全部用户
	users, code := dao.GetAllUser()
	c.JSON(200, gin.H{
		"state":   code,
		"data":    users,
		"message": errmsg.GetErrMsg(code),
	})
}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	// 绑定user错误
	if err != nil {
		fmt.Printf("bind user fail %s", err)
		return
	}
	// 调用dao层添加用户
	code := dao.AddUser(&user)
	c.JSON(200, gin.H{
		"stats":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 修改用户
func EditUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	// 绑定user错误
	if err != nil {
		fmt.Printf("bind user fail %s", err)
		return
	}
	// 调用dao层修改用户
	code := dao.EditUser(&user)
	c.JSON(200, gin.H{
		"stats":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userName := c.Param("name")
	code := dao.DeleteUser(userName)
	c.JSON(200, gin.H{
		"stats":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
