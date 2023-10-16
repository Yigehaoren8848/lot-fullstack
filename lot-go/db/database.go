package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	userModel "meilian/models/sys/user"
)

func ConnectToDatabase() {
	deployMethod, _ := config.String("deployMethod")
	var dataSourse = ""
	if deployMethod == "docker" {
		dataSourse = "root:shitou19990225@tcp(mysql:3306)/youxue"
	} else {
		dataSourse = "root:shitou19990225@tcp(localhost:3306)/youxue"
	}

	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 连接数据库
	err := orm.RegisterDataBase("default", "mysql", dataSourse, 30, 100)
	if err != nil {
		print(err.Error())
	}
	orm.RegisterModel(new(userModel.User))
	orm.Debug = true

	// 记录连接数据库的日志
	if err != nil {
		logs.Error("数据库连接失败：", err)
	} else {
		logs.Info("数据库连接成功")
	}
}
