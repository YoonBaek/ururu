package article

import (
	dataBase "github.com/YoonBaek/ururu-server/db"

	"github.com/gofiber/fiber/v2"
)

// test
// 1. 인스턴스 생성
// 2. value 하드코딩
// 3. post to DB
func create(c *fiber.Ctx) error {
	db := dataBase.DB()
	atc := &Article{}
	atc.Title = "첫 게시글 입니다."
	atc.Content = "첫 게시글 내용 입니다."
	atc.CodeNo = "BTC"
	db.Create(atc)
	err := c.JSON(&atc)
	return err
}

func read(c *fiber.Ctx) error {
	// id, _ := strconv.Atoi(c.Params("post_no"))
	postNo := c.Params("post_no")
	db := dataBase.DB()
	article := &Article{}
	db.Find(article, "post_no = ?", postNo)
	err := c.JSON(*article)
	return err
}
