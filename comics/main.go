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

const allURL = "http://xkcd.com/%d/info.0.json"

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

func getComics(startNum, comicsLength int) ([]Comic, error) {

	comics := make([]Comic, 0, comicsLength)
	for i := startNum; i < startNum+comicsLength; i++ {
		url := fmt.Sprintf(allURL, i)
		comic, err := getComic(url)
		if err != nil {
			return nil, err
		}
		comics = append(comics, *comic)

	}
	return comics, nil
}

func main() {

	lastComic, err := getComic(URL)
	if err != nil {
		log.Fatal(err)
	}

	comicLength := 4
	startComicNum := lastComic.Num - comicLength
	comics, err := getComics(startComicNum, comicLength)
	comics = append(comics, *lastComic)

	for i, com := range comics {
		fmt.Print(i+1, ": ")
		fmt.Printf("%+v\n", com)
	}
}
