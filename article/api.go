package article

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/", create)
}

// test
// 1. 인스턴스 생성
// 2. value 하드코딩
// 3. post to DB
func create(c *fiber.Ctx) error {
	atc := &Article{}
	atc.Title = "첫 게시글 입니다."
	atc.Content = "첫 게시글 내용 입니다."
	atc.CodeNo = "BTC"

	err := c.JSON(&atc)
	return err
}
