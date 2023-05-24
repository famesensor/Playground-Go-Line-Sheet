package ports

type LineRepository interface {
	SendMessageError(replayToken string, msgErr string) error
}
