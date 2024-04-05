package v1

import (
	server "backend/serve"
	"backend/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
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
