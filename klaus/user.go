package klaus

import (
	"encoding/json"

	"github.com/dartt0n/klaus/loc"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	StateStart = iota
	StateRules
	StatePreferences
	StateEnterPref
	StateContinuePref
	StateWait
	StateLocale
)

type User struct {
	ID       int64            `json:"id,omitempty"`
	Username string           `json:"username,omitempty"`
	Alias    string           `json:"alias,omitempty"`
	Lang     string           `json:"lang,omitempty"`
	Messages []int            `json:"messages,omitempty"`
	State    int              `json:"state,omitempty"`
	Prefs    []string         `json:"preferences,omitempty"`
	Loc      loc.Localization `json:"loc,omitempty"`
	GiftFor  int64            `json:"giftFor,omitempty"`
}

func NewUser(upd tg.Update) User {
	tguser := upd.SentFrom()

	return User{
		ID:       tguser.ID,
		Username: tguser.FirstName + " " + tguser.LastName,
		Alias:    tguser.UserName,
		Lang:     tguser.LanguageCode,
		Messages: make([]int, 0),
		State:    StateStart,
		Prefs:    make([]string, 0),
	}
}

func (u *User) UnmarshalJSON(bytes []byte) error {
	var data struct {
		ID       int64    `json:"id,omitempty"`
		Username string   `json:"username,omitempty"`
		Alias    string   `json:"alias,omitempty"`
		Lang     string   `json:"lang,omitempty"`
		Messages []int    `json:"messages,omitempty"`
		State    int      `json:"state,omitempty"`
		Prefs    []string `json:"preferences,omitempty"`
		GiftFor  int64    `json:"giftFor,omitempty"`
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return err
	}

	u.ID = data.ID
	u.Username = data.Username
	u.Alias = data.Alias
	u.Lang = data.Lang
	u.Messages = data.Messages
	u.State = data.State
	u.Prefs = data.Prefs
	u.GiftFor = data.GiftFor

	switch u.Lang {
	case loc.ENG.Lang:
		u.Loc = &loc.ENG
	case loc.RUS.Lang:
		u.Loc = &loc.RUS
	default:
		u.Loc = &loc.ENG
	}

	return nil
}
