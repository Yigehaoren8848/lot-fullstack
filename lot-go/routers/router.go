package routers

import (
	"github.com/beego/beego/v2/server/web"
	"meilian/controllers/lot"
	controller "meilian/controllers/sys/user"
)

func init() {

	web.Router("/ equipments/control", &lot.LotController{}, "GET:ControlEquimpment")
	web.Router("/login", &controller.UserController{}, "post:Login")
	web.Router("/token", &lot.LotController{}, "GET:ResponseToken")
}
