package bot

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/yechentide/mhnow-bot/config"
)

func Run() {
	slog.Info("Botを起動しています ...")
	discord, err := discordgo.New(config.GetBotToken())
	if err != nil {
		slog.Error("Botの初期化に失敗しました", "error", err)
		os.Exit(1)
	}

	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onMessageUpdate)
	discord.AddHandler(onMessageDelete)

	discord.AddHandler(onMessageReactionAdd)
	discord.AddHandler(onMessageReactionRemove)
	discord.AddHandler(onMessageReactionRemoveAll)

	discord.AddHandler(onInteractionCreated)

	err = discord.Open()
	if err != nil {
		slog.Error("Botのログインに失敗しました", "error", err)
		os.Exit(1)
	}

	defer func() {
		discord.Close()
		fmt.Println("")
		slog.Info("Botを終了しました")
	}()

	for _, guild := range discord.State.Guilds {
		RegisterCommands(discord, guild.ID)
	}

	slog.Info("========== ========== ========== Monster Hunter Now - Discord Bot ========== ========== ==========")
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-stopBot
}

/*
GuildMemberAdd
GuildMemberUpdate
GuildMemberRemove
GuildEmojisUpdate
UserUpdate
MessageDeleteBulk
*/
