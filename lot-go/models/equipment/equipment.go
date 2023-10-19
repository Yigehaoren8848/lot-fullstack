package equipment

import (
	"github.com/astaxie/beego/orm"
)

type Equip struct {
	EquipID         int    `orm:"column(equip_id);pk"`
	CategoryID      int    `orm:"column(category_id);null"`
	UserID          int    `orm:"column(user_id);null"`
	EquipName       string `orm:"column(equip_name);null"`
	Topic           string `orm:"column(topic);null"`
	ClientLotID     string `orm:"column(client_lot_id);null"`
	ClientID        string `orm:"column(client_id);null"`
	EquipCategoryId string `orm:"column(equip_category_id);null"`
}

func init() {
	// 注册数据模型
	orm.RegisterModel(new(Equip))
}
