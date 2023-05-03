package main

import (
	"context"
	"fmt"
	"log"

	"github.com/himanshu1221/Learning-GoLang/25mongoapi/model"
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

//MONGODB HELPERS

//INSERT ONE RECORD

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Movie in DB with ide :", inserted.InsertedID)
}
