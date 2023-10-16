package main

import (
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/yechentide/mhnow-bot/bot"
	"github.com/yechentide/mhnow-bot/config"
	"github.com/yechentide/mhnow-bot/dao"
	"github.com/yechentide/mhnow-bot/migrations"
)

func main() {
	config.LoadConfig()

	db := dao.GetDb()
	migrations.Setup()
	defer func() {
		db.Close()
		slog.Info("Disconnected from database.")
	}()

	bot.Run()
}
