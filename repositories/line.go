package repositories

import "github.com/line/line-bot-sdk-go/v7/linebot"

type LineRepository struct {
	lintBot *linebot.Client
}

func NewLineRepository(lintBot *linebot.Client) *LineRepository {
	return &LineRepository{
		lintBot,
	}
}

func (line *LineRepository) SendMessage() {
	return
}
