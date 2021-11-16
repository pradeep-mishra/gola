package {{resourceName}}

import (
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gofiber/fiber/v2"
)

func TestRoutes(t *testing.T){	
	app := fiber.New()
	Route(app)
	req := httptest.NewRequest("GET", "/{{resourceNames}}", nil)
	resp, _ := app.Test(req,1)
	assert.Equalf(t, 200, resp.StatusCode, "should get 200 http status") 
}
