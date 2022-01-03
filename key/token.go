package key

import (
	"time"

	"github.com/YoonBaek/ururu-server/utils"
	"github.com/golang-jwt/jwt/v4"
)

func GetJWT(email, name string) string {
	expire := time.Now().Add(time.Hour * 120).Unix()
	claims := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   expire,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	t, err := token.SignedString(LoadKey())
	utils.HandleErr(err)

	return t
}
