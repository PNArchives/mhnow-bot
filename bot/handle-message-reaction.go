package bot

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

/*
MessageReactionAdd
MessageReactionRemove
MessageReactionRemoveAll
*/

func onMessageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	snowflake := r.Emoji.Name + ":" + r.Emoji.ID
	if r.GuildID == "" {
		slog.Debug(fmt.Sprintf(
			"Direct reaction added: Channel=%s, Guild=%s, MessageID=%s, Emoji=%s",
			r.ChannelID,
			r.GuildID,
			r.MessageID,
			snowflake,
		))
	} else {
		slog.Debug(fmt.Sprintf(
			"Guild reaction added: Channel=%s, Guild=%s, Author=%s(%s)_%-20s, MessageID=%s, Emoji=%s",
			r.ChannelID,
			r.GuildID,
			r.Member.User.Username,
			r.Member.Nick,
			r.Member.User.ID,
			r.MessageID,
			snowflake,
		))
	}
}

func onMessageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	snowflake := r.Emoji.Name + ":" + r.Emoji.ID
	if r.GuildID == "" {
		slog.Debug(fmt.Sprintf(
			"Direct reaction removed: Channel=%s, Guild=%s, User=%s, MessageID=%s, Emoji=%s",
			r.ChannelID,
			r.GuildID,
			r.UserID,
			r.MessageID,
			snowflake,
		))
	} else {
		slog.Debug(fmt.Sprintf(
			"Guild reaction removed: Channel=%s, Guild=%s, User=%s, MessageID=%s, Emoji=%s",
			r.ChannelID,
			r.GuildID,
			r.UserID,
			r.MessageID,
			snowflake,
		))
	}
}

func onMessageReactionRemoveAll(s *discordgo.Session, r *discordgo.MessageReactionRemoveAll) {
	if r.GuildID == "" {
		slog.Debug(fmt.Sprintf(
			"All direct reaction removed: Channel=%s, Guild=%s, User=%s, MessageID=%s",
			r.ChannelID,
			r.GuildID,
			r.UserID,
			r.MessageID,
		))
	} else {
		slog.Debug(fmt.Sprintf(
			"All guild reaction removed: Channel=%s, Guild=%s, User=%s, MessageID=%s",
			r.ChannelID,
			r.GuildID,
			r.UserID,
			r.MessageID,
		))
	}
}
