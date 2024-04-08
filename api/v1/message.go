package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddMessage 添加留言
func AddMessage(c *gin.Context) {
	var message model.Message
	c.ShouldBindJSON(&message)
	code := dao.AddMessage(message)
	c.JSON(200, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryMessage 查询留言
func QueryMessage(c *gin.Context) {
	lid, _ := strconv.Atoi(c.Param("lid"))
	code, messageList := dao.QueryMessage(lid)
	c.JSON(200, gin.H{
		"code":    code,
		"data":    messageList,
		"message": errmsg.GetErrMsg(code),
	})
}
