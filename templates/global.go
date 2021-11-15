package global

import (
	"os"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client
var IsMongoConnected = false
var Database *mongo.Database
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



