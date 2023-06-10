package models

type Mention struct {
	Index  int    `json:"index"`
	Length int    `json:"length"`
	Type   string `json:"type"`
	UserId string `json:"userId"`
}

type Emoji struct {
	Index     int    `json:"index"`
	Length    int    `json:"length"`
	ProductId string `json:"productId"`
	EmojiId   string `json:"emojiId"`
}

type Source struct {
	Type    string `json:"type"`
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Message struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Text    string  `json:"text"`
	Emojis  []Emoji `json:"emojis"`
	Mention Mention `json:"mention"`
}

type Event struct {
	ReplyToken      string          `json:"replyToken"`
	Type            string          `json:"type"`
	Mode            string          `json:"mode"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	WebhookEventId  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
	Message         Message         `json:"message"`
}

type LineMessage struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}
