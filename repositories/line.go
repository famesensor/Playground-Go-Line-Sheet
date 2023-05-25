package repositories

import (
	"log"

	"github.com/famesensor/playground-go-line-sheet/models"
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

func (r *LineRepository) SendMessageWords(replyToken string, words []models.Word) error {
	bubbles := []*linebot.BubbleContainer{}

	for _, v := range words {
		bubbles = append(bubbles, &linebot.BubbleContainer{
			Size: linebot.FlexBubbleSizeTypeMicro,
			Header: &linebot.BoxComponent{
				BackgroundColor: "#4CC764",
				PaddingAll:      linebot.FlexComponentPaddingTypeLg,
				Layout:          linebot.FlexBoxLayoutTypeVertical,
				Spacing:         linebot.FlexComponentSpacingTypeSm,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "🌥️ Good morning wording",
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  "#FFFFFF",
					},
				},
			},
			Body: &linebot.BoxComponent{
				Layout:          linebot.FlexBoxLayoutTypeVertical,
				PaddingAll:      linebot.FlexComponentPaddingTypeLg,
				Spacing:         linebot.FlexComponentSpacingTypeSm,
				BackgroundColor: "#EFEFEF",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Text:   "📄 : " + v.EngWord,
						Weight: linebot.FlexTextWeightTypeBold,
						Size:   linebot.FlexTextSizeTypeSm,
						Wrap:   true,
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "📃 : " + v.ThaiMeaning,
						Size: linebot.FlexTextSizeTypeXs,
						Wrap: true,
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "🕒 : " + v.CreatedDate,
						Size: linebot.FlexTextSizeTypeXxs,
						Wrap: true,
					},
				},
			},
		})
	}

	container := linebot.CarouselContainer{Contents: bubbles}
	flexMessage := linebot.NewFlexMessage("words", &container)

	_, err := r.lintBot.ReplyMessage(replyToken, flexMessage).Do()
	if err != nil {
		return err
	}

	return nil
}

func (r *LineRepository) SendMessageError(replyToken string, msgErr string) error {
	textMessage := linebot.NewTextMessage(msgErr)

	_, err := r.lintBot.ReplyMessage(replyToken, textMessage).Do()
	if err != nil {
		log.Println("error line send error ->", err)
		// TODO: handle error when send message error
	}
	return nil
}
