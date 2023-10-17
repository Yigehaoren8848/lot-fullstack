package lot

import (
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	"meilian/constants"
	controller "meilian/controllers/base"

	"meilian/models/mqtt"
	"meilian/utils/jwt"
)

type LotController struct {
	controller.BaseController
}

func (c *LotController) ResponseToken() {
	//颁发token令牌
	userInfo := make(map[string]interface{})
	userInfo["id"] = "10010"
	userInfo["role"] = "admin"
	userInfo["userName"] = "admin"
	token, err := jwt.GenerateToken(userInfo, constants.TOKEN_EXPIRE_TIME)
	if err != nil {
		c.ResponseJSON(constants.TokenGenerateError, "系统异常", "")
		return
	}
	c.ResponseJSON(constants.SuccessCode, "ok", token)

}
func (c *LotController) ControlEquimpment() {

	params := c.GetString("order")

	token := c.Ctx.Input.Header("token")
	if token == "" {
		c.ResponseJSON(constants.AuthError, "你暂无该权限", "")
		return
	}
	tokenMap, _ := jwt.GetTokenData(token)
	if tokenMap["role"] != "admin" {
		c.ResponseJSON(constants.AuthError, "你暂无该权限", "")
		return
	}
	println(params)
	msgs := mqtt.MqttMessage{
		Msg: "",
	}
	if params == "on" {
		// 在 params 为 "on" 时执行相应的操作
		// 可以在这里添加您的代码逻辑
		msgs.Msg = "on"
	} else if params == "off" {
		// 在 params 为 "off" 时执行相应的操作
		// 可以在这里添加您的代码逻辑
		msgs.Msg = "off"
	} else {
		// 如果 params 不是 "on" 也不是 "off"，执行错误处理
		c.ResponseJSON(0, "params is error!", "")
	}

	msgJson, err := json.Marshal(msgs)
	if err != nil {
		c.ResponseJSON(0, "params is error!", "")
	}
	var _ = c.MQTT.Publish("light", 0, false, msgJson)
	logs.Info("" + "访问了日志接口")
	//2.通知日志
	logs.Notice("这是通知日志")
	//3.警告日志
	logs.Warn("这是警告日志")
	//4.警告日志
	logs.Alert("这是警告日志")
	//5.错误日志
	logs.Error("这是错误日志")
	//6.重要的日志
	logs.Critical("这是重要的日志")
	//7.紧急日志
	logs.Emergency("这是紧急日志")
	//8.Debug日志
	logs.Debug("这是调试日志")

	c.ResponseJSON(2, "ok", msgs)
}
