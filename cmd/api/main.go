package main

import (
	"fmt"
	"log"

	"github.com/pavel97go/gosuslugi/config"
	app "github.com/pavel97go/gosuslugi/internal/app"
	pgstore "github.com/pavel97go/gosuslugi/internal/storage/postgres"
)

func main() {
	// 1) Загружаем конфиг
	cfg, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}

	// 2) Подключаемся к БД
	pool, err := pgstore.NewPool(&cfg.DB)
	if err != nil {
		log.Fatal("db connect error: ", err)
	}
	defer pool.Close()
	log.Println("DB connected OK")

	// 3) Роутер
	router := app.NewRouter(pool)

	// 4) Старт сервера
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Println("Listening on", addr)
	if err := router.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
