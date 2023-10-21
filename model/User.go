package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(64);not null" json:"user_name"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Role     uint   `gorm:"type:int" json:"role"`
	Score    uint   `gorm:"type:int" json:"score"`
}
