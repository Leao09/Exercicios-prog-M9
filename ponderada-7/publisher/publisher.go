package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	Sensor "ponderada7/SensorData"
	imports "ponderada7/imports"
	"strconv"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

const (
	ClientID     = "publisher"
	MQTPTopic    = "ponderada"
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
	type Sens struct {
		NH3_ppm int    `bson:"NH3_ppm"`
		CO_ppm  int    `bson:"CO_ppm"`
		NO2_ppm int    `bson:"NO2_ppm"`
		Sensor  string `bson:"sensor"`
	}

	for {
		data := Sensor.SensorData()
		msg := time.Now().Format(time.RFC3339) + " - " + "sensor" + " - " + strconv.Itoa(data["NH3_ppm"]) + " - " + strconv.Itoa(data["CO_ppm"]) + " - " + strconv.Itoa(data["NO2_ppm"])
		PublishData(client, MQTPTopic, 1, msg)
		log.Println("Publicado:", msg)
		parts := strings.Split(string(msg), " - ")
		sensor := parts[1]
		nh3, _ := strconv.Atoi(parts[2])
		co, _ := strconv.Atoi(parts[3])
		no2, _ := strconv.Atoi(parts[4])

		sensorData := Sens{
			NH3_ppm: nh3,
			CO_ppm:  co,
			NO2_ppm: no2,
			Sensor:  sensor,
		}
		sensorJson, _ := json.Marshal(sensorData)
		imports.WriteToFile(string(sensorJson), "publisher.txt")
		time.Sleep(2 * time.Second)
	}

	
}

func main() {
    Client()
}