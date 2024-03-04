package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	Sensor "goHive/SensorData"
	"os"
	"log"
	"testing"
	"time"
	"encoding/json"
)

const (
	ClientID     = "publisher"
	MQTPTopic    = "/pond-2"
)

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}


func ConfigureMQTTClient() MQTT.Client {
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
	return client
}

func validateFields(t *testing.T, msg map[string]int, expectedFields []string) {
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("Expected field: %s", field)
			return
		}
	}
}

func TestConnection(t *testing.T) {
	client := ConfigureMQTTClient()
	defer client.Disconnect(250)

	t.Log("Connection with broker MQTT succeeded")
}

func TestDataValidation(t *testing.T) {
	msg := Sensor.SensorData()
	expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
	validateFields(t, msg, expectedFields)
	t.Log("Data validation successful")
}

func TestPublisher(t *testing.T) {
	client := ConfigureMQTTClient()
	defer client.Disconnect(250)

	received := make(chan bool)

	token := client.Subscribe(MQTPTopic, 1, func(client MQTT.Client, msg MQTT.Message) {

		var data map[string]int
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Error validating message: %v", err)
			return
		}

		expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
		validateFields(t, data, expectedFields)

		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to subscribe MQTT topic: %v", token.Error())
	}

	msg := Sensor.SensorData()
	jsonData, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Error converting to JSON: %v", err)
	}

	token = client.Publish(MQTPTopic, 0, false, string(jsonData))
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to publish message: %v", token.Error())
	}

	select {
	case <-received:
		t.Log("Message received")
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout")
	}
	
}

func Tests(t *testing.T) {
	t.Run("TestConnection", TestConnection)
    t.Run("TestDataValidation", TestDataValidation)
    t.Run("TestPublisher", TestPublisher)
}