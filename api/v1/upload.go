package v1

import (
	"backend/dao"
	server "backend/serve"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Upload(c *gin.Context) {
	var code uint
	var url string
	taskId := c.Param("taskId")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		code = errmsg.TRANSPORT_ERR
	} else {
		url, code = server.Upload(file, fileHeader, taskId)
		if code == errmsg.SUCCESS {
			id, _ := strconv.Atoi(taskId)
			code = dao.UploadArticle(id, url)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}

func UploadAvatar(c *gin.Context) {
	var code uint
	var url string
	userName := c.Param("userName")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		code = errmsg.TRANSPORT_ERR
	} else {
		url, code = server.Upload1(file, fileHeader, userName)
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}

func UploadAli(c *gin.Context) {
	var code uint
	var url string
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		code = errmsg.TRANSPORT_ERR
	} else {
		code, url = server.UploadAli(file, fileHeader.Filename)
	}
	c.JSON(http.StatusOK, gin.H{
		"state":   code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
