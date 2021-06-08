package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/handlers"
)

func SetupUserRoutes(grp fiber.Router, handlers *handlers.Handler) {
	useRoute := grp.Group("/user")
	useRoute.Post("/register", handlers.UserRegister)
}
