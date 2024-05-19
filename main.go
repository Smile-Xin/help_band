package main

import (
	"backend/dao"
	"backend/routes"
)

func main() {
	//数据库
	dao.InitDb()

	//开启路由
	routes.InitRouter()

}
