package bootstrap

import (
	"todolist-go/config"
	"todolist-go/internal/infra/postgresql"
	jwt2 "todolist-go/internal/pkg/jwt"

	"github.com/joho/godotenv"
)

func Start() error {
	_ = godotenv.Load()

	config.NewConfig()
	db, err := postgresql.New()
	if err != nil {
		panic(err)
	}

	postgresql.Migrate(db)

	jwt := jwt2.NewJWT()

	return error(nil)
}
