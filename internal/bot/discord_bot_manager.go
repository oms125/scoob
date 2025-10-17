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

func (d *DiscordBotManager) SetLogChannel(logChannel string) {
	d.LogChannel = logChannel
	d.LogInfo("Logging Initialized Successfully")
}

func (d *DiscordBotManager) LogInfo(msg string, keysAndValues ...interface{}) {
	if d.LogChannel != "" {
		_, _ = d.Session.ChannelMessageSendEmbed(
			d.LogChannel,
			&discordgo.MessageEmbed{
				Title: "Info Log",
				Color: 0x13f737,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Message",
						Value: msg,
					},
				},
			},
		)
	}
}

func (d *DiscordBotManager) LogError(err error, msg string, keysAndValues ...interface{}) {
	if d.LogChannel != "" {
		_, _ = d.Session.ChannelMessageSendEmbed(
			d.LogChannel,
			&discordgo.MessageEmbed{
				Title: "Error Log",
				Color: 0xf71414,
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Message",
						Value: msg,
					},
					{
						Name:  "Error",
						Value: err.Error(),
					},
				},
			},
		)
	}
}

func (d *DiscordBotManager) SendMessage(channelID string, content string) error {
	_, err := d.Session.ChannelMessageSend(channelID, content)
	return err
}
