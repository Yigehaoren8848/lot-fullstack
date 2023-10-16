package takeout

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Takeout struct {
	Id          int       `orm:"column(id);auto"`
	ShopId      int       `orm:"column(shop_id)"`
	TakeoutName string    `orm:"column(takeout_name)"`
	Price       float64   `orm:"column(price)"`
	AddTime     time.Time `orm:"column(add_time);type(datetime)"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime)"`
	Operate     string    `orm:"column(operate)"`
	Operator    string    `orm:"column(operator)"`
	Deleted     int       `orm:"column(deleted)"`
	Qc          int       `orm:"column(qc)"`
	Size        int       `orm:"column(size)"`
	Quality     int       `orm:"column(quality)"`
}

func init() {
	orm.RegisterModel(new(Takeout))
}
