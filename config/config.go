package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     uint16 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"db_name"`
	SSLMode  string `toml:"ssl_mode"`
}

type DiscordConfig struct {
	BotToken string `toml:"bot_token"`
}

type AppConfig struct {
	DB      DatabaseConfig
	Discord DiscordConfig
}

const (
	xdgConfigName  = "config.toml"
	homeConfigName = ".mhnow.toml"
)

var config *AppConfig = &AppConfig{
	DB: DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "mhnow",
		SSLMode:  "disable",
	},
	Discord: DiscordConfig{
		BotToken: "Your Bot Token",
	},
}

func loadConfigFromToml(filePath string) {
	slog.Info("Load config from " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		slog.Warn("Failed to load config", "error", err)
		generateConfigFile()
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		slog.Warn("Failed to load config", "error", err)
		generateConfigFile()
	}
	if !strings.HasPrefix(config.Discord.BotToken, "Bot") {
		config.Discord.BotToken = "Bot " + config.Discord.BotToken
	}
	slog.Info("Config loaded.", "config", config)
}

func generateConfigFile() {
	configFilePath := ""
	if os.Getenv("XDG_CONFIG_HOME") == "" {
		configFilePath = os.Getenv("HOME") + "/" + homeConfigName
	} else {
		baseDir := os.Getenv("XDG_CONFIG_HOME") + "/mhnow/"
		if _, err := os.Stat(baseDir); os.IsNotExist(err) {
			_ = os.Mkdir(baseDir, os.ModePerm)
		}
		configFilePath = baseDir + xdgConfigName
	}

	// 1. backup
	_, err := os.Stat(configFilePath)
	if err == nil {
		_ = os.Rename(configFilePath, configFilePath+".bak")
		slog.Info("Backup config file in " + configFilePath + ".bak")
	}

	// 2. generate default config file
	b, err := toml.Marshal(config)
	if err != nil {
		slog.Error("Failed to generate default config file", "error", err)
		return
	}
	_ = os.WriteFile(configFilePath, b, 0644)
	slog.Info("デフォルトの設定ファイルを生成しました: " + configFilePath)
	slog.Info("Botのトークンを設定してください。")
}

func LoadConfig() {
	var err error = nil
	configFile01 := os.Getenv("XDG_CONFIG_HOME") + "/mhnow/" + xdgConfigName
	configFile02 := os.Getenv("HOME") + "/" + homeConfigName
	configFile03 := "./" + homeConfigName

	_, err = os.Stat(configFile01)
	if err == nil {
		loadConfigFromToml(configFile01)
		return
	}

	_, err = os.Stat(configFile02)
	if err == nil {
		loadConfigFromToml(configFile02)
		return
	}

	_, err = os.Stat(configFile03)
	if err == nil {
		loadConfigFromToml(configFile03)
		return
	}

	generateConfigFile()
}

// func GetConfig() *AppConfig {
// 	return config
// }

func GetDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.DBName,
		config.DB.SSLMode,
	)
}

func GetBotToken() string {
	return config.Discord.BotToken
}
