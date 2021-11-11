// Please dont edit this file

package server

import (
	"github.com/gofiber/fiber/v2"
	"{{projectName}}/global"
)

var App *fiber.App
var appStarted = false


func StartServer() *fiber.App {
	if App != nil && appStarted{
		return App
	}
	app := fiber.New()
	App = app
	
	// Routing here

	app.Listen(":" + global.GetEnv("PORT"))
	appStarted = true
	return app
}