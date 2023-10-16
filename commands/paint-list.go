package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/yechentide/mhnow-bot/dao"
	"github.com/yechentide/mhnow-bot/utils"
)

func PaintListCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Type:        1,
		Name:        "paint-list",
		Description: "消滅していないモンスターの一覧を表示します",
	}
}

func PaintListHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	query := dao.New(dao.GetDb())
	monsters, err := query.FindHuntableMontersWithGuild(dao.GetCtx())
	if err != nil {
		sendInteractionRespondMessage(s, i, "データの取得に失敗しました")
		return
	}

	msg := ""
	if len(monsters) == 0 {
		msg = "消滅していないモンスターの一覧はありません"
		sendInteractionRespondMessage(s, i, msg)
		return
	}

	emojiMap := utils.GenerateGuildEmojiMap(s, i.GuildID)
	for _, monster := range monsters {
		monsterName := monster.MonsterJpName
		snowflake, ok := (*emojiMap)[monster.MonsterID]
		if ok {
			monsterName = snowflake
		}
		jstTime := monster.DisappearAt.Format("01/02_15:04")
		location := monster.Location.String
		if location == "" {
			location = "???"
		}
		msg += fmt.Sprintf(
			"## %s(R%d)   *%s*まで\n> %sが%sで発見した\n",
			monsterName,
			monster.Rank,
			jstTime,
			monster.HunterName,
			location,
		)
	}
	sendInteractionRespondMessage(s, i, msg)
}
