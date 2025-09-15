package postgresql

import (
	"todolist-go/internal/domain/entity"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Task{},
	)

	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
