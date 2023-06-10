package ports

import (
	"context"

	"github.com/famesensor/playground-go-line-sheet/models"
)

type LineRepository interface {
	SendOperations(replyToken string) error
	SendMessageWords(replyToken string, words []models.Word) error
	SendAddWordDescription(replyToken string) error
	SendAddWord(replyToken string, message string) error
	SendMessageError(replayToken string, msgErr string) error
}

type SheetRepository interface {
	GetValues(spreadsheetId string, rangeSheet string) ([]models.Word, error)
	AddValue(ctx context.Context, spreadSheetId string, word models.Word) error
}
