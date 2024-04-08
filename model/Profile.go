package model

type Profile struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200);default:这个人很懒没有写介绍" json:"desc"`
	Qqchat string `gorm:"type:varchar(200);default:这是qq" json:"qq_chat"`
	Wechat string `gorm:"type:varchar(100);default:这是微信" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200);default:这是微博" json:"weibo"`
	Bili   string `gorm:"type:varchar(200);default:这是b站" json:"bili"`
	Email  string `gorm:"type:varchar(200);default:这是邮箱" json:"email"`
	Img    string `gorm:"type:varchar(200)" json:"img"`
	Avatar string `gorm:"type:varchar(200);default:http://sb3lbh5km.hb-bkt.clouddn.com/touxiang.jpg" json:"avatar"` // 头像
}
