package sheet

import (
	"context"

	"github.com/famesensor/playground-go-line-sheet/configs"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ConnectGoogleSheet(ctx context.Context, config *configs.Config) (*sheets.Service, error) {
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(config.SheetCredentialFileName), option.WithScopes(sheets.SpreadsheetsScope))
	if err != nil {
		return nil, err
	}

	return srv, nil
}
