package main

import (
	"log"

	"github.com/joho/godotenv"
  "{{projectName}}/global"
  "{{projectName}}/server"
  "{{projectName}}/db"
)

func main() {
 
  log.Println("Loading .env file...")
  err := godotenv.Load()
  if err != nil {
    panic("Error loading .env file" + err.Error())
  }

  // connect to mongo
  if global.GetEnv("DATABASE_URL") != ""{
  client,ctx, err := db.ConnectToMongo()
    if err != nil {
      panic(err.Error())
    }
    defer client.Disconnect(*ctx)
  }

  // starting the web server
  server.StartServer()

}

