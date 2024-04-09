package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddMessage 添加留言
func AddMessage(c *gin.Context) {
	var message model.Message
	c.ShouldBindJSON(&message)
	code := dao.AddMessage(message)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryMessage 查询留言
func QueryMessage(c *gin.Context) {
	lid, _ := strconv.Atoi(c.Param("lid"))
	code, messageList := dao.QueryMessage(lid)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    messageList,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryMessageByTwoUserId 查询留言
func QueryMessageByTwoUserId(c *gin.Context) {
	userAId, _ := strconv.Atoi(c.Param("userAId"))
	userBId, _ := strconv.Atoi(c.Param("userBId"))
	code, letterList := dao.QueryMessageByTwoUserId(uint(userAId), uint(userBId))
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"data":    letterList,
		"message": errmsg.GetErrMsg(code),
	})
}
