package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	Imports "ponderada7/imports"
)

type Sensor struct {
	NH3_ppm int    `bson:"NH3_ppm"`
	CO_ppm  int    `bson:"CO_ppm"`
	NO2_ppm int    `bson:"NO2_ppm"`
	Sensor  string `bson:"sensor"`
}


func TestProcessMessage(t *testing.T) {
	msg := &ckafka.Message{
		
	}

	expectedData := Sensor{
		NH3_ppm: 10,
		CO_ppm:  20,
		NO2_ppm: 30,
		Sensor:  "sensor",
	}

	var data Sensor
	Imports.ProcessMessage(msg)

	assert.Equal(t, expectedData, data, "Received data does not match expected data")
}
