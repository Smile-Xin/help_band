package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	//Cid       uint        `gorm:"type:int;not null" json:"cid"`
	//Comment     TaskComment `gorm:"foreignKey:Cid"`                         //任务评价
	Tag           string `gorm:"type:varchar(64);not null" json:"tag"`   //任务标签
	Title         string `gorm:"type:varchar(64);not null" json:"title"` //任务标题
	Briefing      string `gorm:"type:varchar(255)" json:"briefing"`      //任务简介
	Content       string `gorm:"type:longtext" json:"content"`           //任务主要内容
	DemanderId    uint   `gorm:"type:int" json:"demander_id"`            //提出问题者
	DemanderName  string `gorm:"type:varchar(64)" json:"demander_name"`  //提出问题者
	RecipientId   uint   `gorm:"type:int" json:"recipient_id"`           //接受问题者
	RecipientName string `gorm:"type:varchar(64)" json:"recipient_name"` //接受问题者
	Status        int8   `gorm:"type:tinyint;default:0" json:"status"`   //任务当前状态：0：发布未被接受 1：被接受 2：提交文件 3：已完成 4：已评论 5：被退回 6：已重新提交文件 -1：已取消
	Article       string `gorm:"type:varchar(64)" json:"article"`        //文章链接
	//relate      int8
}
