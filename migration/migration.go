package migration

import (
	"fmt"

	"github.com/YoonBaek/ururu-server/article"
	dataBase "github.com/YoonBaek/ururu-server/db"
)

func MakeMigrations() {
	dataBase.DB().AutoMigrate(
		&article.Article{},
	)
	fmt.Println("data migrations done")
}
