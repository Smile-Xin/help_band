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
	fmt.Println("AddLetter", letter)
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
	userAId, _ := strconv.Atoi(c.Param("userId1"))
	userBId, _ := strconv.Atoi(c.Param("userId2"))
	fmt.Println("userAId", userAId, "userBId", userBId)
	code, letterList := dao.QueryLetterByTwoUserId(uint(userAId), uint(userBId))
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

// QueryLetterById 查询信件
func QueryLetterById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code, letter := dao.QueryLetterById(id)
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letter,
		"message": errmsg.GetErrMsg(code),
	})
}
