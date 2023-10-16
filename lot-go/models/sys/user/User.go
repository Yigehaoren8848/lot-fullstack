package user

import (
	"time"
)

// User 是user表的数据模型
type User struct {
	Id            int       `orm:"column(id);auto"`
	Openid        string    `orm:"column(openid);size(80);null"`
	UserName      string    `orm:"column(user_name);size(50);null"`
	Email         string    `orm:"column(email);size(120);null"`
	Pwd           string    `orm:"column(pwd);size(50);null"`
	PwdEncryption string    `orm:"column(pwd_encryption);size(129);null"`
	Phone         string    `orm:"column(phone);size(25);null"`
	GoodsAddr     string    `orm:"column(goods_addr);size(150);null"`
	Dl            string    `orm:"column(dl);size(255);null"`
	Location      string    `orm:"column(location);size(50);null"`
	Vip           string    `orm:"column(vip);size(5);default(0)"`
	UpdateTime    time.Time `orm:"column(update_time);type(datetime);null"`
	AddTime       time.Time `orm:"column(add_time);type(datetime);null"`
	Role          string    `orm:"column(role);size(5);null"`
	Operate       string    `orm:"column(operate);size(5);null"`
	Oprator       string    `orm:"column(oprator);size(50);null"`
}
