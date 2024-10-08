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

var (
	// Client is the MongoDB client
	Client                 *mongo.Client
	Context                context.Context
	CancelFunc             context.CancelFunc
	UserCollection         *mongo.Collection
	TableCollection        *mongo.Collection
	MenuCollection         *mongo.Collection
	ItemCollection         *mongo.Collection
	ItemCategoryCollection *mongo.Collection
	OrderCollection        *mongo.Collection
)

func Setup() {
	// Get the URI from the environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	uri := os.Getenv("MONGO_URI")

	// Setup the database
	Client, Context, CancelFunc = getConnection(uri)
	getCollections()

	UserCollection.Database().Client().Ping(Context, nil)
	TableCollection.Database().Client().Ping(Context, nil)
	MenuCollection.Database().Client().Ping(Context, nil)
	ItemCollection.Database().Client().Ping(Context, nil)
	ItemCategoryCollection.Database().Client().Ping(Context, nil)
	OrderCollection.Database().Client().Ping(Context, nil)

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

	databaseName := os.Getenv("MONGO_DB_NAME")
	database := Client.Database(databaseName)

	UserCollection = database.Collection("users")
	TableCollection = database.Collection("tables")
	MenuCollection = database.Collection("menus")
	ItemCollection = database.Collection("items")
	ItemCategoryCollection = database.Collection("item_categories")
	OrderCollection = database.Collection("order_history")
}
