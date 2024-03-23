package api

import (
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)


func SelectFromMongo() {
    // Executar uma busca para obter os documentos do MongoDB
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.Background())

    // Iterar sobre os documentos e imprimir os dados
    for cursor.Next(context.Background()) {
        var data Sensor
        if err := cursor.Decode(&data); err != nil {
            log.Fatal(err)
        }
        log.Printf("Sensor data: %+v\n", data)
    }
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
}

