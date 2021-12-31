package main

import (
	"github.com/YoonBaek/ururu-server/account"
	"github.com/YoonBaek/ururu-server/article"
	dataBase "github.com/YoonBaek/ururu-server/db"
	"github.com/YoonBaek/ururu-server/migration"
	"github.com/gofiber/fiber/v2"
)

func init() {
	dataBase.InitDataBase()
	migration.MakeMigrations()
}

func main() {
	app := fiber.New()
	article.Routes(app)
	account.Routes(app)
	app.Listen(":3000")
}
