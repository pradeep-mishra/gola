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
	App = fiber.New()
	
	// Routing here

	App.Listen(":" + global.GetEnv("PORT"))
	appStarted = true
	return App
}