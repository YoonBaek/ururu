package account

import (
	dataBase "github.com/YoonBaek/ururu-server/db"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
)

type errorMessage struct {
	ErrorMessage string `json:"error_message"`
}

// 회원 가입 로직.
// 1. 이메일이 중복인지 여부 체크
// 2. 비밀번호 오기를 방지하기 위해 중복 체크
// 3. 비밀번호는 암호화해서 서버에 저장
func signup(c *fiber.Ctx) error {
	db, usim := dataBase.DB(), UserSignInModel()
	err := c.BodyParser(usim)
	utils.HandleErr(err)
	// 1.
	tx := db.Find(&User{}, "email = ?", usim.Email)
	if tx.RowsAffected > 0 {
		return c.JSON(&errorMessage{"email already exists"})
	}
	// 2.
	if usim.Password != usim.Repeat {
		return c.JSON(&errorMessage{"please repeat password correctly"})
	}

	user := &User{}
	user.Email = usim.Email
	// 3.
	user.Password = utils.ToHash(utils.StrToByte(usim.Password))

	db.Create(user)
	return c.JSON(user)
}
