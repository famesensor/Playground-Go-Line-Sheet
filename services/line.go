package services

import "github.com/famesensor/playground-go-line-sheet/ports"

type LineService struct {
	lineRepo      ports.LineRepository
	sheetRepo     ports.SheetRepository
	spreadSheetId string
}

func NewLineService(lineRepo ports.LineRepository, sheetRepo ports.SheetRepository, spreadSheetId string) *LineService {
	return &LineService{
		lineRepo,
		sheetRepo,
		spreadSheetId,
	}
}

func (lineSrv *LineService) GetWords(replyToken string) {
	words, err := lineSrv.sheetRepo.GetValues(lineSrv.spreadSheetId, "Wording!A:C")
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}

	err = lineSrv.lineRepo.SendMessageWords(replyToken, words)
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}
}

func (lineSrv *LineService) SendMessageError(replyToken string) error {
	message := "Operation is wrong or Internal server error"
	return lineSrv.lineRepo.SendMessageError(replyToken, message)
}
