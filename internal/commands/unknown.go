package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Unknown(s *discordgo.Session, i *discordgo.InteractionCreate) {

	response := "不明なコマンドです。"

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: response,
		},
	})
	if err != nil {
		log.Printf("不明なコマンドの応答中にエラーが発生しました: %v", err.Error())
	}
}
