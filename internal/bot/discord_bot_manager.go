package bot

import (
	"context"
	"errors"

	"github.com/bwmarrin/discordgo"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type DiscordBotManager struct {
	Session    *discordgo.Session
	LogChannel string
}

func (d *DiscordBotManager) Start(ctx context.Context) error {
	logger := logf.FromContext(ctx)

	<-ctx.Done()
	logger.Info("Stopping bot")

	return d.GetSession().Close()
}

func (d *DiscordBotManager) GetSession() *discordgo.Session {
	return d.Session
}

func (d *DiscordBotManager) SetSession(session *discordgo.Session) error {
	if d == nil {
		return errors.New("DiscordBotManager nil")
	}
	if session == nil {
		return errors.New("session nil")
	}
	d.Session = session
	return nil
}

func (d *DiscordBotManager) SetLogChannel(logChannel string) error {
	err := d.SendMessage(logChannel, "Logging Initialized Successfully")
	if err != nil {
		return errors.New("failed to initialize logging")
	}
	d.LogChannel = logChannel
	return nil
}

func (d *DiscordBotManager) LogInfo(msg string, keysAndValues ...interface{}) {
	if d.LogChannel != "" {
		_ = d.SendMessage(d.LogChannel, "```"+msg+"```")
	}
}

func (d *DiscordBotManager) LogError(err error, msg string, keysAndValues ...interface{}) {
	if d.LogChannel != "" {
		_ = d.SendMessage(d.LogChannel, "```"+err.Error()+msg+"```")
	}
}

func (d *DiscordBotManager) SendMessage(channelID string, content string) error {
	_, err := d.Session.ChannelMessageSend(channelID, content)
	return err
}
