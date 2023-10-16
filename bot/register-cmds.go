package bot

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	cmds "github.com/yechentide/mhnow-bot/commands"
)

func RegisterCommands(s *discordgo.Session, guildId string) {
	registerMessage := fmt.Sprintf("Registering commands for guild %s ...", guildId)
	slog.Info(registerMessage)

	paintListCmd := cmds.PaintListCommand()
	paintCmd, err := cmds.PaintCommand()
	if err != nil {
		panic(err)
	}
	_, err = s.ApplicationCommandBulkOverwrite(s.State.User.ID, guildId, []*discordgo.ApplicationCommand{
		paintListCmd, paintCmd,
	})
	if err != nil {
		panic(err)
	}
}
