package repository

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Database *mongo.Database

func ConnectDb(dbName string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	Database = client.Database(dbName)
	fmt.Println("Connected to " + Database.Name())
}
