package main

import (
	"mona/config"
	adminv1 "mona/delivery/rest/admin/v1"
	paymentv1 "mona/delivery/rest/payment/v1"
	"mona/domain/repository"
	"mona/domain/service"
	"mona/infrastructure/repository/sqlx"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	env, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	svc := service.New()
	adminApi := adminv1.New(svc, log)
	paymentApi := paymentv1.New(svc, log)
	_ = adminApi
	_ = paymentApi

	repo := repository.New()
	_ = repo
	
	db, err := sqlx.Connect(env.DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to sqlx db")
	}
	defer db.Close()

	log.Info().Msg("successfully connected to sqlx db")
}
