package main

import tele "gopkg.in/telebot.v3"

func IncludeRoutes(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("menu", mainMenu(c))
	})
	b.Handle("/rbs", func(c tele.Context) error {
		markup := *simpleMarkup
		markup.Reply(
			markup.Row(btnHelp),
			markup.Row(btnContact),
			markup.Row(btnLocation),
			markup.Row(btnPollAny),
			markup.Row(btnPollRegular),
			markup.Row(btnPollQuiz),
		)
		return c.Send("Hello!", &markup)
	})
	b.Handle(&iBtnDataBase, func(c tele.Context) error {
		markup := *simpleMarkup
		markup.Inline(
			markup.Row(iBtnURL, iBtnQuery),
			markup.Row(iBtnData()),
			markup.Row(iBtnQueryChat),
			//markup.Row(iBtnLogin),
		)
		_, err := b.EditReplyMarkup(c.Message(), &markup)
		return err
	})
	b.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("help!")
	})
	b.Handle(tele.OnContact, func(c tele.Context) error {
		return c.EditOrSend("this is re result of sending you contact!")
	})
}
