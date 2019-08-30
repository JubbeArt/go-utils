package fchan

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"strings"
)

type Post struct {
	Board  string
	Number int64 `json:"no"`

	Subject string `json:"sub"`
	Comment string `json:"com"`
	Name    string `json:"name"`

	Time      int64 `json:"time"`
	TimeMilli int64 `json:"tim"`

	FileName  string `json:"filename"`
	Extension string `json:"ext"`
	MD5       string `json:"md5"`

	Replies int `json:"replies"`
	Images  int `json:"images"`
}

func (p *Post) Text() string {
	if p.Subject == "" {
		return p.Comment
	} else if p.Comment == "" {
		return p.Subject
	}

	return p.Subject + " " + p.Comment
}

func (p *Post) FullFileName() string {
	return p.FileName + p.Extension
}

func (p *Post) HasImage() bool {
	return p.Extension != ""
}

func (p *Post) ImageUrl() string {
	return fmt.Sprintf("https://i.4cdn.org/%v/%v%v", p.Board, p.TimeMilli, p.Extension)
}

func (p *Post) ThumbnailUrl() string {
	return fmt.Sprintf("https://i.4cdn.org/%v/%vs.jpg", p.Board, p.TimeMilli)
}

func Thread(board string, number int64) ([]Post, error) {
	var threadJSON struct {
		Posts []Post `json:"posts"`
	}

	err := getJson(fmt.Sprintf("https://a.4cdn.org/%v/thread/%v.json", board, number), &threadJSON)

	if err != nil {
		return nil, err
	}

	posts := threadJSON.Posts

	for i := range posts {
		posts[i].Board = board
		posts[i].Subject = cleanString(posts[i].Subject)
		posts[i].Comment = cleanString(posts[i].Comment)

	}

	return posts, nil
}

func Catalog(board string) ([]Post, error) {
	threads := make([]Post, 0)

	type Page struct {
		Threads []Post `json:"threads"`
	}

	var pages []Page

	err := getJson(fmt.Sprintf("https://a.4cdn.org/%v/catalog.json", board), &pages)

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		threads = append(threads, page.Threads...)
	}

	for i := range threads {
		threads[i].Board = board
		threads[i].Subject = cleanString(threads[i].Subject)
		threads[i].Comment = cleanString(threads[i].Comment)
	}

	return threads, nil
}

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return errors.New(fmt.Sprintf("url: %q responded with status code: %v", url, resp.StatusCode))
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&target)

	if err != nil {
		return err
	}

	return nil
}

func cleanString(text string) string {
	text = html.UnescapeString(text)
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "<br>", " ")
	text = strings.ReplaceAll(text, "<wbr>", "")
	text = strings.ReplaceAll(text, "<span class=\"quote\">", "")
	text = strings.ReplaceAll(text, "</span>", "")
	return text
}
