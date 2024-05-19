package dao

import (
	"backend/model"
	"backend/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var (
	err error
	db  *gorm.DB
)

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		// 打印日志
		Logger: logger.Default.LogMode(logger.Info),

		NamingStrategy: schema.NamingStrategy{
			// 禁用自动变复数
			SingularTable: true,
		},
	})

	fmt.Println("dsn:" + dsn)

	if err != nil {
		fmt.Println("数据库连接错误")
	}

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("连接池错误")
		return
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// 迁移 schema
	_ = db.AutoMigrate(
		&model.User{},
		&model.Task{},
		&model.TaskComment{},
		&model.Article{},
		&model.Profile{},
		&model.Message{},
		&model.Letter{},
	)

}
