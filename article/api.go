package article

import (
	dataBase "github.com/YoonBaek/ururu-server/db"

	"github.com/gofiber/fiber/v2"
)

// test
// 1. 인스턴스 생성
// 2. 하드코딩 제거. ctx에서 body data 불러와서 article로 파싱
// 3. post to DB
func create(c *fiber.Ctx) error {
	db := dataBase.DB()
	article := &Article{}
	c.BodyParser(article)
	db.Create(article)
	return nil
}

// GET
func read(c *fiber.Ctx) error {
	// id, _ := strconv.Atoi(c.Params("post_no"))
	postNo := c.Params("post_no")
	db := dataBase.DB()
	article := &Article{}
	db.Find(article, "post_no = ?", postNo)
	err := c.JSON(*article)
	return err
}

// PUT
func update(c *fiber.Ctx) error {
	postNo := c.Params("post_no")
	db := dataBase.DB()
	article := &Article{}
	db.Find(article, "post_no = ?", postNo)
	c.BodyParser(article)
	db.Save(article)
	c.JSON(article)
	return nil
}

// DELETE
func delete(c *fiber.Ctx) error {
	db := dataBase.DB()
	db.Delete(&Article{}, c.Params("post_no"))
	return nil
}
