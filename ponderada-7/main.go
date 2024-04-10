package main

import (
	"fmt"
	"log"
	"os"
    Imports "ponderada7/imports"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	topics := []string{"ponderada"}
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.username":     os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":     os.Getenv("CONFLUENT_API_SECRET"),
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"group.id":          os.Getenv("CLUSTER_ID"),
		"auto.offset.reset": "latest",
	}

	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topics: %v", err)
	}

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			Imports.ProcessMessage(msg)
		} else {
			log.Printf("Error consuming message: %v\n", err)
		}
	}
}

