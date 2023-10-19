package controller

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"meilian/models/sys/user"
	"meilian/utils/jwt"
)

type BaseController struct {
	web.Controller
	O    orm.Ormer
	MQTT mqtt.Client // 添加 MQTT 客户端字段
}

var clientMap []map[string]interface{}

func (b *BaseController) Prepare() {
	b.O = orm.NewOrm()
	if b.MQTT == nil {

		// MQTT 服务器地址和端口
		broker := "tcp://192.168.1.169:1883"
		uuidObj, err := uuid.NewUUID()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		_ = uuidObj.String()
		// 创建 MQTT 客户端连接配置
		opts := mqtt.NewClientOptions().AddBroker(broker)
		opts.SetClientID("esp8266_client3")
		//opts.SetWill("r","offline",byte(0),true)
		//opts.SetConnectionLostHandler()
		//opts.SetOnConnectHandler()
		b.MQTT = mqtt.NewClient(opts)

		if token := b.MQTT.Connect(); token.Wait() && token.Error() != nil {

		}

	} else {
	}
}
func (c *BaseController) ResponseJSON(code int, msg string, data interface{}) {
	response := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

// 通过token解析用户信息
func (c *BaseController) GetUser() *user.User {
	userModel := &user.User{}
	token := c.Ctx.Input.Header("token")
	if token == "" {
		token = c.GetString("token")
	}
	if token == "" {

		return nil
	}
	userClams, _ := jwt.ParseToken(token)
	id := userClams["id"]
	_ = c.O.QueryTable(new(user.User)).Filter("id", id).One(userModel)
	return userModel
}

// 查询有topic用于程序启动时的订阅
func (c *BaseController) getAllTopics() []orm.Params {
	sql := " SELECT eq.topic FROM equip WHERE eq.delete = 0"
	var topics []orm.Params
	_, err := c.O.Raw(sql).Values(&topics)
	if err != nil {
		return nil

	}
	return topics
}
