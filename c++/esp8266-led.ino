#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <ArduinoJson.h>
// WiFi网络信息
const char* ssid = "TP-LINK_6357";
const char* password = "shitou.123";

// MQTT服务器信息
const char* mqttServer = "192.168.1.169";
const int mqttPort = 1883;
const char* mqttUser = "";      // 如果需要认证，请填写用户名
const char* mqttPassword = "";  // 如果需要认证，请填写密码

// MQTT主题
const char* topic = "light";


// 创建WiFi客户端和MQTT客户端
WiFiClient espClient;
PubSubClient client(espClient);
//控制继电器开关的引脚
const int led = 0;
const int esp_led = 0;


void sendMsg(String msg, String topicString, bool remain) {
  char topic[topicString.length() + 1];
  strcpy(topic, topicString.c_str());
  client.publish(topic, msg.c_str(), remain);
}


void setup() {
  // 连接WiFi
  Serial.begin(300);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    // Serial.println("Connecting to WiFi...");
  }
  // Serial.println("Connected to WiFi");


  // 连接MQTT服务器
  client.setServer(mqttServer, mqttPort);
  client.setKeepAlive(5);
  client.setCallback(callback);
  pinMode(led, OUTPUT);
  pinMode(LED_BUILTIN, OUTPUT);
  digitalWrite(LED_BUILTIN, LOW);
  //初始化引脚为低电平
  digitalWrite(led, HIGH);
  String willTopicString = "lw";
  String willMsg = "离线";
  char willTopic[willTopicString.length() + 1];
  strcpy(willTopic, willTopicString.c_str());
  if (client.connect("esp8266_client442", willTopic, 0, true, willMsg.c_str())) {
    // Serial.println("Connected to MQTT");
    // 订阅控制主题
    client.subscribe(topic);
    // Serial.print("已订阅：");
    // Serial.print(topic);
    sendMsg("上线", "lw", true);

  } else {
    // Serial.print("Failed, rc=");
    // Serial.print(client.state());
    // Serial.println(" Try again in 5 seconds");
    delay(1000);
  }
}

void loop() {

  if (!client.connected()) {
    reconnect();
  }
  // 在这里执行其他任务

  client.loop();
  delay(20);
}

void reconnect() {
  while (!client.connected()) {
    // Serial.println("Connecting to MQTT...");
    if (client.connect("esp8266_client442")) {
      Serial.println("Connected to MQTT");
      // 订阅控制主题
      client.subscribe(topic);
      // Serial.print("已订阅：");
      // Serial.print(topic);
      sendMsg("下线", "lw", true);

    }
  }
}
void callback(char* topic, byte* payload, unsigned int length) {
  //生产环境不可打印任何内容
  // Serial.print("Message arrived ：");

  // for (int i = 0; i < length; i++) {
  //   Serial.print((char)payload[i]);
  // }


  // 解析控制命令并控制灯
  if (strcmp(topic, "light") == 0) {

    // 创建JSON文档对象
    DynamicJsonDocument doc(1024);  // 适当的大小取决于您的JSON消息大小

    // 解析JSON消息
    DeserializationError error = deserializeJson(doc, payload, length);
    if (error) {
      // Serial.println("Failed to parse JSON");
      return;
    }

    // 提取消息中的字段
    const char* msg = doc["msg"];
    if (strcmp(msg, "on") == 0) {
      digitalWrite(LED_BUILTIN, LOW);
    }
    if (strcmp(msg, "off") == 0) {
      digitalWrite(LED_BUILTIN, HIGH);
    }
  }
}
