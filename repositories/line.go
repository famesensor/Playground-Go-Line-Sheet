package repositories

import (
	"log"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineRepository struct {
	lintBot *linebot.Client
}

func NewLineRepository(lintBot *linebot.Client) *LineRepository {
	return &LineRepository{
		lintBot,
	}
}

func (line *LineRepository) SendMessageError(replyToken string, msgErr string) error {
	textMessage := linebot.NewTextMessage(msgErr)

	_, err := line.lintBot.ReplyMessage(replyToken, textMessage).Do()
	if err != nil {
		log.Println("error line send error ->", err)
		// TODO: handle error when send message error
	}
	return nil
}
