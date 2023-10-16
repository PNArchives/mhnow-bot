package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	fmt.Println(">>> " + msg)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}

func SendReply(s *discordgo.Session, channelID string, reference *discordgo.MessageReference, msg string) {
	_, err := s.ChannelMessageSendReply(channelID, msg, reference)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}

func SendEmbedMessage(s *discordgo.Session, channelID, title, desc string, color int) {
	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Color:       color,
	}
	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		fmt.Println("Error sending embed message: ", err)
	}
}

func SendReaction(s *discordgo.Session, channelID, messageID, reaction string) {
	err := s.MessageReactionAdd(channelID, messageID, reaction)
	if err != nil {
		fmt.Println("Error add a reaction: ", err)
	}
}
