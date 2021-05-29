package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/clubhouse-clone-biz/config"
	"github.com/jason-shen/clubhouse-clone-biz/ent"
	"github.com/jason-shen/clubhouse-clone-biz/ent/migrate"
	"github.com/jason-shen/clubhouse-clone-biz/middleware"
	"github.com/jason-shen/clubhouse-clone-biz/routes"
	"github.com/jason-shen/clubhouse-clone-biz/utils"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	conf := config.New()

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Name, conf.Database.Password))
	if err != nil {
		utils.Fatalf("Database connection failed: ", err)
	}
	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		utils.Fatalf("Migration fail: ", err)
	}

	app := fiber.New()

	middleware.SetMiddleware(app)

	routes.SetupApiV1(app)

	port := "8000"

	addr := flag.String("addr", port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
