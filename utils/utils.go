package utils

import "log"

func ToByte(message string) []byte {
	return []byte(message)
}

func HandleErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
