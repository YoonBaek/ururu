package migration

import (
	"fmt"

	"github.com/YoonBaek/ururu-server/account"
	"github.com/YoonBaek/ururu-server/article"
	dataBase "github.com/YoonBaek/ururu-server/database"
)

func MakeMigrations() {
	dataBase.DB().AutoMigrate(
		&article.Article{},
		&account.User{},
		&account.UserAuth{},
	)
	fmt.Println("data migrations done")
}
