package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID       int    `gorm:"primaryKey" json:"id"`
	Lid      uint   `gorm:"type:int" json:"lid"` //letter id
	Letter   Letter `gorm:"foreignKey:Lid" json:"letter"`
	SenderID uint   `gorm:"type:int" json:"sender_id"` // 发送者id
	Content  string `gorm:"type:varchar(200)" json:"content"`
}
