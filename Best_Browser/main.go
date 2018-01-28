package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	postgres "postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bb"
	password = "bb"
	dbname   = "best_browser"
)

type Index struct {
	Title string
	Body  string

	Links []Link
}

type Link struct {
	URL, Title, Img string
	Count           int
}

type Browser struct {
	Title string
	URL   string
	Icon  string
	Count int
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

var voteTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("votes.tmpl"))

var browsers = map[string]*Browser{
	"firefox": {"FIREFOX",
		"https://ffp4g1ylyit3jdyti1hqcvtb-wpengine.netdna-ssl.com/firefox/files/2017/12/firefox-logo-300x310.png",
		"https://cdn2.iconfinder.com/data/icons/squareplex/512/firefox.png",
		0},
	"chrome": {"CHROME",
		"https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		"http://flaticns.com/web-icons/google%20chrome.png",
		0},
	"explorer": {"EXPLORER",
		"https://www.di.net.au/wp-content/uploads/2009/07/ie-logo.png",
		"http://iconshow.me/media/images/social/flat-style-Metro-UI-Icons/Browser/png/512/MetroUI-Internet-Explorer.png",
		0},
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var p_db, _ = postgres.NewDB("postgres", psqlInfo)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/vote/", voteHandler)
	log.Fatal(http.ListenAndServe("localhost:1234", nil))
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Best Browser",
		Body:  "Vote the best internet browser",
	}
	for name, browser := range browsers {
		data.Links = append(data.Links, Link{
			URL:   "/vote/" + name,
			Title: browser.Title,
			Img:   browser.Icon,
			Count: p_db.GetVotes(name),
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// voteHandler is an HTTP handler that serves the image pages.
func voteHandler(w http.ResponseWriter, r *http.Request) {
	browser := strings.TrimPrefix(r.URL.Path, "/vote/")
	p_db.InsertVote(browser)
	browsers[browser].Count = p_db.GetVotes(browser)
	data, ok := browsers[browser]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := voteTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
