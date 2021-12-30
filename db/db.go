package dataBase

import (
	"fmt"

	"github.com/YoonBaek/ururu-server/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func DB() *gorm.DB {
	return db
}

func InitDataBase() {
	db, err = gorm.Open(sqlite.Open("ururu.db"))
	utils.HandleErr(err)
	fmt.Println("------------------------")
	fmt.Println("database init suceed")
	fmt.Println("------------------------")
}
