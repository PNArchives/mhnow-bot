package migrations

import (
	"embed"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/yechentide/mhnow-bot/dao"
)

//go:embed *.sql
var embedMigrations embed.FS

func Setup() {
	db := dao.GetDb()

	slog.Info("Running migrations ...")
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("Failed to set dialect", "error", err)
		os.Exit(1)
	}
	if err := goose.Up(db, "."); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}
}
