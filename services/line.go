package services

import (
	"context"
	"strings"
	"time"

	"github.com/famesensor/playground-go-line-sheet/models"
	"github.com/famesensor/playground-go-line-sheet/ports"
)

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

func (lineSrv *LineService) GetOperations(replyToken string) {
	err := lineSrv.lineRepo.SendOperations(replyToken)
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
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

func (lineSrv *LineService) AddWordDescription(replyToken string) {
	err := lineSrv.lineRepo.SendAddWordDescription(replyToken)
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}
}

func (lineSrv *LineService) AddWord(ctx context.Context, replyToken, word, mean string) {
	checkWord, err := lineSrv.checkWordDuplicate(word)
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}

	if checkWord {
		err = lineSrv.lineRepo.SendAddWord(replyToken, "Word duplicate")
		if err != nil {
			lineSrv.SendMessageError(replyToken)
			return
		}
		return
	}

	wordModel := models.Word{EngWord: word, ThaiMeaning: mean, CreatedDate: time.Now().Format("01/02/2006")}
	err = lineSrv.sheetRepo.AddValue(ctx, lineSrv.spreadSheetId, wordModel)
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}

	err = lineSrv.lineRepo.SendAddWord(replyToken, "Add word successfully")
	if err != nil {
		lineSrv.SendMessageError(replyToken)
		return
	}
}

func (lineSrv *LineService) checkWordDuplicate(word string) (bool, error) {
	words, err := lineSrv.sheetRepo.GetValues(lineSrv.spreadSheetId, "Wording!A:C")
	if err != nil {
		return false, err
	}

	mapWords := map[string]bool{}
	for _, v := range words {
		mapWords[strings.ToLower(v.EngWord)] = true
	}

	return mapWords[word], nil
}

func (lineSrv *LineService) SendMessageError(replyToken string) error {
	message := "Operation is wrong or Internal server error"
	return lineSrv.lineRepo.SendMessageError(replyToken, message)
}
