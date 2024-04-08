package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddLetter 创建信件
func AddLetter(c *gin.Context) {
	var letter model.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		fmt.Printf("bind letter fail:%s", err)
		return
	}
	code := dao.AddLetter(letter)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letter,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryLetterByUid 查询信件
func QueryLetterByUid(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Param("uid"))
	code, letterList := dao.QueryLetterByUid(uid)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letterList,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryLetterByTwoUserId 查询信件
func QueryLetterByTwoUserId(c *gin.Context) {
	userAId, _ := strconv.Atoi(c.Param("userAId"))
	userBId, _ := strconv.Atoi(c.Param("userBId"))
	code, letterList := dao.QueryLetterByTwoUserId(userAId, userBId)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letterList,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryLetterByUserId 查询信件
func QueryLetterByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	letterList, code := dao.QueryLetterByUserId(userId)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letterList,
		"message": errmsg.GetErrMsg(code),
	})
}
