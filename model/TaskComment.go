package model

import "gorm.io/gorm"

type TaskComment struct {
	gorm.Model
	TaskId      uint   `gorm:"type:int;not null" json:"task_id"`
	AppraiserId uint   `gorm:"type:int;not null" json:"appraiser_id"` //评价者
	ReceiverId  uint   `gorm:"type:int;not null" json:"receiver_id"`  //被评价者
	Content     string `gorm:"type:longtext" json:"content"`          //评价的具体内容
	Score       uint   `gorm:"type:tinyint" json:"score"`             //评分
	Status      uint   `gorm:"type:tinyint;default:0" json:"status"`  //状态码 0：未评论 1：买家评论 2：卖家评论
}
