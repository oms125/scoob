package bot

import (
	"context"
	"errors"

	"github.com/bwmarrin/discordgo"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type DiscordBotManager struct {
	session *discordgo.Session
}

func (d *DiscordBotManager) Start(ctx context.Context) error {
	logger := logf.FromContext(ctx)

	<-ctx.Done()
	logger.Info("Stopping bot")

	return d.GetSession().Close()
}

func (d *DiscordBotManager) GetSession() *discordgo.Session {
	return d.session
}

func (d *DiscordBotManager) SetSession(session *discordgo.Session) error {
	if d == nil {
		return errors.New("DiscordBotManager nil")
	}
	if session == nil {
		return errors.New("session nil")
	}
	d.session = session
	return nil
}

func (d *DiscordBotManager) SendMessage(channelID string, content string) error {
	_, err := d.session.ChannelMessageSend(channelID, content)
	return err
}
