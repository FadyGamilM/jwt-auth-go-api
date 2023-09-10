package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	// get env config vars ready
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("error trying to load env variables : %v", err)
	}
	uri := os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// connect
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("error trying to connect to mongodb : %v", err)
	}

	// ping to the client
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("error trying to ping the db instance : %v", err)
	}

	log.Println("connected to mongodb ...")
	return client
}
