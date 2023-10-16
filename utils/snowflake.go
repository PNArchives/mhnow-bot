package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func GetEmojiSnowflake(s *discordgo.Session, guildID, monsterID string) string {
	m := GenerateGuildEmojiMap(s, guildID)
	snowflake, ok := (*m)[monsterID]
	if ok {
		return snowflake
	}
	return monsterID
}

func GenerateGuildEmojiMap(s *discordgo.Session, guildID string) *map[string]string {
	m := map[string]string{}
	emojis, _ := s.GuildEmojis(guildID)
	for _, e := range emojis {
		m[e.Name] = fmt.Sprintf("<:%s:%s>", e.Name, e.ID)
	}
	return &m
}
