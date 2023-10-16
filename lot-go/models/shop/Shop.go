package shop

import "github.com/astaxie/beego/orm"

type Shop struct {
	Id                 int    `orm:"auto"`
	UserId             int    `orm:"column(user_id);null"`
	BgImageId          string `orm:"column(bg_image_id);null"`
	StorefrontId       int    `orm:"column(storefront_id);null"`
	ShopName           string `orm:"column(shop_name);null"`
	BusinessTime       string `orm:"column(business_time);null"`
	ClosingTime        string `orm:"column(closing_time);null"`
	Business           int    `orm:"column(business);null"`
	Address            string `orm:"column(address);null"`
	Lat                string `orm:"column(lat);null"`
	Lng                string `orm:"column(lng);null"`
	Location           string `orm:"column(location)"`
	Distance           string
	Deleted            string `orm:"column(deleted);null"`
	Module             string `orm:"column(module);null"`
	HasBusinessLicense int    `orm:"column(business_license);null"`
}

func init() {
	orm.RegisterModel(new(Shop))
}
