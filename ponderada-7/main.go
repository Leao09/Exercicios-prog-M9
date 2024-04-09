package main

import (
	"context"
	"fmt"
	"log"
	"os"
	api "ponderada7/API"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    mongoUser := os.Getenv("MONGO_USER")
    mongoPassword := os.Getenv("MONGO_PASSWORD")

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.a2vstab.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", mongoUser, mongoPassword)).
        SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        fmt.Println(err)
    }
    defer func() {
        
        if err = client.Disconnect(context.TODO()); err != nil {
            fmt.Println("Iniciando o serviço...")
            fmt.Println(err)
        }
    }()
    
    // for {
    //     api.InitMongoDb(client)
    //     api.ConsumerKafka()
    //     log.Println("Publicado:")
	// 	time.Sleep(2 * time.Second)
    // }
    api.InitMongoDb(client)
    api.ConsumerKafka()
	
	}