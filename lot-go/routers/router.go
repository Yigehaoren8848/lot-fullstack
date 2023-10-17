package routers

import (
	"github.com/beego/beego/v2/server/web"
	"meilian/controllers/lot"
)

func init() {

	web.Router("/ equipments/control", &lot.LotController{}, "GET:ControlEquimpment")
	web.Router("/token", &lot.LotController{}, "GET:ResponseToken")
}
