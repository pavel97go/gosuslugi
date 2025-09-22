package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pavel97go/gosuslugi/config"
	app "github.com/pavel97go/gosuslugi/internal/app"
	pgstore "github.com/pavel97go/gosuslugi/internal/storage/postgres"
)

func main() {

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}).
		With().
		Timestamp().
		Logger()

	cfg, err := config.Load("")
	if err != nil {
		log.Fatal().Err(err).Msg("load config")
	}

	pool, err := pgstore.NewPool(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("db connect")
	}
	defer pool.Close()

	r := app.NewRouter(pool)
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info().Str("addr", addr).Msg("listening")
	if err := r.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}
