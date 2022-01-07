package account

import (
	"errors"
	"fmt"
	"time"

	dataBase "github.com/YoonBaek/ururu-server/database"
	"github.com/YoonBaek/ururu-server/token"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	auth                        = &UserAuth{}
	errAcountNotExists    error = errors.New("가입되지 않은 이메일이에요")
	errWrongPw            error = errors.New("비밀번호를 확인해주세요")
	errEmailAlreadyExists error = errors.New("이미 가입된 이메일입니다")
	errWrongPwRepeat      error = errors.New("비밀번호 확인이 일치하지 않습니다")
)

// [POST] 유저로부터 회원가입을 받습니다. 기존 회원과 중복되어 있는지를 체크하고,
// 비밀번호 확인이 일치하면 User 테이블에 새로운 계정을 생성합니다.
// 비밀번호는 UserAuth 테이블을 통해 따로 관리합니다.
func signup(c *fiber.Ctx) error {
	signUpForm := getUserSignUpModel()
	err := c.BodyParser(signUpForm)
	utils.HandleErr(err)

	if isMember(signUpForm.Email) {
		return c.JSON(&utils.ErrorMessage{errEmailAlreadyExists.Error()})
	}
	if signUpForm.Password != signUpForm.Repeat {
		return c.JSON(&utils.ErrorMessage{errWrongPwRepeat.Error()})
	}

	auth = &UserAuth{
		User: User{
			Email:    signUpForm.Email,
			Nickname: initRandomNick(),
		},
		Password: utils.ToHash(signUpForm.Password),
	}

	dataBase.DB().Create(auth)
	return c.JSON(auth.User)
}

// [POST] 유저로부터 로그인 요청을 받습니다.
// 회원인지 여부와 암호를 확인합니다. 암호 확인 과정은 암호화를 거쳐 이루어집니다.
// 회원 여부가 확인되면 JWT 토큰울 반환합니다.
func login(c *fiber.Ctx) error {
	db := dataBase.DB()
	loginForm := &userLogInModel{}
	c.BodyParser(loginForm)

	if !isMember(loginForm.Email) {
		return c.JSON(&utils.ErrorMessage{errAcountNotExists.Error()})
	}

	auth = &UserAuth{}
	db.Joins("User", db.Where(&User{Email: loginForm.Email})).Find(auth)

	if !pwValidation(auth.Password, loginForm.Password) {
		return c.JSON(&utils.ErrorMessage{errWrongPw.Error()})
	}

	return c.JSON(fiber.Map{
		"token": token.SignJWT(jwt.MapClaims{
			"iss":      "ururu.com",
			"exp":      time.Now().Add(time.Hour * 120).Unix(),
			"nickname": auth.User.Nickname,
		}),
	})
}

// [GET] 유저로부터 로그아웃 요청을 받습니다
// goodbye 메시지를 출력하고 만료된 토큰을 반환합니다.
func logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	return c.JSON(fiber.Map{
		"goodbye": fmt.Sprintf("%s님 안녕히 가세요\n세션을 안전하게 종료했습니다.", user["nickname"]),
		"expiredToken": token.SignJWT(jwt.MapClaims{
			"iss":      "ururu.com",
			"exp":      time.Now().Unix() - 1,
			"nickname": auth.User.Nickname,
		}),
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
