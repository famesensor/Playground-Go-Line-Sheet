package ports

import "context"

type LineService interface {
	GetOperations(replyToken string)
	GetWords(replyToken string)
	AddWordDescription(replyToken string)
	AddWord(ctx context.Context, replayToken, word, mean string)
	SendMessageError(replyToken string) error
}
