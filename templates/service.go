package ${{resourceName}}

import (
	//"os"
	//"${{projectName}}/global"
	"github.com/gofiber/fiber/v2"
)

// add all your service functions for resource ${{resourceName}} here

func Get${{resourceNameCapitalized}}(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello from ${{resourceName}}",
	})
	// to get data from .env file
	// os.Getenv("VAR_NAME")
}
