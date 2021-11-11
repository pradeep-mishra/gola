package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
  "${{projectName}}/global"
  "${{projectName}}/server"
)

func main() {
  loadEnv()
  dbConnect()
  startApp()
}

func loadEnv() {
  log.Println("Loading .env file...")
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading .env file", err)
  }
}

func dbConnect() {
  if os.Getenv("DATABASE_URL") != ""{
    _,err := global.ConnectToMongo()
    if err != nil {
      log.Fatal(err)
    }
  }
}

func startApp() {
  server.StartServer()
}

