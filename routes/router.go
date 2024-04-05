package routes

import (
	"backend/api/v1"
	"backend/middleware"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 设置项目状态
	gin.SetMode(utils.AppMode)
	//创建新的router
	router := gin.New()
	router.Use(middleware.Cors())
	router.Use(middleware.Logger())

	group1 := router.Group("api/v1")
	{
		// user业务
		group1.GET("user/QueryByName/:name", v1.QueryUserByName)
		// 添加用户
		group1.POST("user/AddUser", v1.AddUser)
		// 登录
		group1.POST("user/Login", v1.Login)

		// task业务
		// 获取全部任务(未接受） 有分页查询 pageSize:每页数量 pageNum:页码
		group1.GET("task/GetAll/:pageSize/:pageNum", v1.GetAll)
		// 根据tag查询任务 有分页查询
		group1.GET("task/QueryTaskByTage/:pageSize/:pageNum/:tag", v1.QueryTaskByTag)
		// 根据提出问题者查询任务
		group1.GET("task/QueryTaskByDemander/:pageSize/:pageNum/:demander", v1.QueryTaskByDemander)
		// 根据接受者查询任务
		group1.GET("task/QueryTaskByReceiver/:pageSize/:pageNum/:recipient", v1.QueryTaskByReceiver)
		// 根据id查询任务
		group1.GET("task/QueryTaskByID/:id", v1.QueryTaskByID)
		// 根据id和状态查询任务
		group1.GET("task/QueryTaskByUserIdStatus/:pageSize/:pageNum/:userName/:status", v1.QueryTaskByUserIdStatus)

		// taskComment业务
		// 根据taskID查询评论
		group1.GET("taskComment/QueryCommentByTaskID/:taskID", v1.QueryTaskCommentByTaskID)
		// 根据Receiver（被评论者）查询评论
		group1.GET("taskComment/QueryCommentByReceiver/:receiver", v1.QueryTaskCommentByReceiver)
	}
	group2 := router.Group("api/v1", middleware.JwtToken())
	{
		// user业务
		// 获取全部用户
		group2.GET("user/GetAll", v1.GetAllUser)
		// 修改用户
		group2.POST("user/EditUser", v1.EditUser)
		// 删除用户
		group2.DELETE("user/DeleteUser/:name", v1.DeleteUser)

		// task业务

		// 添加任务
		group2.POST("task/AddTask", v1.AddTask)
		// 修改任务
		group2.POST("task/EditTask", v1.EditTask)
		// 删除任务
		group2.DELETE("task/DeleteTask/:id", v1.DeleteTask)

		// taskComment业务
		// 添加评论
		group2.POST("taskComment/AddTaskComment", v1.AddTaskComment)
		// 修改评论
		group2.POST("taskComment/EditTaskComment", v1.EditTaskComment)
		// 删除评论
		group2.DELETE("taskComment/DeleteTaskComment/:id", v1.DeleteTaskComment)

		// 文件上传
		group2.POST("qiniu/upload/:taskId", v1.Upload)
	}
	_ = router.Run(utils.HttpPort)
}
