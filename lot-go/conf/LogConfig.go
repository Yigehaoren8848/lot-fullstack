package conf

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"time"
)

func LogAccessInfo(ip, userAgent, method, url string) {
	accessLog := fmt.Sprintf("%s - %s - %s - %s - %s", time.Now().Format("2006-01-02 15:04:05"), ip, userAgent, method, url)
	logs.Info(accessLog)
}
func InitLogConfig() {
	logs.EnableFuncCallDepth(true)
	//get logConf conf
	logPath := "logs/app.log"
	maxsize := 10 * 1024 * 11024
	maxDays := 7
	//set common logger that saves all logConf
	logConf := fmt.Sprintf(`{"filename":"%s","maxsize":%d,"daily":true,"maxdays":%d,"rotate":true,"level":%d}`, logPath, maxsize, maxDays, logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, logConf)
	logs.SetLogger(logs.AdapterConsole)
	// 设置日志异步输出（可选）
	logs.Async()
}
