package model

import "gorm.io/gorm"

type TaskComment struct {
	gorm.Model
	TaskId        uint   `gorm:"type:int;not null" json:"task_id"`                //评论的任务id
	AppraiserId   uint   `gorm:"type:int;not null" json:"appraiser_id"`           //评价者id
	AppraiserName string `gorm:"type:varchar(64);not null" json:"appraiser_name"` //评价者name
	ReceiverId    uint   `gorm:"type:int;not null" json:"receiver_id"`            //被评价者id
	ReceiverName  string `gorm:"type:varchar(64);not null" json:"receiver_name"`  //被评价者name
	Content       string `gorm:"type:longtext" json:"content"`                    //评价的具体内容
	Score         uint   `gorm:"type:tinyint" json:"score"`                       //评分
	Status        uint   `gorm:"type:tinyint;default:0" json:"status"`            //状态码 0：未评论 1：这条评论是买家评论的 2：这条评论是卖家评论的 -1：这条评论是退回的信息
}
