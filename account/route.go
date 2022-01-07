package account

import (
	"github.com/YoonBaek/ururu-server/token"
	"github.com/gofiber/fiber/v2"
)

const appName URL = "/account"

var authRequried = token.LoginRequired()

type URL string

func (u URL) getURL(path string) string {
	return string(u) + path
}

func Routes(app *fiber.App) {
	app.Post(appName.getURL("/signup"), signup)
	app.Post(appName.getURL("/login"), login)
	app.Post(
		appName.getURL("/logout"),
		authRequried,
		logout,
	)
}
