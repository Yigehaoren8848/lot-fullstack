package lot

import (
	"encoding/json"
	"meilian/constants"
	controller "meilian/controllers/base"
	models "meilian/models/mall"
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

	c.ResponseJSON(2, "ok", msgs)
}

func (c *LotController) close() {
	var courses []models.Courses

	_, _ = c.O.QueryTable("courses").All(&courses, "Id", "Title", "Subtitle", "Time")
	c.ResponseJSON(2, "ok", "")
}
