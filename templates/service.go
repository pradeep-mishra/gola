package {{resourceName}}

import (
	//"context"
	//"time"
	"log"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"{{projectName}}/global"
)


var collection *mongo.Collection


func SetCollection(m *mongo.Database){
	log.Println("setting {{resourceNames}} collection")
	collection = m.Collection("{{resourceNames}}")
}

//--------------------------------------------------------------------------------------//
// add all your service functions for resource {{resourceName}} here


func Get{{resourceNameCapitalized}}(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": global.GetEnv("WELCOME_MESSAGE"),
	})
}

// to Get and Set data from Env, from .env file
// global.GetEnv("VAR_NAME")
// global.SetEnv("VAR_NAME")
// global.LookupEnv("VAR_NAME", "Fallback Value")


func New{{resourceNameCapitalized}}(c *fiber.Ctx) error {
	dto := new({{resourceNameCapitalized}}Dto)
	
	if err := c.BodyParser(dto);err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}

	if err:= global.Validate.Struct(dto); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error" : err.Error(), 
		})
	}

	// add object in mongoDB
	/*
	{{resourceName}}Obj := &{{resourceNameCapitalized}}Schema{
		ID: primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name : dto.Name,
		Email : dto.Email,
	} 
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	if  _, err := collection.InsertOne(ctx, {{resourceName}}Obj);err != nil{
		return c.Status(500).JSON(fiber.Map{
			"error" : err.Error(), 
		})
	}
	return c.Status(200).JSON({{resourceName}}Obj)
	*/

	return c.Status(200).JSON(dto)
}
