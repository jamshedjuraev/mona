package main

import (
	"fmt"
	"mona/config"
	"mona/infrastructure/repository/sqlx"
	"mona/pkg/glog"
)

func main() {
	log := glog.New()

	env, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	fmt.Printf("DB DSN: %s\n", env.DB.DSN)

	db, err := sqlx.Connect(env.DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to sqlx db")
	}
	defer db.Close()

	log.Info().Msg("successfully connected to sqlx db")
}
