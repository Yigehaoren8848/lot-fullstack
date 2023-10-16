package wx

import (
	"bytes"
	"encoding/json"
	"fmt"
	controller "meilian/controllers/base"
	"net/http"
)

type PushMsgController struct {
	controller.BaseController
}

func (c *PushMsgController) PushMsg() {
	// 假设您已经获取了 access_token，这里只是示意代码
	access_token := "71_1eNUjHsgHt3PPAxp2ul-1p4Yl5O0Jh3L5LENyVrfN7A60rORTD7F-e5oy8gUrjRNrbW16JIuXBI8v0cHM5kHlRl9Msrd0DiHnyxMm_Mr9Tj6dg-Jtjpt-8KT2OgHMNeABAOTU"

	// 构建消息内容
	data := map[string]map[string]string{
		"thing1": {"value": "2023-08-09"},

		// 其他模板字段
		"thing9": {"value": "测试商品"},
	}

	jsonData := map[string]interface{}{
		"touser":      "o5qWM4rHlcOYiFSl39cvLUMAT2T0",
		"template_id": "pkFiQffPPG7xKOVjFPrCgI1gQz_4CQ8Q1Dm6VhVOREI",
		"data":        data,
	}

	jsonDataBytes, _ := json.Marshal(jsonData)

	// 发送订阅消息
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s", access_token)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonDataBytes))
	if err != nil {
		// 处理错误
		c.ResponseJSON(0, err.Error(), "")
	}

	defer resp.Body.Close()

	// 处理响应
	// ...
	c.ResponseJSON(1, "ok", "")

}
