package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var slashcommands []*discordgo.ApplicationCommand = []*discordgo.ApplicationCommand{
	{
		Name:        "grouping",
		Description: "メンバーのチーム分けを行います",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "チーム数",
				Description: "チーム数を指定します",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "メンバーリスト",
				Description: "メンバーリストをカンマ区切りで指定します (例: user1,user2,user3)。指定しない場合はボイスチャンネル内のメンバーが対象になります",
				Required:    false,
			},
		},
	},
}

func RegisterSlashCommands(s *discordgo.Session) error {
	for _, sc := range slashcommands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", sc)
		if err != nil {
			log.Printf("スラッシュコマンドの登録に失敗しました: %v", err)
			return err
		} else {
			log.Printf("スラッシュコマンド '%s' を登録しました", sc.Name)
		}
	}

	return nil
}
