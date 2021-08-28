package main

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`

}

type RestResponse struct {
	Result []Update 	`json:"result"`
}

type Message struct{
	Chat Chat			`json:"chat"`
	text string			`json:"text"`
}

type Chat struct {
		ChatId int		`json:"id"`
}

type BotMessage struct{
	ChatId int			`json:"chat_id"`
	Text string			`json:"text"`
}