package dao

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/yechentide/mhnow-bot/config"
)

var db *sql.DB = nil

func openDB() {
	slog.Info("Connecting to database ...")
	var err error
	db, err = sql.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Info("Connected!")
}

func GetDb() *sql.DB {
	if db == nil {
		openDB()
	}
	return db
}

func GetCtx() context.Context {
	return context.Background()
}
