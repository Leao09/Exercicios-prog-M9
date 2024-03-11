package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

const (
	ClientID     = "subscriber"
	MQTPTopic    = "/pond-2"
)
var db *sql.DB

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("[SUBSCRIBER][%s] %s \n", msg.Topic(), msg.Payload())
	result := strings.Split(string(msg.Payload()), " - ")
	name := result[1]
	nh3, _ := strconv.Atoi(result[2])
	co, _ := strconv.Atoi(result[3])
	no2, _ := strconv.Atoi(result[4])

	data := Sensor{NH3_ppm: nh3, CO_ppm: co, NO2_ppm: no2, sensor: name}
	Insert(data, db)
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

type Sensor struct {
	NH3_ppm, CO_ppm, NO2_ppm int
	sensor                   string
}

func Insert(data Sensor, db *sql.DB) {

	_, err := db.Exec(fmt.Sprintf("INSERT INTO sensor (sensor, NH3_ppm, CO_ppm, NO2_ppm) VALUES ('%v', %v, %v, %v)", data.sensor, data.NH3_ppm, data.CO_ppm, data.NO2_ppm))
	if err != nil {
		log.Fatal(err)
	}
}

func Subscriber(dbPointer *sql.DB) {

	db = dbPointer
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(ClientID)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}

	if token := client.Subscribe(MQTPTopic , 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	select {}
}