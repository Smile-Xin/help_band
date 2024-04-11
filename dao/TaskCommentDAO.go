package dao

import (
	"backend/model"
	"backend/utils/errmsg"
	"fmt"
)

// ExistTaskCommentByID 判断任务评论是否存在
func ExistTaskCommentByID(id int) bool {
	result := db.Where("id = ?", id).Find(&model.TaskComment{})
	fmt.Println("result.RowsAffected", result.RowsAffected)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// AddTaskComment 添加任务评论
func AddTaskComment(taskComment *model.TaskComment) (code uint) {
	// 判断是否存在要评论的任务
	if !ExistTaskById(taskComment.TaskId) {
		return errmsg.TASK_NOT_EXIST
	}
	err := db.Create(&taskComment).Error
	if err != nil {
		return errmsg.DATABASE_WRITE_FAIL
	}
	return errmsg.SUCCESS
}

//	EditTaskComment 编辑任务评论
func EditTaskComment(taskComment *model.TaskComment) (code uint) {
	// 判断评论是否存在
	if ExistTaskCommentByID(int(taskComment.ID)) {
		// 存在
		err := db.Model(&model.TaskComment{}).Where("id = ?", taskComment.ID).Updates(&taskComment).Error
		if err != nil {
			return errmsg.DATABASE_WRITE_FAIL
		}
		return errmsg.SUCCESS
	} else {
		return errmsg.TASK_COMMENT_NOT_EXIST
	}

}

//	DeleteTaskComment 删除任务评论
func DeleteTaskComment(id int) (code uint) {
	if ExistTaskCommentByID(id) {
		err := db.Where("id = ?", id).Delete(&model.TaskComment{}).Error
		if err != nil {
			return errmsg.DATABASE_WRITE_FAIL
		}
		return errmsg.SUCCESS
	} else {
		return errmsg.TASK_COMMENT_NOT_EXIST
	}

}

//	QueryTaskCommentByTaskID 根据taskID查询评论
func QueryTaskCommentByTaskID(taskID int) (taskComment model.TaskComment, code uint) {
	result := db.Where("task_id = ? and status != ?", taskID, -1).Find(&taskComment)
	if result.Error != nil {
		return taskComment, errmsg.DATABASE_WRITE_FAIL
	}
	// 判断是否存在
	//if result.RowsAffected == 0 {
	//	return taskComment, errmsg.TASK_COMMENT_NOT_EXIST
	//}
	return taskComment, errmsg.SUCCESS
}

//	QueryTaskCommentByReceiver 根据Receiver（被评论者）查询评论
func QueryTaskCommentByReceiver(receiver int) (taskCommentList []model.TaskComment, code uint) {
	result := db.Where("receiver_id = ?", receiver).Find(&taskCommentList)
	if result.Error != nil {
		return taskCommentList, errmsg.DATABASE_WRITE_FAIL
	}
	//fmt.Printf("result.RowsAffected:%x", result.RowsAffected)
	// 判断是否存在
	if result.RowsAffected == 0 {
		return taskCommentList, errmsg.TASK_COMMENT_NOT_EXIST
	}
	return taskCommentList, errmsg.SUCCESS
}
