package account

import (
	"errors"
	"time"

	dataBase "github.com/YoonBaek/ururu-server/database"
	"github.com/YoonBaek/ururu-server/key"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/goombaio/namegenerator"
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

func signup(c *fiber.Ctx) error {
	signUpModel := UserSignUpModel()
	err := c.BodyParser(signUpModel)
	utils.HandleErr(err)
	// 1.
	if isMember(signUpModel.Email) {
		return c.JSON(&errorMessage{errEmailAlreadyExists.Error()})
	}
	// 2.
	if signUpModel.Password != signUpModel.Repeat {
		return c.JSON(&errorMessage{errWrongPwRepeat.Error()})
	}
	// 3.
	auth := &UserAuth{
		User: User{
			Email:    signUpModel.Email,
			Nickname: initRandomNick(),
		},
		Password: utils.ToHash(utils.StrToByte(signUpModel.Password)),
	}

	dataBase.DB().Create(auth)
	return c.JSON(auth.User)
}

func login(c *fiber.Ctx) error {
	db := dataBase.DB()
	loginForm := &userLogInModel{}
	c.BodyParser(loginForm)

	if !isMember(loginForm.Email) {
		return c.JSON(&errorMessage{errAcountNotExists.Error()})
	}

	auth := &UserAuth{}
	db.Joins("User", db.Where(&User{Email: loginForm.Email})).Find(auth)
	if auth.Password != utils.ToHash(utils.StrToByte(loginForm.Password)) {
		return c.JSON(&errorMessage{errWrongPw.Error()})
	}

	expire := time.Now().Add(time.Hour * 120).Unix()

	claims := jwt.MapClaims{
		"email": auth.User.Email,
		"name":  auth.User.Nickname,
		"exp":   expire,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(key.LoadKey())
	utils.HandleErr(err)

	return c.JSON(fiber.Map{"token": t})
}

func isMember(email string) bool {
	db := dataBase.DB()
	tx := db.Find(&User{}, "email = ?", email)
	return tx.RowsAffected > 0
}

func initRandomNick() string {
	seed := time.Now().UTC().UnixNano()
	genName := namegenerator.NewNameGenerator(seed)

	name := genName.Generate()
	return name
}
