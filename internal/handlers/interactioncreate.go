package handlers

import (
	"github.com/takapdayon/discord-friends-bot/internal/commands"

	"github.com/bwmarrin/discordgo"
)

func commandSelector(s *discordgo.Session, i *discordgo.InteractionCreate) {

	switch i.ApplicationCommandData().Name {
	case "grouping":
		commands.Grouping(s, i)
	default:
		commands.Unknown(s, i)
	}
}

func HandleInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	commandSelector(s, i)
}
