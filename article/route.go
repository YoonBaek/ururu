package article

import "github.com/gofiber/fiber/v2"

const appName URL = "/articles"

type URL string

func (u URL) getURL(path string) string {
	return string(u) + path
}

func Routes(app *fiber.App) {
	app.Get(appName.getURL("/:post_no"), read)
	app.Post(appName.getURL("/create"), create)
}
