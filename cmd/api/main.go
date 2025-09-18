package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/pavel97go/gosuslugi/config"
	app "github.com/pavel97go/gosuslugi/internal/app"
	pgstore "github.com/pavel97go/gosuslugi/internal/storage/postgres"
)

func main() {
	// zerolog настройка
	zerolog.TimeFieldFormat = time.RFC3339
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen})

	cfg, err := config.Load("")
	if err != nil {
		zlog.Fatal().Err(err).Msg("load config")
	}

	pool, err := pgstore.NewPool(&cfg.DB)
	if err != nil {
		zlog.Fatal().Err(err).Msg("db connect error")
	}
	defer pool.Close()

	r := app.NewRouter(pool)
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	zlog.Info().Msgf("listening on %s", addr)

	if err := r.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
