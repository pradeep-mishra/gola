package {{resourceName}}

import (
	"{{projectName}}/global"
	"github.com/gofiber/fiber/v2"
)


//var collection *global.DB.Collection

func init(){
	// collection = global.DB.Collection("{{resourceNames}}")
}

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
	dataModel := new({{resourceNameCapitalized}}Dto)
	
	if err := c.BodyParser(dataModel);err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}

	if err:= global.Validate.Struct(dataModel); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error" : err.Error(), 
		})
	}

	return c.Status(200).JSON(dataModel)
}
