//package main
//
//import (
//	_ "meilian/routers"
//	beego "github.com/beego/beego/v2/server/web"
//)
//
//func main() {
//	beego.Run()
//}

package main

import (
	//"meilian/filter"

	//a"meilian/controllers"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"meilian/conf"
	"meilian/db"
	_ "meilian/routers"
)

func init() {
	// 连接数据库
	db.ConnectToDatabase()

}

func main() {
	conf.InitLogConfig()
	// 注册跨域拦截器
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"}, // 允许所有域
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type","token"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

//web.AutoRouter(&controllers.MainController{})

	//web.InsertFilter("*",web.BeforeRouter,filter.AuthMiddleware)
	web.Run()
}



//package main
//
//import (
//	"fmt"
//	"os"
//	"time"
//
//	MQTT "github.com/eclipse/paho.mqtt.golang"
//)
//
//func main() {
//	// MQTT连接参数
//	opts := MQTT.NewClientOptions()
//	opts.AddBroker("tcp://localhost:1883") // MQTT Broker的地址
//	opts.SetClientID("dd")
//
//
//	// 创建MQTT客户端
//	client := MQTT.NewClient(opts)
//	if token := client.Connect(); token.Wait() && token.Error() != nil {
//		fmt.Println(token.Error())
//		os.Exit(1)
//	}
//
//	// 订阅主题
//	topic := "lamp/control"
//	if token := client.Subscribe("lw", 0, func(client MQTT.Client, msg MQTT.Message) {
//		// 处理接收到的消息
//		payload := msg.Payload()
//		fmt.Printf("收到消息: %s\n", payload)
//	}); token.Wait() && token.Error() != nil {
//		fmt.Println(token.Error())
//		os.Exit(1)
//	}
//
//	// 控制灯
//	for {
//		// 向主题发送消息来控制灯的开关
//		message := "on" // 控制灯开启
//		client.Publish(topic, 0, false, message)
//		time.Sleep(2 * time.Second)
//
//		message = "off" // 控制灯关闭
//		client.Publish(topic, 0, false, message)
//		time.Sleep(2 * time.Second)
//	}
//}