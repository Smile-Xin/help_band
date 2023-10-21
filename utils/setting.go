package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbHost     string
	DbName     string
)

func Init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件出错")
	}

	LoadData(file)

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("236753")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbName = file.Section("database").Key("DbName").MustString("help_band01")
}
