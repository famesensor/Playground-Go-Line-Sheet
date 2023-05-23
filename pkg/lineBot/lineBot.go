package lineBot

import (
	"github.com/famesensor/playground-go-line-sheet/configs"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func ConnectLineBot(config *configs.Config) (*linebot.Client, error) {
	return linebot.New(config.LineChannelSecret, config.LineChannelToken)
}
