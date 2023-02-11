package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"komekshi/pkg/logger"
	"time"
)

type Eventer interface {
	Create() *discordgo.GuildScheduledEvent
}

type Event struct {
	Name        string
	Description string
	StartAfter  time.Duration
	Length      time.Duration
	ChanelID    string
	S           *discordgo.Session
	L           *logger.Logger
}

func (e Event) Create() *discordgo.GuildScheduledEvent {
	startingTime := time.Now().Add(e.StartAfter * time.Minute)
	endingTime := startingTime.Add(e.Length * time.Minute)
	guild := e.S.State.Guilds
	scheduledEvent, err := e.S.GuildScheduledEventCreate(guild[0].ID, &discordgo.GuildScheduledEventParams{
		Name:               e.Name,
		Description:        e.Description,
		ScheduledStartTime: &startingTime,
		ScheduledEndTime:   &endingTime,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          e.ChanelID,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})
	if err != nil {
		e.L.Error(err)
	}

	e.L.Info("Created: " + scheduledEvent.Name)
	return scheduledEvent
}

func Start(s *discordgo.Session, event *discordgo.GuildScheduledEvent, l *logger.Logger) {
	result, err := s.GuildScheduledEventEdit(event.GuildID, event.ID, &discordgo.GuildScheduledEventParams{Status: 2})
	if err != nil {
		l.Error(err)
		return
	}

	l.Info("Start event " + result.Name)
}

func List(s *discordgo.Session) []*discordgo.GuildScheduledEvent {
	l := logger.GetLogger()
	result, err := s.GuildScheduledEvents(s.State.Guilds[0].ID, true)
	if err != nil {
		l.Error(err)
	}

	return result
}

func GetEventUsers(s *discordgo.Session, event *discordgo.GuildScheduledEvent) []*discordgo.GuildScheduledEventUser {
	result, err := s.GuildScheduledEventUsers(event.GuildID, event.ID, 100, true, "", "")
	if err != nil {
		logrus.Error(err)
	}
	return result
}
