package db

import (
	"github.com/YoonBaek/ururu-server/article"
	"github.com/YoonBaek/ururu-server/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func DB(whichDB string) *gorm.DB {
	if db == nil {
		db, err = gorm.Open(sqlite.Open(whichDB))
		utils.HandleErr(err)
	}
	db.AutoMigrate(&article.Article{})
	return db
}
