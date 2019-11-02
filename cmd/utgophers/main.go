package main

import (
	"fmt"
	"os"

	"github.com/pkg/browser"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app  = kingpin.New("utgophers", "A UT Gophers Help App")
	open = app.Flag("open", "Automatically open the browser to linked sites").
		Default("false").Bool()

	slack = app.Command("slack", "Get Help Joining Slack")

	meetup = app.Command("meetup", "Get Help Joining Meetup.com")

	home    = app.Command("home", "Go To Homepage")
	subPage = home.Flag("sub-page", "Sub page to open").
		PlaceHolder("PAGENAME").Short('s').
		Enum("about", "companies", "presentations", "resources")
)

const (
	homeURL   = "http://utahgolang.com/"
	slackURL  = "http://bit.ly/forgeutahinvite)!"
	meetupURL = "https://www.meetup.com/utahgophers)!"
)

func main() {
	app.HelpFlag.Short('h')

	var msg string
	var url string

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Home Page
	case home.FullCommand():
		msg = "Our home is"
		if subPage == nil {
			url = homeURL
		} else {
			url = fmt.Sprintf("%s%s", homeURL, *subPage)
		}
	// Slack help
	case slack.FullCommand():
		msg = "Join us on Slack"
		url = slackURL
	// Meetup.com
	case meetup.FullCommand():
		msg = "Join the meetup"
		url = meetupURL
	}

	fmt.Printf("%s at %s\n", msg, url)
	if *open {
		browser.OpenURL(url)
	}
}
