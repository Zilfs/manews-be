package seeds

import (
	"manews/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("admin123"), 14)
	if err != nil {
		log.Fatal().Err(err).Msg("Error cereating password hash")
	}

	admin := model.User{
		Name: "Admin",
		Email: "admin@mail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email:"admin@mail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin user")
	}else{
		log.Info().Msg("Admin user seeded successfully")
	}
}