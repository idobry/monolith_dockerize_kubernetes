package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Index struct {
	Title string
	Body  string

	Links []Link
}

type Link struct {
	URL, Title, Img string
	Count           string
}

type Browser struct {
	Title string
	URL   string
	Icon  string
	Count string
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

var voteTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("votes.tmpl"))

var browsers = map[string]*Browser{
	"firefox": {"firefox",
		"https://ffp4g1ylyit3jdyti1hqcvtb-wpengine.netdna-ssl.com/firefox/files/2017/12/firefox-logo-300x310.png",
		"https://cdn2.iconfinder.com/data/icons/squareplex/512/firefox.png",
		"0"},
	"chrome": {"chrome",
		"https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		"http://flaticns.com/web-icons/google%20chrome.png",
		"0"},
	"explorer": {"explorer",
		"https://www.di.net.au/wp-content/uploads/2009/07/ie-logo.png",
		"http://iconshow.me/media/images/social/flat-style-Metro-UI-Icons/Browser/png/512/MetroUI-Internet-Explorer.png",
		"0"},
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/vote/", voteHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
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
			Count: GetVotes(name),
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func GetVotes(browser string) string {
	// Make a get request
	rs, err := http.Get("http://getVotes:3000/get_votes/" + browser)
	// Process response
	if err != nil {
		panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("browser :", browser, "result :", string(bodyBytes))
	return string(bodyBytes)
}

func InsertVote(browser string) {
	// Make a get request
	rs, err := http.Post("http://insertVote:3000/insert_vote/"+browser, "text/plain", nil)
	// Process response
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("browser :", browser, "result :", string(bodyBytes))
}

// voteHandler is an HTTP handler that serves the image pages.
func voteHandler(w http.ResponseWriter, r *http.Request) {
	browser := strings.TrimPrefix(r.URL.Path, "/vote/")
	InsertVote(browser)
	browsers[browser].Count = GetVotes(browser)
	data, ok := browsers[browser]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := voteTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
