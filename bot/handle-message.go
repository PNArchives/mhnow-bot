package bot

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

/*
MessageCreate
MessageUpdate
MessageDelete
*/

func isBotMentionedInMessage(s *discordgo.Session, mentions []*discordgo.User) bool {
	for _, mention := range mentions {
		if mention.ID == s.State.User.ID {
			return true
		}
	}
	return false
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.GuildID == "" {
		onDirectMessageCreated(s, m)
	} else if isBotMentionedInMessage(s, m.Mentions) {
		onGuildMessageCreated(s, m)
	}
}

func onMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if !isBotMentionedInMessage(s, m.Mentions) {
		return
	}
	slog.Info(fmt.Sprintf(
		"Message updated: Channel=%s, Guild=%s, Author=%s_%-20s\n%s",
		m.ChannelID,
		m.GuildID,
		m.Author.Username,
		m.Author.ID,
		m.Content,
	))
}

func onMessageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {}

func onGuildMessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	slog.Info(fmt.Sprintf(
		"Guild message created: Channel=%s, Guild=%s, Author=%s(%s)_%-20s\n%s",
		m.ChannelID,
		m.GuildID,
		m.Author.Username,
		m.Member.Nick,
		m.Author.ID,
		m.Content,
	))
}

func onDirectMessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	slog.Info(fmt.Sprintf(
		"Direct message created: Channel=%s, Author=%s_%-20s\n%s",
		m.ChannelID,
		m.Author.Username,
		m.Author.ID,
		m.Content,
	))
}
