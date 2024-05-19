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

		// article业务
		group1.GET("article/query", v1.QueryArticle)
		group1.GET("article/list", v1.QueryArtList)
		group1.GET("article/get", v1.GetArticle)
		// category业务
		//group1.GET("category/get", v1.GetCategory)
		//group1.GET("category/query", v1.QueryCategory)
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
		// 获取用户头像
		group2.GET("user/GetAvatar/:id", v1.GetAvatar)

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

		// article业务
		//auth.GET("article/query", v1.QueryArticle)
		//auth.GET("article/list", v1.QueryArtList)
		//auth.GET("article/get", v1.GetArticle)
		group2.POST("article/add", v1.AddArticle)
		group2.POST("article/edit", v1.EditArticle)
		group2.DELETE("article/delete", v1.DeleteArticle)

		// category业务
		//auth.GET("category/get", v1.GetCategory)
		//auth.GET("category/query", v1.QueryCategory)
		//group2.POST("category/add", v1.AddCategory)
		//group2.POST("category/edit", v1.EditCategory)
		//group2.DELETE("category/delete", v1.DeleteCategory)

		// profile 业务
		//auth.GET("profile/:id", v1.GetProfile)
		group2.POST("profile/edit/:id", v1.EditProfile)
		group2.GET("admin/profile/:id", v1.GetProfile)

		// 私信业务
		// 添加私信
		group2.POST("message/AddMessage", v1.AddMessage)
		// 查询私信
		group2.GET("message/QueryMessage/:lid", v1.QueryMessage)
		// 根据两个userId查message
		group2.GET("letter/QueryMessageByTwoUserId/:userId1/:userId2", v1.QueryMessageByTwoUserId)

		// 添加对话
		group2.POST("letter/AddLetter", v1.AddLetter)
		// 根据userId查询对话
		group2.GET("letter/QueryLetterByUserId/:userId", v1.QueryLetterByUserId)
		// 根据两个userId查对话
		group2.GET("letter/QueryLetterByTwoUserId/:userId1/:userId2", v1.QueryLetterByTwoUserId)
		// 根据uid查询对话
		group2.GET("letter/QueryLetterByUid/:uid", v1.QueryLetterByUid)
		// 根据id查询对话
		group2.GET("letter/QueryLetterById/:id", v1.QueryLetterById)

		// 文件上传
		group2.POST("qiniu/upload/:taskId", v1.Upload)
		group2.POST("qiniu/uploadAvatar/:userName", v1.UploadAvatar)

	}
	_ = router.Run(utils.HttpPort)
}
