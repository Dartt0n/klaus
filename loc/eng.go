package loc

import "fmt"

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
	return `And of course we didn't forget about the staff! Don't forget to grab a present on the way home!`
}

func (e *engloc) GiftForMessage(name, alias string, prefs []string) string {
	prefs_str := ""
	for _, pref := range prefs {
		prefs_str += " • " + pref + "\n"
	}

	return fmt.Sprintf(`And... the moment of thuth... You are Secret Santa for...

	%s (@%s)!!!
	His/her preferences are: 
	%s
	
	Remember, you should prepare a present before the 21th of Decemeber and bring it to the 319 office.
	Don't forget to sign whom the gift is to!`, name, alias, prefs_str)
}

func (e *engloc) RegistrationClosed() string {
	return `🎅 Ho Ho Ho! The sleigh is full, and Santa's workshop is bustling with preparations for the big day! We're sorry to say that the registration for this year's Secret Santa has officially closed. But fear not, for the spirit of giving and joy is still very much alive! Keep spreading the holiday cheer, and who knows, maybe a little surprise might find its way to you too! 🎅`
}

func (e *engloc) InfoGiftMessage(name, alias string, prefs []string) string {
	prefs_str := ""
	for _, pref := range prefs {
		prefs_str += " • " + pref + "\n"
	}

	return fmt.Sprintf(`Ho ho ho! 🎁

Remember, what you should do:
1️⃣ You need to prepare present:
The maximum value of the gift is 300 rubles
You need to make a gift before the 21th of December (inclusive)
Don't forget to attach a card with nice words, and sign whom the gift is to

2️⃣ Then gift is ready, bring it to the 319 office

3️⃣ If you are leaving early, you can take your gift in the 319 office (please contact the administrator beforehand)

4️⃣ If you are not, be ready to have fun and receive your gift at 🎉New Year Party🎉 (it will take place on the 22th December)


You are Secret Santa for
%s (@%s)!!!
His/her preferences are: 
%s`, name, alias, prefs_str)
}
