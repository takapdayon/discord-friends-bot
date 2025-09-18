package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type groupingContext struct {
	teamCount  int
	memberList []string
}

func NewGroupingContext(s *discordgo.Session, i *discordgo.InteractionCreate) *groupingContext {
	options := i.ApplicationCommandData().Options
	gc := &groupingContext{}

	for _, opt := range options {
		// FIXME: Name以外での判定を考える
		switch opt.Name {
		case "チーム数":
			gc.teamCount = int(opt.IntValue())
		case "メンバーリスト":
			if opt.StringValue() != "" {
				gc.memberList = strings.Split(opt.StringValue(), ",")
			}
		}
	}

	return gc
}

func Grouping(s *discordgo.Session, i *discordgo.InteractionCreate) {
	gc := NewGroupingContext(s, i)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "チーム分け結果",
					Description: "チーム分けを実行しました！\n",
					Color:       0x006400,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "チーム数",
							Value:  strconv.Itoa(gc.teamCount),
							Inline: true,
						},
						{
							Name:   "メンバーリスト",
							Value:  strings.Join(gc.memberList, ", "),
							Inline: true,
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("チーム分けコマンドの応答中にエラーが発生しました: %v", err.Error())
	}
}
