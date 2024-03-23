package main

import (
	"context"
	"fmt"
	"log"
	"os"
	api "pond6/API"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Carregar variáveis de ambiente do arquivo .env
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    // Recuperar usuário e senha do arquivo .env
    mongoUser := os.Getenv("MONGO_USER")
    mongoPassword := os.Getenv("MONGO_PASSWORD")

    // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.a2vstab.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", mongoUser, mongoPassword)).
        SetServerAPIOptions(serverAPI)

    // Create a new client and connect to the server
    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

    // Inicializar o MongoDB no pacote api
    api.InitMongoDb(client)
	api.Subscriber()

	}