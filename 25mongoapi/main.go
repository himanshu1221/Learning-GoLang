package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://himanshuchhatwal9295:himanshu2314@golang.9spcknm.mongodb.net/test"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongo db

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reffrece is ready")
}
