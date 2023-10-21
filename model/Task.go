package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	//Cid       uint        `gorm:"type:int;not null" json:"cid"`
	Comment     TaskComment `gorm:"foreignKey:Cid"`                         //任务评价
	Tag         string      `gorm:"type:varchar(64);not null" json:"tag"`   //任务标签
	Title       string      `gorm:"type:varchar(64);not null" json:"title"` //任务标题
	Content     string      `gorm:"type:longtext" json:"content"`           //任务主要内容
	DemanderId  uint        `gorm:"type:int;not null" json:"demander"`      //提出问题者
	RecipientId uint        `gorm:"type:int;not null" json:"recipient"`     //接受问题者
	Status      int8        `gorm:"type:tinyint;default:0" json:"status"`   //任务当前状态：0：发布未被接受 1：被接受 2：已完成未评论 3：已评论

}
