package model

type Message struct {
	ID       int  `gorm:"primaryKey" json:"id"`
	Lid      uint `gorm:"type:int" json:"lid"`
	SenderID uint `gorm:"type:int" json:"sender_id"`
}
