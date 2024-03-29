package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MqttCmsBi : Khoi tao doi tuong mqtt
var MqttCmsBi mqtt.Client

// CmsHostBi : host MQTT
const CmsHostBi string = "tcp://test.mosquitto.org:1883"

// const CmsHostBi string = "tcp://192.168.2.9:1883"

// CmsAccessTokenBi : User
const CmsAccessTokenBi = ""

// CmsPassBi : Password
const CmsPassBi = ""

// idBox : Phân định device
var mac = ""

// CmsTopicIn : Server to Box
// var CmsTopicIn = "/v1/devices/monitor/" + mac + "/in"

// CmsTopicOut : Box to Server
// var CmsTopicOut = "/v1/devices/monitor/" + mac + "/out/+"
const CmsTopicOut = "/v1/devices/monitor/out/+"

//PublishData : Function
func PublishData(mac string, payload string) { // idBox : Mac of device
	CmsTopicIn := "/v1/devices/monitor/" + mac + "/in"
	fmt.Println("Mac: " + mac)
	fmt.Println("TOPIC IN :" + CmsTopicIn)
	// If Test device with static Topic
	//CmsTopicIn = "TNQ_MQTT"
	fmt.Println("Test device with static Topic")
	fmt.Println("TOPIC IN :" + CmsTopicIn)
	//var payload string = "{" + "\"ip_private\":" + "\"" + ip + "\"" +"," + "\"box_id\":" + "\"" + id_cam + "\"" + "}"

	fmt.Printf("Payload = %v\n\n", payload)
	Token1 := MqttCmsBi.Publish(CmsTopicIn, 1, false, payload)
	if Token1.Wait() && Token1.Error() != nil {
		fmt.Printf("Error Publish message : %v\n", Token1.Error())
	} else {
		fmt.Println("Send message done")
	}

	// payloadTest := "{" + "\"status\":" + "1" + "}"
	// fmt.Printf("Payload = %v\n\n", payloadTest)
	// time.Sleep(5 * time.Second)
	// Token1 = MqttCmsBi.Publish(CmsTopicOut, 1, false, payloadTest)
	// if Token1.Wait() && Token1.Error() != nil {
	// 	fmt.Printf("Error Publish message : %v\n", Token1.Error())
	// } else {
	// 	fmt.Println("Send message test done")
	// }
	fmt.Println("-------------------------------------------------------------")
}

// MqttBegin : Khoi tao MQTT
func MqttBegin() {

	OptsCmsBI := mqtt.NewClientOptions()
	OptsCmsBI.AddBroker(CmsHostBi)
	OptsCmsBI.SetUsername(CmsAccessTokenBi)
	OptsCmsBI.SetPassword(CmsPassBi)
	OptsCmsBI.SetCleanSession(true)
	OptsCmsBI.SetConnectionLostHandler(MQTTLostConnectHandler)
	OptsCmsBI.SetOnConnectHandler(MQTTOnConnectHandler)

	MqttCmsBi = mqtt.NewClient(OptsCmsBI)
	if Token1 := MqttCmsBi.Connect(); Token1.Wait() && Token1.Error() == nil {
		fmt.Println("MQTT CMS  Connected")
		MqttCmsBi.Subscribe(CmsTopicOut, 0, MqttMessageHandler)
	} else {
		fmt.Println("MQTT CMS  cant not Connected")
		fmt.Printf("Loi CMS  : %v \n", Token1.Error())
		fmt.Println("-------------------")
	}
}

// number_repub : check number push msg faild
var number_repub = 0

//MQTTLostConnectHandler: Check lost connect mqtt server - can't publish msg
func MQTTLostConnectHandler(c mqtt.Client, err error) {
	// c.Disconnect(10)
	// MqttStt = false
	number_repub = number_repub + 1
	fmt.Println("MQTT CMS  Lost Connect")
	fmt.Println("Number reconnect publish msg: " + strconv.Itoa(number_repub))
	fmt.Println(err)
	fmt.Println("------------------------------------------------------------------")
}

var number_resub = 0

// Check lost connect mqtt server - can't subcriber msg from cms
func MQTTOnConnectHandler(client mqtt.Client) {
	number_resub = number_resub + 1
	fmt.Println("Lostconnect chanel subcriber: MQTT_OnConnectHandler")
	fmt.Println("Number reconnect subcriber msg: " + strconv.Itoa(number_resub))
	// client.Unsubscribe(CmsTopicOut)
	// time.Sleep(10)
	//
	if token := client.Subscribe(CmsTopicOut, 0, MqttMessageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		fmt.Println("Subscriber Error MQTT")

	} else {
		fmt.Println("Subcriber is MQTTOnConnectHandler () OKIE ")
	}
}
func MqttMessageHandler(MqttBI mqtt.Client, message mqtt.Message) {
	fmt.Println("=================== MqttMessageHandler ====================")
	fmt.Printf("TOPIC: %s\n", message.Topic())
	fmt.Printf("MSG:\n %s\n", message.Payload())
	fmt.Println("=========================== + = + ==========================")
	dec := json.NewDecoder(bytes.NewReader(message.Payload()))
	var list map[string]interface{}
	if err := dec.Decode(&list); err != nil {
		fmt.Printf("Error:%v\n", err)
		fmt.Println("Message:Loi form message\n")
	} else {
		//***********************************************//
		if list["method"] == "sub_begin" { // Test subcriber
			fmt.Println("================>> Subcriber is OKIE <========== \n")

		} else if list["method"] == "upgrade_engine" {
			fmt.Println(`list["method"] == "upgrade_engine"`)
			s, err := json.Marshal(list["params"])
			if err != nil {
				//fmt.Println(err)
			}
			topic := strings.Replace(message.Topic(), "request", "response", 1)
			msg := `{"method":"upgrade_engine","status":` + strconv.Itoa(1) + `}`
			//file.Println(topic)
			CmsResponse(MqttBI, topic, msg)
			params := string(s)
			fmt.Println(params)
			// UpgradeEngine(params)
		} else {
			//fmt.Println(list)
		}
	}

}

// CmsResponse : Phan hoi message tu box to Server
func CmsResponse(c mqtt.Client, topic string, msg string) {
	c.Publish(topic, 0, false, msg)
}
