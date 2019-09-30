package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const URL = "https://xkcd.com/info.0.json"

func getLastComic() (*Comic, error) {
	comic, err := getComic(URL)
	if err != nil {
		return nil, err
	}
	return comic, nil
}

func getComic(url string) (*Comic, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return nil, err
	}
	return &comic, nil
}

func getLastFiveComics(comic *Comic, comicSlice *[]Comic) error {

	for i := comic.Num - 1; i > comic.Num-5; i-- {
		url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", i)
		temporaryComic, err := getComic(url)
		if err != nil {
			return err
		}
		*comicSlice = append(*comicSlice, *temporaryComic)
	}
	return nil
}

func (c Comic) String() string {
	return fmt.Sprintf("Month : %s\tNum : %d\tLink : %s\tYear : %s\tNews : %s\tTransscript : %s\tAlt: %s\tImg : %s\tTitle : %s\tDay : %s\t", c.Month, c.Num, c.Link, c.Year, c.News, c.Transcript, c.Alt, c.Img, c.Title, c.Day)
}

func main() {

	comic, err := getLastComic()
	if err != nil {
		log.Fatal(err)
	}
	comics := []Comic{}
	comics = append(comics, *comic)

	err1 := getLastFiveComics(comic, &comics)
	if err1 != nil {
		log.Fatal(err1)
	}

	for i, com := range comics {
		fmt.Print(i+1, ": ")
		fmt.Println(com.String())
	}
}
