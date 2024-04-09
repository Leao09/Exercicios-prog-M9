package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var collection *mongo.Collection

func InitMongoDb(client *mongo.Client) {
	collection = client.Database("test").Collection("sensor")
}

type Sensor struct {
	NH3_ppm int    `bson:"NH3_ppm"`
	CO_ppm  int    `bson:"CO_ppm"`
	NO2_ppm int    `bson:"NO2_ppm"`
	Sensor  string `bson:"sensor"`
}

func Insert(data Sensor) {
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateKafkaConsumer(configMap *ckafka.ConfigMap, topics []string, msgChan chan *ckafka.Message) (*ckafka.Consumer, error) {
	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		for msg := range msgChan {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			var result map[string]interface{}
			err := json.Unmarshal(msg.Value, &result)
			if err != nil {
				log.Fatal(err)
			}

			name := result["sensor"].(string)
			nh3, _ := strconv.Atoi(result["NH3_ppm"].(string))
			co, _ := strconv.Atoi(result["CO_ppm"].(string))
			no2, _ := strconv.Atoi(result["NO2_ppm"].(string))

			data := Sensor{
				NH3_ppm: nh3,
				CO_ppm:  co,
				NO2_ppm: no2,
				Sensor:  name,
			}
			Insert(data)
			fmt.Println("Recebido", msg)
		}
	}()

	fmt.Println("Kafka consumer has been started")

	return consumer, nil
}

func ConsumerKafka() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}
	topics := []string{"ponderada"}
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
		"security.protocol":  "SASL_SSL",
		"sasl.mechanisms":    "PLAIN",
		"group.id":           os.Getenv("CLUSTER_ID"),
		"auto.offset.reset":  "latest",
	}

	msgChan := make(chan *ckafka.Message)
	consumer, err := CreateKafkaConsumer(configMap, topics, msgChan)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	fmt.Println("Kafka consummer has been consumer")
	defer consumer.Close()
}