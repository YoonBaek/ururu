package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"os"

	"github.com/YoonBaek/ururu-server/utils"
)

const keyFileName string = "secret.key"

var (
	serverKey *rsa.PrivateKey
)

func GenerateKey() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	utils.HandleErr(err)
	saveKey(key)
}

func saveKey(key *rsa.PrivateKey) {
	bytes := x509.MarshalPKCS1PrivateKey(key)
	err := os.WriteFile(keyFileName, bytes, 0644)
	utils.HandleErr(err)
}

func LoadKey() *rsa.PrivateKey {
	keyAsBytes, err := os.ReadFile(keyFileName)
	utils.HandleErr(err)
	serverKey, err = x509.ParsePKCS1PrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return serverKey
}
