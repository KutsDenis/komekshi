package discord

import (
	"github.com/bwmarrin/discordgo"
	"komekshi/internal/discord/events"
	"komekshi/pkg/logger"
	"strconv"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	l := logger.GetLogger()
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "need event":
		eventer := events.Event{
			Name:        "Тестовое",
			Description: "Тестируем событие",
			StartAfter:  1,
			Length:      2,
			ChanelID:    "570601549718355978",
			S:           s,
			L:           &l,
		}

		event := eventer.Create()
		_, _ = s.ChannelMessageSend(m.ChannelID, "done "+strconv.Itoa(int(event.Status)))
	case "event users":
		e := events.List(s)
		u := events.GetEventUsers(s, e[0])
		_, _ = s.ChannelMessageSend(m.ChannelID, u[0].User.Mention())
	}
}
