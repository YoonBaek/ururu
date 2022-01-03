package utils

import (
	"crypto/sha256"
	"fmt"
	"log"
)

func StrToByte(message string) []byte {
	return []byte(message)
}

func HandleErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func ToHash(toBeHashed string) string {
	bytes := StrToByte(toBeHashed)
	crypted := sha256.Sum256(bytes)
	return fmt.Sprintf("%x", crypted)
}
