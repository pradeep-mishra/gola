package global

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var IsMongoConnected = false
var DB *mongo.Database

func ConnectToMongo() (*mongo.Client, error) {
  client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
  if err != nil {
      return nil, err
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
      return nil, err
  }
  IsMongoConnected = true
  MongoClient = client
  if os.Getenv("DATABASE_NAME") != "" {
    DB = client.Database(os.Getenv("DATABASE_NAME"))
  }
  return client, nil
}


