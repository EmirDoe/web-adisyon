package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://bm-admin:ec2UIVgkVz1LVFZU@cluster0.ijsfclu.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

var (
	// Client is the MongoDB client
	Client          *mongo.Client
	Context         context.Context
	CancelFunc      context.CancelFunc
	TableCollection *mongo.Collection
	MenuCollection  *mongo.Collection
	ItemCollection  *mongo.Collection
	OrderHistory    *mongo.Collection
)

func Setup() {
	// Setup the database
	Client, Context, CancelFunc = getConnection(uri)
	getCollections()
	TableCollection.Database().Client().Ping(Context, nil)
	MenuCollection.Database().Client().Ping(Context, nil)
	ItemCollection.Database().Client().Ping(Context, nil)
	OrderHistory.Database().Client().Ping(Context, nil)

}

func getConnection(uri string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Failed to create client: %v", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err.Error())
	}

	return client, ctx, cancel
}

func getCollections() {
	databaseName := "web-adisyon"
	database := Client.Database(databaseName)

	TableCollection = database.Collection("tables")
	MenuCollection = database.Collection("menus")
	ItemCollection = database.Collection("items")
	OrderHistory = database.Collection("order_history")
}
