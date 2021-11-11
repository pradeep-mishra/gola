package ${{resourceName}}

import (
	"os"
	//"${{projectName}}/global"
	"github.com/gofiber/fiber/v2"
)


func Load(app *fiber.App) {
	
	${{resourceName}} := app.Group("/${{resourceNames}}") // grouping of all your routes

	// add all your routes here

	// sample route
	${{resourceName}}.Get("/", Get${{resourceNameCapitalized}})
	
}


