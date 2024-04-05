package dao

import (
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
)

// ExistTaskById 判断任务是否存在
func ExistTaskById(id uint) bool {
	result := db.Where("id = ?", id).Find(&model.Task{})

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// QueryTaskByTag 查询任务列表
/*
pageSize: 每页数量	pageNum: 页码	tag: 关键词
判断是否存在关键词，存在则模糊查询关键词，不存在则查询全部
*/
func QueryTaskByTag(pageSize int, pageNum int, tag string) (taskList []model.Task, total int64, code uint) {
	if tag != "" {
		// 总人数
		err := db.Where("tag like ?", "%"+tag+"%").Find(&taskList).Count(&total).Error
		if err != nil {
			fmt.Printf("get task total fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		// 查询
		err = db.Where("tag like ?", "%"+tag+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
		if err != nil {
			fmt.Printf("query task fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
	} else {
		// 没有关键词的总人数
		err := db.Find(&taskList).Count(&total).Error
		if err != nil {
			fmt.Printf("get task total fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		// 没有关键词查询
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
		if err != nil {
			fmt.Printf("query task fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
	}
	code = errmsg.SUCCESS
	return
}

// GetAll 获取全部任务(未接受)
func GetAll(pageSize int, pageNum int) (taskList []model.Task, total int64, code uint) {
	// 总人数
	err := db.Where("`status` = ?", 0).Find(&taskList).Count(&total).Error
	if err != nil {
		fmt.Printf("get task total fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	// 查询
	err = db.Where("`status` = ?", 0).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
	if err != nil {
		fmt.Printf("query task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

// QueryTaskByID 根据id查询任务
func QueryTaskByID(id int) (task model.Task, code uint) {
	if ExistTaskCommentByID(id) {
		code = errmsg.TASK_NOT_EXIST
		return
	}
	err := db.Where("id = ?", id).Find(&task).Error
	if err != nil {
		fmt.Printf("query task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return

}

// QueryTaskByUserIdStatus 根据userId和状态查询任务
func QueryTaskByUserIdStatus(pageSize int, pageNum int, userName string, status int, title string) (taskList []model.Task, total int64, code uint) {
	// 判断是否存在该用户
	if !ExistUser(userName) {
		code = errmsg.INEXISTENCE_USER
		return
	}
	// 总人数
	err := db.Where("(demander_name = ? OR recipient_name = ?) AND status = ? AND title LIKE ?", userName, userName, status, "%"+title+"%").Find(&taskList).Count(&total).Error
	if err != nil {
		fmt.Printf("get task total fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return

	}
	// 查询
	err = db.Where("(demander_name = ? OR recipient_name = ?) AND status = ? AND title LIKE ?", userName, userName, status, "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
	code = errmsg.SUCCESS
	return

}

// QueryTaskByDemander 根据提出问题者查询任务
func QueryTaskByDemander(pageSize int, pageNum int, demander string, title string) (taskList []model.Task, total int64, code uint) {
	if !ExistUser(demander) {
		code = errmsg.INEXISTENCE_USER
		return
	}
	// 总人数
	err := db.Where("demander_name = ? AND title LIKE ?", demander, "%"+title+"%").Find(&taskList).Count(&total).Error
	if err != nil {
		fmt.Printf("get task total fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	// 查询
	err = db.Where("demander_name = ? AND title LIKE ?", demander, "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
	if err != nil {
		fmt.Printf("query task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

// QueryTaskByReceiver 根据接受者查询任务
func QueryTaskByReceiver(pageSize int, pageNum int, recipient_name string, title string) (taskList []model.Task, total int64, code uint) {
	if !ExistUser(recipient_name) {
		code = errmsg.INEXISTENCE_USER
		return
	}
	// 总人数
	err := db.Where("recipient_name = ? AND title LIKE ?", recipient_name, "%"+title+"%").Find(&taskList).Count(&total).Error
	if err != nil {
		fmt.Printf("get task total fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	// 查询
	err = db.Where("recipient_name = ? AND title LIKE ?", recipient_name, "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at desc").Find(&taskList).Error
	if err != nil {
		fmt.Printf("query task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return

}

// AddTask 添加任务
func AddTask(task *model.Task) (code uint) {
	err := db.Create(&task).Error
	if err != nil {
		fmt.Printf("add task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

// EditTask 修改任务
func EditTask(task *model.Task) (code uint) {
	if ExistTaskById(task.ID) {
		//err := db.Save(&task).Error
		db.Model(&model.Task{}).Where("id = ?", task.ID).Updates(&task)
		if err != nil {
			fmt.Printf("edit task fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
		return
	} else {
		code = errmsg.TASK_NOT_EXIST
		return
	}

}

// DeleteTask 删除任务
func DeleteTask(id int) (code uint) {
	err := db.Where("id = ?", id).Delete(&model.Task{}).Error
	if err != nil {
		fmt.Printf("delete task fail: %s", err)
		code = errmsg.DATABASE_WRITE_FAIL
		return
	}
	code = errmsg.SUCCESS
	return
}

// UploadArticle 上传文章
func UploadArticle(taskId int, article string) (code uint) {
	if ExistTaskById(uint(taskId)) {
		var task model.Task
		db.Where("id = ?", taskId).Find(&task)
		// 任务状态为重新提交文件
		if task.Status == 5 {
			err := db.Model(&model.Task{}).Where("id = ?", taskId).Updates(model.Task{Article: article, Status: 6}).Error
			if err != nil {
				fmt.Printf("upload article fail: %s", err)
				code = errmsg.DATABASE_WRITE_FAIL
				return
			}
			code = errmsg.SUCCESS
			return
		}
		// 任务状态为提交文件
		err := db.Model(&model.Task{}).Where("id = ?", taskId).Updates(model.Task{Article: article, Status: 2}).Error
		if err != nil {
			fmt.Printf("upload article fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
		code = errmsg.SUCCESS
		return
	} else {
		code = errmsg.TASK_NOT_EXIST
		return
	}
}
