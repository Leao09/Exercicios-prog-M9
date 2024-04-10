package imports

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)


type Sensor struct {
	NH3_ppm int    `bson:"NH3_ppm"`
	CO_ppm  int    `bson:"CO_ppm"`
	NO2_ppm int    `bson:"NO2_ppm"`
	Sensor  string `bson:"sensor"`
}
func ProcessMessage(msg *ckafka.Message) {
	log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

	// Parse the message string
    parts := strings.Split(string(msg.Value), " - ")
    // if len(parts) != 7 {
    //     log.Printf("Invalid message format: %s", string(msg.Value))
    //     return
    // }

    sensor := parts[1]
    nh3, _ := strconv.Atoi(parts[2])
    co, _ := strconv.Atoi(parts[3])
    no2, _ := strconv.Atoi(parts[4])

	data := Sensor{
		NH3_ppm: nh3,
		CO_ppm:  co,
		NO2_ppm: no2,
		Sensor:  sensor,
	}
	sensorData, _ := json.Marshal(data)
	fmt.Println("Received:", sensorData)
	WriteToFile(string(sensorData), "consumer.txt")
} 

func WriteToFile(message string, filename string) error {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(message + "\n")
    if err != nil {
        return err
    }

    fmt.Printf("Mensagem escrita com sucesso no arquivo %s\n", filename)
    return nil
}