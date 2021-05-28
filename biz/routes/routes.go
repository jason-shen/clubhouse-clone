package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SetupApiV1(app *fiber.App)  {
	v1 := app.Group("/api/v1")
	Route := v1.Group("/test")
	Route.Get("/hello", func(ctx *fiber.Ctx) error {
		fmt.Println("its all working")

		return nil
	})
}