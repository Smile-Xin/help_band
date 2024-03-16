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

// ExistTaskById 判断任务是否存在

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
		err = db.Where("tag like ?", "%"+tag+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&taskList).Error
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
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&taskList).Error
		if err != nil {
			fmt.Printf("query task fail: %s", err)
			code = errmsg.DATABASE_WRITE_FAIL
			return
		}
	}
	code = errmsg.SUCCESS
	return
}

// GetAll 获取全部任务
func GetAll() (taskList []model.Task, code uint) {
	err := db.Find(&taskList).Error
	if err != nil {
		fmt.Printf("get all task fail: %s", err)
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
