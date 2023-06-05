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

func (r *LineRepository) SendOperations(replyToken string) error {
	bubble := linebot.BubbleContainer{
		Size: linebot.FlexBubbleSizeTypeKilo,
		Header: &linebot.BoxComponent{
			BackgroundColor: "#4CC764",
			PaddingAll:      linebot.FlexComponentPaddingTypeXl,
			Layout:          linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "🌥️ Good morning wording",
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#FFFFFF",
				},
			},
		},
		Body: &linebot.BoxComponent{
			Layout:     linebot.FlexBoxLayoutTypeVertical,
			PaddingAll: linebot.FlexComponentPaddingTypeLg,
			Spacing:    linebot.FlexComponentSpacingTypeMd,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Style:  linebot.FlexButtonStyleTypePrimary,
					Action: linebot.NewMessageAction("Get words", "getwords"),
					Height: linebot.FlexButtonHeightTypeSm,
				},
				&linebot.ButtonComponent{
					Style:  linebot.FlexButtonStyleTypePrimary,
					Action: linebot.NewMessageAction("Add word", "addword"),
					Height: linebot.FlexButtonHeightTypeSm,
				},
			},
		},
	}

	container := linebot.CarouselContainer{Contents: []*linebot.BubbleContainer{&bubble}}
	flexMessage := linebot.NewFlexMessage("operations", &container)

	_, err := r.lintBot.ReplyMessage(replyToken, flexMessage).Do()
	if err != nil {
		return err
	}

	return nil
}

func (r *LineRepository) SendMessageWords(replyToken string, words []models.Word) error {
	bubbles := []*linebot.BubbleContainer{}

	for _, v := range words {
		bubbles = append(bubbles, &linebot.BubbleContainer{
			Size: linebot.FlexBubbleSizeTypeMicro,
			Header: &linebot.BoxComponent{
				BackgroundColor: "#4CC764",
				PaddingAll:      linebot.FlexComponentPaddingTypeXl,
				Layout:          linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Text:   "🌥️ Good morning wording",
						Weight: linebot.FlexTextWeightTypeBold,
						Color:  "#FFFFFF",
					},
				},
			},
			Body: &linebot.BoxComponent{
				Layout:     linebot.FlexBoxLayoutTypeVertical,
				PaddingAll: linebot.FlexComponentPaddingTypeLg,
				Spacing:    linebot.FlexComponentSpacingTypeMd,
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

func (r *LineRepository) SendAddWordDescription(replyToken string) error {
	bubble := linebot.BubbleContainer{
		Size: linebot.FlexBubbleSizeTypeMega,
		Header: &linebot.BoxComponent{
			BackgroundColor: "#4CC764",
			PaddingAll:      linebot.FlexComponentPaddingTypeXl,
			Layout:          linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "🌥️ Good morning wording",
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#FFFFFF",
				},
			},
		},
		Body: &linebot.BoxComponent{
			Layout:     linebot.FlexBoxLayoutTypeVertical,
			PaddingAll: linebot.FlexComponentPaddingTypeXl,
			Spacing:    linebot.FlexComponentSpacingTypeMd,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Text:   "Add Word",
					Size:   linebot.FlexTextSizeTypeMd,
					Weight: linebot.FlexTextWeightTypeBold,
				},
				&linebot.BoxComponent{
					Layout:  linebot.FlexBoxLayoutTypeBaseline,
					Spacing: linebot.FlexComponentSpacingTypeMd,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Text:  "Command",
							Size:  linebot.FlexTextSizeTypeSm,
							Color: "#aaaaaa",
							Flex:  linebot.IntPtr(2),
						},
						&linebot.TextComponent{
							Text:  "aw|\n\"word to add\"|\n\"meaning of word\"",
							Size:  linebot.FlexTextSizeTypeSm,
							Wrap:  true,
							Color: "#666666",
							Flex:  linebot.IntPtr(3),
						},
					},
				},
			},
		},
	}

	container := linebot.CarouselContainer{Contents: []*linebot.BubbleContainer{&bubble}}
	flexMessage := linebot.NewFlexMessage("operations", &container)

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
