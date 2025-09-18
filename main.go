package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/takapdayon/discord-friends-bot/internal/commands"
	"github.com/takapdayon/discord-friends-bot/internal/handlers"
)

var state struct {
	Token string
}

func init() {
	// 初期化処理
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".envファイルの読み込み中にエラーが発生しました: %v", err)
	}
	state.Token = os.Getenv("DISCORD_BOT_TOKEN")
	if state.Token == "" {
		log.Fatalf("環境変数DISCORD_BOT_TOKENが設定されていません")
	}
}

func main() {
	dg, err := discordgo.New("Bot " + state.Token)
	if err != nil {
		log.Fatalf("Discordセッションの作成中にエラーが発生しました: %v", err.Error())
	}

	dg.AddHandler(handlers.HandleInteractionCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalf("接続の開始中にエラーが発生しました: %v", err.Error())
	}
	defer dg.Close()

	err = commands.RegisterSlashCommands(dg)
	if err != nil {
		log.Fatalf("スラッシュコマンドの登録中にエラーが発生しました: %v", err)
	}

	log.Println("Botが起動しました。終了するにはCTRL-Cを押してください。")
	select {}

}
