package mqtt

type MqttMessage struct {
	Msg string `json:"msg"`
}

type MqttStatus struct {
	isconnected string
	clientId    string
}
