package loc

var (
	_   Localization = (*rusloc)(nil)
	RUS              = rusloc{}
)

type rusloc struct {
}

func (e *rusloc) UnexpectedMessageText() string {
	return "Пожалуйста, используйте кнопки на клавиатуре :)"
}

func (e *rusloc) EnterPrefReplySuccess() string {
	return "Отлично! Ваш текущий список предпочтений:\n\n"
}

func (e *rusloc) EnterPrefReplyRemove() string {
	return "Хорошо! Ваш текущий список предпочтений:\n\n"
}

func (e *rusloc) EnterPrefReplyNext() string {
	return `Хорошо! Введите следующее предпочтение:`
}

func (e *rusloc) EnterPrefReplyUnkwonMessage() string {
	return "Извините, я не понимаю это сообщение"
}

func (e *rusloc) EnterPrefReplyZeroPref() string {
	return "Пожалуйста, введите хотя бы одно предпочтение"
}

func (e *rusloc) EnterPrefReplyFinish() string {
	return `Здорово! Мы все подготовили!   

Теперь вам нужно ждать 11 декабря! В этот день, в 11:00, я отправлю вам имя человека, для которого вы подготавливаете подарок! Удачи, мой друг!`
}

func (e *rusloc) EnterPrefButtonContinue() string {
	return "Я хочу ввести еще одно предпочтение"
}

func (e *rusloc) EnterPrefButtonEnd() string {
	return "Все, что я хотел"
}

func (e *rusloc) EnterPrefButtonRemove() string {
	return "Удалить последнее предпочтение"
}

func (e *rusloc) PrefIntroMessage() string {
	return `Хо хо хо! 
Возможно, человек, который получит ваше имя, ничего о вас не знает. 

Вам нужно помочь ему подготовить хороший подарок для вас!

Итак, расскажите мне о вашем первом предпочтении:`
}

func (e *rusloc) StartButtonYes() string {
	return "Поехали!"
}

func (e *rusloc) StartMessage() string {
	return `Привет, мой друг! Это я, Санта! И я рад видеть тебя здесь! 🎅

Мои умные эльфы решили помочь мне с подарками для добрых людей в Иннополисе. 

Они создали этого бота, где вы можете участвовать в обмене замечательными впечатлениями через ваши подарки.  

Вы готовы к чуду?`
}

func (e *rusloc) RulesButtonYes() string {
	return "Да, я готов"
}

func (e *rusloc) RulesMessage() string {
	return `Отлично! Тогда позвольте мне объяснить процесс:

Мои эльфы отправят вам имя случайного человека, для которого вы подготовите подарок 🎉
				
Как это будет работать: 
1️⃣ Вам нужно зарегистрироваться до 11 декабря (включительно)
2️⃣ 11 декабря в 11:00 вы получите имя человека, для которого вы подготовите подарок
3️⃣ Вам нужно подготовить подарок:
Максимальная стоимость подарка - 500 рублей
Вы должны сделать это до 19 декабря (включительно).
Добавьте заметку или маленькую открытку, которая включает имя человека, для которого подарок.

4️⃣ Принесите свой подарок в офис 319

Хотите указать предпочтения о вашем подарке?`
}
