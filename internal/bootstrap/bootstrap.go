package bootstrap

import (
	"todolist-go/config"
	handler2 "todolist-go/internal/handler"
	"todolist-go/internal/infra/postgresql"
	"todolist-go/internal/middleware"
	bcrypt2 "todolist-go/internal/pkg/bcrypt"
	jwt2 "todolist-go/internal/pkg/jwt"
	"todolist-go/internal/repository"
	"todolist-go/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Start() error {
	_ = godotenv.Load()

	app := fiber.New()

	config.NewConfig()
	db, err := postgresql.New()
	if err != nil {
		panic(err)
	}

	postgresql.Migrate(db)

	bcrypt := bcrypt2.NewBcrypt(10)
	jwt := jwt2.NewJWT()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(usecase.UsecaseDependencies{
		Repository: repository,
		Bcrypt:     *bcrypt,
		JWT:        *jwt,
	})

	middleware := middleware.NewMiddleware(jwt, *usecase)

	v1 := app.Group("/api/v1")

	handler := handler2.NewHandler(v1, *usecase, *middleware)
	handler.MountEndpoint()

	return app.Listen(":3000")
}
