package klaus

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	StateUndefined = iota
	StateGreeted
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Alias    string `json:"alias,omitempty"`
	Lang     string `json:"lang,omitempty"`
	Messages []int  `json:"messages,omitempty"`
	State    int    `json:"state,omitempty"`
}

func NewUser(upd tg.Update) User {
	tguser := upd.SentFrom()

	return User{
		ID:       tguser.ID,
		Username: tguser.FirstName + " " + tguser.LastName,
		Alias:    tguser.UserName,
		Lang:     tguser.LanguageCode,
		Messages: make([]int, 0),
		State:    StateUndefined,
	}
}
