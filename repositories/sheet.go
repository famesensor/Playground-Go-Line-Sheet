package repositories

import (
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
