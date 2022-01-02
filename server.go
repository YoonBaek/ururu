package main

import (
	"github.com/YoonBaek/ururu-server/account"
	"github.com/YoonBaek/ururu-server/article"
	dataBase "github.com/YoonBaek/ururu-server/database"
	"github.com/YoonBaek/ururu-server/key"
	"github.com/YoonBaek/ururu-server/migration"
	"github.com/gofiber/fiber/v2"
)

func init() {
	key.GenerateKey()
	dataBase.InitDataBase()
	migration.MakeMigrations()
}

func main() {
	app := fiber.New()
	article.Routes(app)
	account.Routes(app)
	app.Listen(":3000")

}
