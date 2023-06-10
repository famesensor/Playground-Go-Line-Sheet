package repositories

import (
	"context"
	"fmt"

	"github.com/famesensor/playground-go-line-sheet/models"
	"google.golang.org/api/sheets/v4"
)

type SheetRepository struct {
	sheetsConn *sheets.Service
}

func NewSheetRepository(sheetsConn *sheets.Service) *SheetRepository {
	return &SheetRepository{
		sheetsConn,
	}
}

func (r *SheetRepository) GetValues(spreadSheetId string, rangeSheet string) ([]models.Word, error) {
	res, err := r.sheetsConn.Spreadsheets.Values.Get(spreadSheetId, rangeSheet).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return []models.Word{}, nil
	}

	data := make([]models.Word, len(res.Values))
	for i, row := range res.Values {
		data[i] = models.Word{EngWord: fmt.Sprintf("%v", row[0]), ThaiMeaning: fmt.Sprintf("%v", row[1]), CreatedDate: fmt.Sprintf("%v", row[2])}
	}

	return data, nil
}

func (r *SheetRepository) AddValue(ctx context.Context, spreadSheetId string, word models.Word) error {
	row := &sheets.ValueRange{
		Values: [][]interface{}{{word.EngWord, word.ThaiMeaning, word.CreatedDate}},
	}

	res, err := r.sheetsConn.Spreadsheets.Values.Append(spreadSheetId, "Wording", row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(ctx).Do()
	if err != nil || res.HTTPStatusCode != 200 {
		return err
	}

	return nil
}
