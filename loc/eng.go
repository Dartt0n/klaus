package loc

var (
	_   Localization = (*engloc)(nil)
	ENG              = engloc{"en"}
)

type engloc struct {
	Lang string `json:"lang,omitempty"`
}

func (e *engloc) UnexpectedMessageText() string {
	return "Please, use keyboard buttons :)"
}

func (e *engloc) EnterPrefReplySuccess() string {
	return "Awesome! Your current list of preferences is:\n\n"
}

func (e *engloc) EnterPrefReplyRemove() string {
	return "Ok! Your current list of preferences is:\n\n"
}

func (e *engloc) EnterPrefReplyNext() string {
	return `Nice! Enter next preference:`
}

func (e *engloc) EnterPrefReplyUnkwonMessage() string {
	return "Sorry, I can not understand this message"
}

func (e *engloc) EnterPrefReplyZeroPref() string {
	return "Please, enter atleast on preference"
}

func (e *engloc) EnterPrefReplyFinish() string {
	return `Cool! We made all preparations!   

Now you should wait for the 11th of December! That day, at 11:00 AM, I'll send you the name of the person you're preparing a gift for! Good luck, my friend!`
}

func (e *engloc) EnterPrefButtonContinue() string {
	return "I want to enter one more preference"
}

func (e *engloc) EnterPrefButtonEnd() string {
	return "That's all"
}

func (e *engloc) EnterPrefButtonRemove() string {
	return "Remove last preference"
}

func (e *engloc) PrefIntroMessage() string {
	return `Ho ho ho! 
Probably, the person who would receive your name will not know about you anything. 

You need to help him/her to prepare a good gift for you!

So, tell me your 1st preference:`
}

func (e *engloc) StartButtonYes() string {
	return "Let's go!"
}

func (e *engloc) StartMessage() string {
	return `Hello, my friend! It's me, Santa! And I'm glad to see you here! 🎅

My clever elves decided to help me with presents for kind people in Innopolis University. 

They created this bot where you can participate in sharing wonderful vibes through your gifts.  

Are you ready for a miracle?`
}

func (e *engloc) RulesButtonYes() string {
	return "Yes, I'm ready"
}

func (e *engloc) RulesMessage() string {
	return `Great! Then, let me explain the process:

My elves will send you the name of a random person for whom you will prepare a present 🎉
				
How it will work: 
1️⃣ You need to register before the 11th of December (inclusively)
2️⃣ 11th of December at 11:00 AM you'll receive the name of the person you're preparing a gift for
3️⃣ You need to prepare present:
The maximum value of the gift is 500 rubles
You need to make it before 22th of December (inclusive).
Add the note or little postcard, which includes the name of the person for whom the present is.

4️⃣ Bring your gift to the 319 office

Want to specify preferences about your present?`
}

func (e *engloc) StartupMessage() string {
	return `Ho ho ho! 🎁

Attention, my dear friend! Tomorrow, after 11:00 AM, you'll receive the name of the person you're preparing a gift for! But wait, there's still time for you to tell us your preferences. This is your last chance to ensure your gift will bring joy to its recipient! ⏰

Remember, in order to receive the name of the person you're preparing a gift for, you must add at least one preference. If you haven't done so yet, don't miss this opportunity to make your gift extra special!

Can't wait to enter new preferences or register again? Simply type /start to begin the process again and share your thoughtful preferences. Let the magic of giving continue! 🎅`
}
