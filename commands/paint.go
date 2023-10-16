package commands

import (
	"database/sql"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/yechentide/mhnow-bot/dao"
	"github.com/yechentide/mhnow-bot/utils"
)

func PaintCommand() (*discordgo.ApplicationCommand, error) {
	query := dao.New(dao.GetDb())
	monters, err := query.ListDiscoverableMonsters(dao.GetCtx())
	if err != nil {
		return nil, err
	}

	choices := []*discordgo.ApplicationCommandOptionChoice{}
	for _, m := range monters {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  m.JpName,
			Value: m.ID,
		})
	}

	return &discordgo.ApplicationCommand{
		Type:        1,
		Name:        "paint",
		Description: "ペイントしたモンスターを登録します",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "ランク",
				Type:        4,
				Required:    true,
				Description: "数字で入力してください",
			},
			{
				Name:        "種類",
				Type:        3,
				Required:    true,
				Description: "モンスター名を入力してください",
				Choices:     choices,
			},
			{
				Name:        "消滅日時",
				Type:        3,
				Required:    true,
				Description: "例: 09/09 12:30",
			},
			{
				Name:        "場所",
				Type:        3,
				Required:    false,
				Description: "任意です",
			},
		},
	}, err
}

func registerHunterIfNotExist(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	hQuery := dao.New(dao.GetDb())
	name := i.Member.User.Username
	nickName := i.Member.Nick

	_, err := hQuery.GetHunter(dao.GetCtx(), i.Member.User.ID)
	if err != sql.ErrNoRows {
		return err
	}
	_, err = hQuery.RegisterHunter(dao.GetCtx(), dao.RegisterHunterParams{
		ID:          i.Member.User.ID,
		Name:        name,
		DisplayName: sql.NullString{String: nickName, Valid: true},
	})
	return err
}

func PaintHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	rank := opts[0].IntValue()
	monsterID := opts[1].StringValue()
	disappearAt, err := utils.ConvertDateInput(opts[2].StringValue(), "Asia/Tokyo")
	if err != nil {
		msg := fmt.Sprintf("<@%s>\n日時のフォーマットを見直してください\n入力した値: %s", i.Member.User.ID, opts[2].StringValue())
		sendInteractionRespondMessage(s, i, msg)
		return
	}
	location := ""
	if len(opts) >= 4 {
		location = opts[3].StringValue()
	}

	err = registerHunterIfNotExist(s, i)
	if err != nil {
		msg := fmt.Sprintf("<@%s>\nハンター情報が見つかりません", i.Member.User.ID)
		sendInteractionRespondMessage(s, i, msg)
		return
	}

	phQuery := dao.New(dao.GetDb())
	_, err = phQuery.PaintMonster(dao.GetCtx(), dao.PaintMonsterParams{
		Rank:        int16(rank),
		HunterID:    i.Member.User.ID,
		MonsterID:   monsterID,
		DisappearAt: disappearAt,
		Location:    sql.NullString{String: location, Valid: true},
	})
	if err != nil {
		msg := fmt.Sprintf("<@%s>\nデータの保存に失敗しました", i.Member.User.ID)
		sendInteractionRespondMessage(s, i, msg)
		return
	}

	snowflake := utils.GetEmojiSnowflake(s, i.GuildID, monsterID)
	msg := fmt.Sprintf("%s(R%d) %sまで", snowflake, rank, opts[2].StringValue())
	if location != "" {
		msg += fmt.Sprintf("\n%s", location)
	}
	s.InteractionRespond(
		i.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: msg,
			},
		},
	)
}
