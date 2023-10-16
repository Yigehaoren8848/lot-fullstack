package filter

import (
	"encoding/json"
	"github.com/beego/beego/v2/adapter/logs"
	"net/http"
)

// ErrorResponse 结构体用于表示错误响应的 JSON 对象
type ErrorResponse struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// WriteJSONResponse 将给定的数据转换成 JSON 字符串并作为响应发送
func WriteJSONResponse(w http.ResponseWriter, msg string, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := ErrorResponse{
		Msg:  msg,
		Code: code,
		Data: data,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		logs.Error("Error encoding JSON response:", err)
	}
}
