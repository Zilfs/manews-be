package seeds

import (
	"manews/internal/core/domain/model"
	"manews/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error cereating password hash")
	}

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@mail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@mail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin user")
	} else {
		log.Info().Msg("Admin user seeded successfully")
	}
}
