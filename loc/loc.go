package loc

type Localization interface {
	UnexpectedMessageText() string
	EnterPrefReplySuccess() string
	EnterPrefReplyRemove() string
	EnterPrefReplyNext() string
	EnterPrefReplyUnkwonMessage() string
	EnterPrefReplyZeroPref() string
	EnterPrefReplyFinish() string
	EnterPrefButtonContinue() string
	EnterPrefButtonEnd() string
	EnterPrefButtonRemove() string
	PrefIntroMessage() string
	StartButtonYes() string
	StartMessage() string
	RulesButtonYes() string
	RulesMessage() string
	StartupMessage() string
	GiftForMessage(name, alias string, prefs []string) string
	RegistrationClosed() string
	InfoGiftMessage(name, alias string, prefs []string) string
}
