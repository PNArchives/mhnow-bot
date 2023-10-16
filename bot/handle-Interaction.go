package bot

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	cmds "github.com/yechentide/mhnow-bot/commands"
)

func onInteractionCreated(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		slog.Debug("Interaction Created", "Command", i)
		commandHandler(s, i)
	case discordgo.InteractionMessageComponent:
		slog.Debug("Interaction Created", "MessageComponent", i)
	case discordgo.InteractionModalSubmit:
		slog.Debug("Interaction Created", "ModalSubmit", i)
	default:
		slog.Debug("Interaction Created", i.Type, i)
	}
}

func commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	hunterName := i.Member.Nick
	if len(hunterName) == 0 {
		hunterName = i.Member.User.Username
	}
	slog.Info(fmt.Sprintf("%s がコマンド %s を実行しました", hunterName, i.ApplicationCommandData().Name))
	switch i.ApplicationCommandData().Name {
	case "paint":
		cmds.PaintHandler(s, i)
	case "paint-list":
		cmds.PaintListHandler(s, i)
	}
}
