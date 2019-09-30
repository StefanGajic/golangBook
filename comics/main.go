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
	Safe_title string
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
		log.Fatal(err)
		return nil, err
	}
	return comic, nil
}

func getComic(url string) (*Comic, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &comic, nil
}

func main() {

	comic, err := getLastComic()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(comic)
}
