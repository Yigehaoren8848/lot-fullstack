package db

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() {

	runmode, _ := config.String("runmode")
	if runmode == "dev" {
		web.LoadAppConfig("ini", "./conf/app.dev.conf")
	} else {
		web.LoadAppConfig("ini", "./conf/app.prod.conf")
	}
	mqtt, _ := web.AppConfig.String("db.port")
	print(mqtt)
	// 获取数据库连接信息
	dbHost, _ := web.AppConfig.String("db.host")
	dbPort, _ := web.AppConfig.String("db.port")
	dbUser, _ := web.AppConfig.String("db.user")
	dbPassword, _ := web.AppConfig.String("db.password")
	dbName, _ := web.AppConfig.String("db.name")

	// 构建数据库连接字符串
	dbConnStr := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8"

	// 注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册数据库连接
	err := orm.RegisterDataBase("default", "mysql", dbConnStr)

	// 记录连接数据库的日志
	if err != nil {
		logs.Error("数据库连接失败：", err)
	} else {
		logs.Info("数据库连接成功")
	}
}
