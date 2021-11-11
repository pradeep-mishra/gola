package ${{controllerName}}

import (
	"os"
	//"${{projectName}}/global"
	"github.com/gofiber/fiber/v2"
)


func Load(app *fiber.App) {
	
	${{controllerName}} := app.Group("/${{controllerNames}}") // grouping of all your routes

	// add all your routes here

	// sample route
	${{controllerName}}.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(os.Getenv("WELCOME_MESSAGE"))
	})
	
}


