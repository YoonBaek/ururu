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

// 로컬에 이미 키가 존재한다면 매번 사용할 필요 없습니다
// 주기적으로 갱신하는 로직은 추후 구현
func Generate() {
	save := func(key *rsa.PrivateKey) {
		bytes := x509.MarshalPKCS1PrivateKey(key)
		err := os.WriteFile(keyFileName, bytes, 0644)
		utils.HandleErr(err)
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	utils.HandleErr(err)
	save(key)
}

// 키 생성을 위한 내부 비공개 키를 반환
func LoadPvKey() *rsa.PrivateKey {
	if serverKey != nil {
		return serverKey
	}
	keyAsBytes, err := os.ReadFile(keyFileName)
	utils.HandleErr(err)
	serverKey, err = x509.ParsePKCS1PrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return serverKey
}

// 검증을 위한 공개 키를 반환
func LoadPbKey() *rsa.PublicKey {
	return &LoadPvKey().PublicKey
}
