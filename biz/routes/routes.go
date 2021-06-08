package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/handlers"
)

func SetupApiV1(app *fiber.App, handlers *handlers.Handler) {
	v1 := app.Group("/api/v1")
	SetupUserRoutes(v1, handlers)
}
