package ports

import "github.com/famesensor/playground-go-line-sheet/models"

type LineRepository interface {
	SendMessageWords(replyToken string, words []models.Word) error
	SendMessageError(replayToken string, msgErr string) error
}

type SheetRepository interface {
	GetValues(spreadsheetId string, rangeSheet string) ([]models.Word, error)
}
