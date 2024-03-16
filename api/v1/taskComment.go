package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddTaskComment 添加任务评论
func AddTaskComment(c *gin.Context) {
	var taskComment model.TaskComment
	_ = c.ShouldBindJSON(&taskComment)
	code := dao.AddTaskComment(&taskComment)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditTaskComment 修改任务评论
func EditTaskComment(c *gin.Context) {
	var taskComment model.TaskComment
	_ = c.ShouldBindJSON(&taskComment)
	code := dao.EditTaskComment(&taskComment)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteTaskComment 删除任务评论
func DeleteTaskComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := dao.DeleteTaskComment(id)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryTaskCommentByTaskID 根据taskID查询评论
func QueryTaskCommentByTaskID(c *gin.Context) {
	taskID, _ := strconv.Atoi(c.Param("taskID"))
	taskComment, code := dao.QueryTaskCommentByTaskID(taskID)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    taskComment,
		"message": errmsg.GetErrMsg(code),
	})
}

// QueryTaskCommentByReceiver 根据Receiver（被评论者）查询评论
func QueryTaskCommentByReceiver(c *gin.Context) {
	receiver, _ := strconv.Atoi(c.Param("receiver"))
	taskComment, code := dao.QueryTaskCommentByReceiver(receiver)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    taskComment,
		"message": errmsg.GetErrMsg(code),
	})
}
