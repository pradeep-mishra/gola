package ${{resourceName}}

import (
	"${{projectName}}/global"
	"github.com/gofiber/fiber/v2"
)

// add all your service functions for resource ${{resourceName}} here

func Get${{resourceNameCapitalized}}(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello from ${{resourceName}}",
	})
}


func GetWelcome(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": global.LookupEnv("WELCOME_MESSAGE", "Hola User")
	})

	// to get/set data from env, data from .env file
	// global.GetEnv("VAR_NAME")
	// global.SetEnv("VAR_NAME")
	// global.LookupEnv("VAR_NAME", "Fallback_Value")
}
