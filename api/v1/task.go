package v1

import (
	"backend/dao"
	"backend/model"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// QueryTaskByTag 查询任务列表
func QueryTaskByTag(c *gin.Context) {
	/*pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	tag := c.Query("tag")*/
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	tag := c.Param("tag")
	taskList, total, code := dao.QueryTaskByTag(pageSize, pageNum, tag)
	c.JSON(200, gin.H{
		"state":   code,
		"data":    taskList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetAll 获取全部任务
func GetAll(c *gin.Context) {
	taskList, code := dao.GetAll()
	c.JSON(200, gin.H{
		"state":   code,
		"data":    taskList,
		"message": errmsg.GetErrMsg(code),
	})

}

// AddTask 添加任务
func AddTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	code := dao.AddTask(&task)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditTask 修改任务
func EditTask(c *gin.Context) {
	var task model.Task
	_ = c.ShouldBindJSON(&task)
	code := dao.EditTask(&task)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteTask 删除任务
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := dao.DeleteTask(id)
	c.JSON(200, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
	})
}
