package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Letter struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserA_id   uint   `gorm:"type:int;not null" json:"userA_id"`
	UserA      User   `gorm:"foreignKey:UserA_id" json:"userA"`
	UserB      User   `gorm:"foreignKey:UserB_id" json:"userB"`
	UserB_id   uint   `gorm:"type:int;not null" json:"userB_id"`
	UserA_name string `gorm:"type:varchar(64)" json:"userA_name"`
	UserB_name string `gorm:"type:varchar(64)" json:"userB_name"`
	Uid        uint64 `gorm:"type:int" json:"uid"`
}

func (l *Letter) BeforeSave(tx *gorm.DB) (err error) {
	l.Uid = GetUid(l.UserA_id, l.UserB_id)
	return
}

func GetUid(id1 uint, id2 uint) (Uid uint64) {
	fmt.Println(id1, id2)
	if id1 > id2 {
		Uid = uint64(id2 | (id1 << 32))
	} else {
		Uid = uint64(id1 | (id2 << 32))
	}
	fmt.Println("Uid", Uid)
	return
}
