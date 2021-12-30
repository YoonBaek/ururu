package board

import (
	u "github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	const path string = "/board"

	app.Get(path+"/:id", article)
	app.Get(path, articles)

	app.Post(path+"/create", create)
	app.Post(path+"/:id/delete", delete)
	app.Post(path+"/board/:id/update", update)
	app.Post(path+"/board/:id/up", voteUp)
	app.Post(path+"/board/:id/down", voteDown)
}

// POST
func voteDown(c *fiber.Ctx) error {
	id := c.Params("id")
	c.Send(u.ToByte("id: " + id))
	c.Send(u.ToByte("vote down"))
	return nil
}

// POST
func voteUp(c *fiber.Ctx) error {
	id := c.Params("id")
	c.Send(u.ToByte("id: " + id))
	c.Send(u.ToByte("vote up"))
	return nil
}

// POST
// create article
func create(c *fiber.Ctx) error {
	c.Send(u.ToByte("create article"))
	return nil
}

// GET
// View an article
func article(c *fiber.Ctx) error {
	id := c.Params("id")
	c.Send(u.ToByte("id: " + id))
	c.Send(u.ToByte("view article"))
	return nil
}

// GET
// View article list
func articles(c *fiber.Ctx) error {
	c.Send(u.ToByte("view articles"))
	return nil
}

// POST
// delete article
func delete(c *fiber.Ctx) error {
	id := c.Params("id")
	c.Send(u.ToByte("id: " + id))
	c.Send(u.ToByte("delete article"))
	return nil
}

// POST
// update article
func update(c *fiber.Ctx) error {
	id := c.Params("id")
	c.Send(u.ToByte("id: " + id))
	c.Send(u.ToByte("update article"))
	return nil
}
