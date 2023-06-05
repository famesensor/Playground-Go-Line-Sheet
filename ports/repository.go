package ports

import "github.com/famesensor/playground-go-line-sheet/models"

type LineRepository interface {
	SendOperations(replyToken string) error
	SendMessageWords(replyToken string, words []models.Word) error
	SendAddWordDescription(replyToken string) error
	SendMessageError(replayToken string, msgErr string) error
}

type SheetRepository interface {
	GetValues(spreadsheetId string, rangeSheet string) ([]models.Word, error)
}
