package category_nurzhas_store

type ChatIdStore interface {
	Create(ch *ChatId) (*ChatId, error)
	List(telegramBotId string) ([]ChatId, error)
}
