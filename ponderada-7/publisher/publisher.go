package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	Sensor "pond6/SensorData"
	"os"
	"log"
	"time"
	"strconv"
)

const (
	ClientID     = "publisher"
	MQTPTopic    = "/pond-2"
)

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}


func PublishData(client MQTT.Client, topic string, qos byte, data string) {
	token := client.Publish(topic, qos, false, data)
	token.Wait()
}


func ConfigureMQTTClient() *MQTT.ClientOptions {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(ClientID)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}

func Client() {
	opts := ConfigureMQTTClient()
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	for {
		data := Sensor.SensorData()

		msg := time.Now().Format(time.RFC3339) + " - " + "sensor" + " - " + strconv.Itoa(data["NH3_ppm"]) + " - " + strconv.Itoa(data["CO_ppm"]) + " - " + strconv.Itoa(data["NO2_ppm"])
		PublishData(client, MQTPTopic, 1, msg)
		log.Println("Publicado:", msg)
		time.Sleep(2 * time.Second)
	}

	
}

func main() {
    Client()
}