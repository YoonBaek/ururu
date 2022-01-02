package account

import "github.com/gofiber/fiber/v2"

const appName URL = "/account"

type URL string

func (u URL) getURL(path string) string {
	return string(u) + path
}

func Routes(app *fiber.App) {
	app.Post(appName.getURL("/signup"), signup)
	app.Post(appName.getURL("/login"), login)
}
