package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"strings"
)

var (
	r = &tele.ReplyMarkup{}

	// Reply buttons.
	btnHelp        = r.Text("ℹ Help")
	btnContact     = r.Contact("⚙ Contact")
	btnLocation    = r.Location("Location")
	btnPollAny     = r.Poll("PollAny", tele.PollAny)
	btnPollRegular = r.Poll("PollRegular", tele.PollRegular)
	btnPollQuiz    = r.Poll("PollQuiz", tele.PollQuiz)

	// Inline buttons.
	iBtnDataBase  = r.Data("⬅", "prev", "default_data")
	iBtnURL       = r.URL("Visit", "https://google.com")
	iBtnQuery     = r.Query("Search", "query")
	iBtnQueryChat = r.QueryChat("Share", "query")
	//iBtnLogin     = r.Login("Login", &tele.Login{
	//	URL:         "https://google.com",
	//	Text:        "",
	//	Username:    "",
	//	WriteAccess: true})

	simpleMarkup = &tele.ReplyMarkup{}
)

func iBtnData(data ...string) tele.Btn {
	btn := iBtnDataBase
	btn.Data = strings.Join(data, "|")
	return btn
}

func mainMenu(c tele.Context) *tele.ReplyMarkup {
	fmt.Println(c) // just to use context for now
	markup := *simpleMarkup
	markup.Inline(
		markup.Row(iBtnData()),
		markup.Row(iBtnURL),
		markup.Row(iBtnQuery),
		markup.Row(iBtnQueryChat),
		//markup.Row(iBtnLogin),
	)
	return &markup
}
