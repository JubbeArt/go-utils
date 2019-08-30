package fchan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Post struct {
	Board  string
	Number int64 `json:"no"`

	Comment string `json:"com"`
	Name    string `json:"name"`

	Time        int64 `json:"time"`
	TimeInMilli int64 `json:"tim"`

	FileName  string `json:"filename"`
	Extension string `json:"ext"`
}

type Thread struct {
	*Post

	Subject string `json:"sub"`

	Replies int `json:"replies"`
	Images  int `json:"images"`
}

func (p *Post) FullFileName() string {
	return p.FileName + p.Extension
}

func (p *Post) HasImage() bool {
	return p.Extension != ""
}

func (p *Post) ImageUrl() string {
	return fmt.Sprintf("https://i.4cdn.org/%v/%v%v", p.Board, p.TimeInMilli, p.Extension)
}

func (p *Post) ThumbnailUrl() string {
	return fmt.Sprintf("https://i.4cdn.org/%v/%vs.jpg", p.Board, p.TimeInMilli)
}

func GetThread(board string, number int64) ([]Post, error) {
	posts := make([]Post, 0)

	err := getJson(fmt.Sprintf("https://a.4cdn.org/%v/thread/%v.json", board, number), &posts)

	if err != nil {
		return nil, err
	}

	for i := range posts {
		posts[i].Board = board
	}

	return posts, nil
}

func Catalog(board string) ([]Thread, error) {
	threads := make([]Thread, 0)

	type Page struct {
		Threads []Thread `json:"threads"`
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
