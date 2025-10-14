package bot

import (
	"github.com/go-logr/logr"
)

type BotLogger struct {
	Logger            logr.LogSink
	DiscordBotManager *DiscordBotManager
}

func (b *BotLogger) Init(info logr.RuntimeInfo) {
	b.Logger.Init(info)
}

func (b *BotLogger) Enabled(level int) bool {
	return b.Logger.Enabled(level)
}

func (b *BotLogger) Info(level int, msg string, keysAndValues ...interface{}) {
	b.DiscordBotManager.LogInfo(msg, keysAndValues...)

	b.Logger.Info(level, msg, keysAndValues...)
}

func (b *BotLogger) Error(err error, msg string, keysAndValues ...interface{}) {

	b.Logger.Error(err, msg, keysAndValues...)
}

func (b *BotLogger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	return &BotLogger{
		Logger:            b.Logger.WithValues(keysAndValues...),
		DiscordBotManager: b.DiscordBotManager,
	}
}

func (b *BotLogger) WithName(name string) logr.LogSink {
	return &BotLogger{
		Logger:            b.Logger.WithName(name),
		DiscordBotManager: b.DiscordBotManager,
	}
}
