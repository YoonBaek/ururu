package token

import (
	"errors"

	"github.com/YoonBaek/ururu-server/key"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var (
	errNotAuthenticated error = errors.New("인증되지 않은 사용자입니다")
)

// Verify 기능을 담당하는 핸들러
// route와 연동해서 사용
func LoginRequired() func(c *fiber.Ctx) error {
	return jwtware.New(
		jwtware.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.JSON(utils.ErrorMessage{errNotAuthenticated.Error()})
			},
			SigningKey:    key.LoadPbKey(),
			SigningMethod: "RS256",
		},
	)
}
