package category_nurzhas_store

type CreateTelegramBotCommand struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

func (cmd *CreateTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).CreateTelegramBot(cmd)
}

type GetTelegramBotCommand struct {
	Id string `json:"id"`
}

func (cmd *GetTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).GetTelegramBot(cmd)
}

type DeleteTelegramBotCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(TelegramService).DeleteTelegramBot(cmd)
}

type ListTelegramBotCommand struct {
}

func (cmd *ListTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).ListTelegramBot(cmd)
}

type SendMessageCommand struct {
	TelegramBoId string `json:"_"`
	Message      string `json:"_"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	ParseMode    string `json:"_"`
}

func (cmd *SendMessageCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(TelegramService).SendMessage(cmd)
}

type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	Chat Chat `json:"chat"`
}

type Result struct {
	Message Message `json:"message"`
}

type GetUpdates struct {
	Result []Result `json:"result"`
}

type SendTelegramMessage struct {
	ChatId    string `json:"chat_id"`
	ParseMode string `json:"parse_mode"`
	Text      string `json:"text"`
}
