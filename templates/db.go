package db

import (
  "context"
  "log"
  "os"
  "time"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "{{projectName}}/global"
)




func ConnectToMongo() (*mongo.Client, *context.Context, error) {
  //log.Println("Connecting to database...")
  client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
  if err != nil {
      return nil, nil,err
  }
  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
      return nil, nil, err
  }
  global.IsMongoConnected = true
  global.MongoClient = client
  if os.Getenv("DATABASE_NAME") != "" {
    db := client.Database(os.Getenv("DATABASE_NAME"))
    global.Database = db
		// dont remove this comment
		// collection set
  }
  //log.Println("Database connected")
  return client, &ctx, nil
}
