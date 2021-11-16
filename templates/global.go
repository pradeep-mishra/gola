package global

import (
	"os"
	"strings"

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

type ErrorItem struct {
  Error string `json:"error"`
  Validation string `json:"validation"`
  Value interface{} `json:"value"`
}

func FormatValidationError(err error) []*ErrorItem {
  var errors []*ErrorItem
  if err != nil {
    for _, err := range err.(validator.ValidationErrors){
      var item ErrorItem
      item.Error = err.StructField() + " field is invalid"
      item.Validation = err.Tag()
      item.Value = err.Value()
      errors = append(errors, &item)
    }
  }
  return errors
}

func FormatJSONParserError(err error) []*ErrorItem {
  var errors []*ErrorItem
  var errMessage string
  er := err.Error()
  errPart := ""
  errs := strings.Split(er, ".")
  if len(errs) >= 2{  
    errPart = errs[len(errs)-1]
    errMessage = "Required field, " + errPart
  }else{
    errMessage = er
  }

  var item ErrorItem
  item.Error = errMessage
  item.Validation = errPart
  item.Value = ""
  errors = append(errors, &item)
  return errors
}



