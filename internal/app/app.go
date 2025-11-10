package app

import (
	"manews/config"

	"github.com/rs/zerolog/log"
)

func RunServer(){
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil{
		log.Fatal().Msgf("Failed to connect to database: %v", err)
		return
	}	
}