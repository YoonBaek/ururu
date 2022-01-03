package account

import (
	"errors"
	"fmt"

	dataBase "github.com/YoonBaek/ururu-server/database"
	"github.com/YoonBaek/ururu-server/key"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
)

var (
	errAcountNotExists    error = errors.New("가입되지 않은 이메일이에요")
	errWrongPw            error = errors.New("비밀번호를 확인해주세요")
	errEmailAlreadyExists error = errors.New("이미 가입된 이메일입니다")
	errWrongPwRepeat      error = errors.New("비밀번호 확인이 일치하지 않습니다")
)

type errorMessage struct {
	ErrorMessage string `json:"error_message"`
}

// [POST] 유저로부터 회원가입을 받습니다. 기존 회원과 중복되어 있는지를 체크하고,
// 비밀번호 확인이 일치하면 User 테이블에 새로운 계정을 생성합니다.
// 비밀번호는 UserAuth 테이블을 통해 따로 관리합니다.
func signup(c *fiber.Ctx) error {
	signUpForm := getUserSignUpModel()
	err := c.BodyParser(signUpForm)
	utils.HandleErr(err)

	if isMember(signUpForm.Email) {
		return c.JSON(&errorMessage{errEmailAlreadyExists.Error()})
	}
	if signUpForm.Password != signUpForm.Repeat {
		return c.JSON(&errorMessage{errWrongPwRepeat.Error()})
	}

	auth := &UserAuth{
		User: User{
			Email:    signUpForm.Email,
			Nickname: initRandomNick(),
		},
		Password: utils.ToHash(signUpForm.Password),
	}

	dataBase.DB().Create(auth)
	return c.JSON(auth.User)
}

// [POST] 유저로부터 로그인 신청을 받습니다.
// 회원인지 여부와 암호를 확인합니다. 암호 확인 과정은 암호화를 거쳐 이루어집니다.
// 회원 여부가 확인되면 JWT 토큰울 반환합니다.
func login(c *fiber.Ctx) error {
	db := dataBase.DB()
	loginForm := &userLogInModel{}
	c.BodyParser(loginForm)

	if !isMember(loginForm.Email) {
		return c.JSON(&errorMessage{errAcountNotExists.Error()})
	}

	auth := &UserAuth{}
	db.Joins("User", db.Where(&User{Email: loginForm.Email})).Find(auth)

	if !pwValidation(auth.Password, loginForm.Password) {
		return c.JSON(&errorMessage{errWrongPw.Error()})
	}

	return c.JSON(fiber.Map{
		"greetings": fmt.Sprintf("%s님 안녕하세요", auth.User.Nickname),
		"token":     key.GetJWT(auth.User.Email, auth.User.Nickname),
	})
}

func isMember(email string) bool {
	db := dataBase.DB()
	tx := db.Find(&User{}, "email = ?", email)
	return tx.RowsAffected > 0
}

func pwValidation(answer string, query string) bool {
	return answer == utils.ToHash(query)
}
