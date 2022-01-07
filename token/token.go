package token

import (
	"github.com/YoonBaek/ururu-server/key"
	"github.com/YoonBaek/ururu-server/utils"
	"github.com/golang-jwt/jwt/v4"
)

// 유저의 정보를 받아 payload를 작성,
// RS256 기반 서명 후 토큰을 반환
func SignJWT(payload jwt.Claims) string {
	// 추후 기능 확장에 따라 payload를 업데이트
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	t, err := token.SignedString(key.LoadPvKey())
	utils.HandleErr(err)

	return t
}
