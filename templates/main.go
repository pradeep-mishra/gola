package main

import (
	"log"

	"github.com/joho/godotenv"
  "{{projectName}}/global"
  "{{projectName}}/server"
)

func main() {
 
  log.Println("Loading .env file...")
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file", err)
  }

  // connect to mongo
  if global.GetEnv("DATABASE_URL") != ""{
    _,err := global.ConnectToMongo()
    if err != nil {
      log.Fatal(err)
    }
  }

  // starting the webserver
  server.StartServer()

}

