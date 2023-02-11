package discord

import (
	"github.com/bwmarrin/discordgo"
	"komekshi/internal/cfg"
	"komekshi/pkg/logger"
)

func Start(l *logger.Logger) {
	bot, err := discordgo.New("Bot " + cfg.Get.Token)
	if err != nil {
		l.Error(err)
	}

	bot.AddHandler(messageHandler)

	err = bot.Open()
	if err != nil {
		l.Fatal(err)
	}
	l.Info("Bot is running")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "ping":
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	default:
		_, _ = s.ChannelMessageSend(m.ChannelID, "hey")
	}
}
