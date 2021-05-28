package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/middleware"
	"github.com/jason-shen/clubhouse-clone-biz/routes"
	"log"
)

func main()  {
	app := fiber.New()

	middleware.SetMiddleware(app)

	routes.SetupApiV1(app)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
