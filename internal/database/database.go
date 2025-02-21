package database

import (
	"context"
	"github.com/KasiditR/netflix-clone-go-gin-API/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
	"time"
)

type Mongo struct {
	client   *mongo.Client
	Database *mongo.Database
	Err      error
}

var (
	_mongo         Mongo
	ShortWaitTime  time.Duration = 2
	MediumWaitTime time.Duration = 5
	LongWaitTime   time.Duration = 10
)

func ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), LongWaitTime*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.LoadConfig().MongoURI)
	_mongo.client, _mongo.Err = mongo.Connect(clientOptions)
	if _mongo.Err != nil {
		panic(_mongo.Err)
	}

	if err := _mongo.client.Ping(ctx, readpref.Primary()); err != nil {
		panic(_mongo.Err)
	}

	_mongo.Database = _mongo.client.Database(config.LoadConfig().MongoDatabase)

	log.Println("Connected to MongoDB")
}

func Get() Mongo {
	if _mongo.client == nil {
		ConnectDatabase()
	}

	return _mongo
}

func UserCollection() *mongo.Collection {
	return _mongo.Database.Collection("users")
}
