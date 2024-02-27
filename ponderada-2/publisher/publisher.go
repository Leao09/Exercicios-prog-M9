package main

import (
	"encoding/json"
	"log"
	"time"
	Sensor "paho-go/SensorData"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTTBroker   = "tcp://localhost:4567"
	ClientID     = "publisher"
	MQTPTopic    = "/pond-2"
)

func ConfigureMQTTClient() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(MQTTBroker)
	opts.SetClientID(ClientID)
	return opts
}

func PublishData(client MQTT.Client, topic string, qos byte, data string) {
	token := client.Publish(topic, qos, false, data)
	token.Wait()
}


func Client() {
	opts := ConfigureMQTTClient()
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Error connecting to MQTT broker:", token.Error())
		return
	}

	for {
		data := Sensor.SensorData()
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error converting data to JSON", err)
			return
		}

		msg := time.Now().Format(time.RFC3339) + " " + string(jsonData)
		PublishData(client, MQTPTopic, 1, msg)

		log.Println("[PUBLISHER] ", msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Client()
}