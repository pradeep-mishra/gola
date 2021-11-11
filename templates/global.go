package global

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var IsMongoConnected = false
var DB *mongo.Database
var Validate *validator.Validate


func init(){
  Validate = validator.New()
}


func GetEnv(key string) string {
  return os.Getenv(key)
}

func SetEnv(key string, value string) error {
  return os.Setenv(key, value)
}

func LookupEnv(key string, defaultValue string) string {
  val, ok := os.LookupEnv(key)
  if !ok {
    return defaultValue
  }
  return val
}


func ConnectToMongo() (*mongo.Client, error) {
  log.Println("Connecting to database...")
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
  log.Println("Database connected")
  return client, nil
}


