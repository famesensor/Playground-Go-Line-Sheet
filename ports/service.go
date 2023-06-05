package ports

type LineService interface {
	GetOperations(replyToken string)
	GetWords(replyToken string)
	AddWordDescription(replyToken string)
	// AddWord()
	// DeleteWord()
	// SendMessage(replayToken string) error
	SendMessageError(replyToken string) error
}
