package model

import (
	"gorm.io/gorm"
)

type Letter struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserA_id uint   `gorm:"type:int;not null" json:"userA_id"`
	UserB_id uint   `gorm:"type:int;not null" json:"userB_id"`
	Uid      uint64 `gorm:"type:int" json:"uid"`
}

func (l *Letter) BeforeSave(tx *gorm.DB) (err error) {
	l.Uid = GetUid(l.UserA_id, l.UserB_id)
	return
}

func GetUid(id1 uint, id2 uint) (Uid uint64) {
	if id1 > id2 {
		Uid = uint64(id2) | uint64(id1)<<32
	} else {
		Uid = uint64(id1) | uint64(id2)<<32
	}
	return
}
