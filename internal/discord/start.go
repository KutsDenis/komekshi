package discord

import (
	"github.com/bwmarrin/discordgo"
	"komekshi/internal/cfg"
	"komekshi/pkg/logger"
)

var Bot *discordgo.Session

func Start(l *logger.Logger) {
	Bot = &discordgo.Session{}

	Bot, _ = discordgo.New("Bot " + cfg.Get.Token)

	Bot.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		l.Info("Bot is ready")
	})

	Bot.AddHandler(messageHandler)

	err := Bot.Open()
	if err != nil {
		l.Fatal(err)
	}

	l.Info("Bot is running")
}
