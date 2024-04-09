package v1

import (
	"backend/dao"
	"backend/middleware"
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
		"state":   code,
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
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	userName := c.Param("name")
	code := dao.DeleteUser(userName)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// Login 登录
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	// 调用dao层登录
	u, code := dao.Login(&user)
	// 判断是否登录成功
	if code != errmsg.SUCCESS {
		c.JSON(200, gin.H{
			"state":   code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	// 登录成功，生成token
	token, code := middleware.JwtGenerateToken(user.UserName, time.Hour*7*24)
	c.JSON(200, gin.H{
		"state":   code,
		"token":   token,
		"name":    u.UserName,
		"id":      u.ID,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetAvatar 获取头像
func GetAvatar(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, avatar := dao.GetAvatar(uint(id))
	c.JSON(200, gin.H{
		"state":   code,
		"data":    avatar,
		"message": errmsg.GetErrMsg(code),
	})
}
