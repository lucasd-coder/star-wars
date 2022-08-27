package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/lucasd-coder/star-wars/config"
	"github.com/lucasd-coder/star-wars/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func SetUpMongoDB(cfg *config.Config) {
	url := fmt.Sprintf("mongodb://%s:%s", cfg.MongoDbHost, cfg.MongoDbPort)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		logger.Log.Fatal(err.Error())
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		logger.Log.Fatalf("Error MongoDB connection: ", err.Error())
	} else {
		logger.Log.Infoln("MongoDB Connected ")
	}

	client = mongoClient
}

func GetClientMongoDB() *mongo.Client {
	return client
}

func CloseConnMongoDB() error {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
