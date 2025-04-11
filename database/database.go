package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Database *mongo.Database

func ConnectDB(uri string, dbname string) *mongo.Database {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to Database: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging mongo client: %v", err)
	}

	Database = client.Database(dbname)

	return Database
}
