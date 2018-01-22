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
	user     = "idob"
	password = "idob"
	dbname   = "best_browser"
)

type Index struct {
	Title string
	Body  string

	Links []Link
}

type Link struct {
	URL, Title, Img string
}

type Image struct {
	Title string
	URL   string
	Icon  string
	Count int
}

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

var imageTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("image.tmpl"))

var images = map[string]*Image{
	"firefox": {"FIREFOX",
		"https://ffp4g1ylyit3jdyti1hqcvtb-wpengine.netdna-ssl.com/firefox/files/2017/12/firefox-logo-300x310.png",
		"https://cdn2.iconfinder.com/data/icons/squareplex/512/firefox.png",
		0},
	"chrom": {"CHROM",
		"https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		"http://flaticns.com/web-icons/google%20chrome.png",
		0},
	"explorer": {"EXPLORER",
		"https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		"http://iconshow.me/media/images/social/flat-style-Metro-UI-Icons/Browser/png/512/MetroUI-Internet-Explorer.png",
		0},
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var p_db, _ = postgres.NewDB("postgres", psqlInfo)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image/", imageHandler)
	log.Fatal(http.ListenAndServe("localhost:1234", nil))
}

// indexHandler is an HTTP handler that serves the index page.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := &Index{
		Title: "Best Browser",
		Body:  "Vote the best internet browser",
	}
	for name, img := range images {
		data.Links = append(data.Links, Link{
			URL:   "/image/" + name,
			Title: img.Title,
			Img:   img.Icon,
		})
	}
	if err := indexTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

// imageHandler is an HTTP handler that serves the image pages.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	browser := strings.TrimPrefix(r.URL.Path, "/image/")
	p_db.InsertVote(browser)
	images[browser].Count = p_db.GetVotes(browser)
	data, ok := images[browser]
	if !ok {
		http.NotFound(w, r)
		return
	}
	if err := imageTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}
