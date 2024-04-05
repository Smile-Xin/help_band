package model

import (
	"backend/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(64);not null" json:"user_name"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	Role     uint   `gorm:"type:int" json:"role"`
	Score    uint   `gorm:"type:int" json:"score"` // 用户评分
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password, err = utils.ScryptPW(u.Password)
	return
}
