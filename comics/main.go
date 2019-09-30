package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Comics struct {
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

var comic = []Comics{
	{Month: "9", Num: 2208, Link: "", Year: "2019", News: "", Safe_title: "Drone Fishing", Transcript: "", Alt: "Today's consumers who order their drones off the internet don't know the joy of going out in nature and returning with a drone that you caught yourself, whose angry owners you fought off with your own two hands.", Img: "https://imgs.xkcd.com/comics/drone_fishing.png", Title: "Drone Fishing", Day: "27"},
}

var comic2 = []Comics{
	{Month: "9", Num: 2207, Link: "", Year: "2019", News: "", Safe_title: "Math Work", Transcript: "", Alt: "I could type this into a solver, which MIGHT help, but would also mean I have to get a lot of parentheses right...", Img: "https://imgs.xkcd.com/comics/math_work.png", Title: "Math Work", Day: "25"},
}

var comic3 = []Comics{
	{Month: "9", Num: 2206, Link: "https://www.fonts.com/content/learning/fontology/level-3/numbers/oldstyle-figures", Year: "2019", News: "", Safe_title: "Mavis Beacon", Transcript: "", Alt: "There are actually lowercase-like 'oldstyle' forms of normal numbers with more pronounced ascenders and descenders, which is why some numbers like '5' in books sometimes dangle below the line. But the true capital numbers remain the domain of number maven Mavis Beacon.", Img: "https://imgs.xkcd.com/comics/mavis_beacon.png", Title: "Mavis Beacon", Day: "23"},
}

var comic4 = []Comics{
	{Month: "9", Num: 2205, Link: "", Year: "2019", News: "", Safe_title: "Types of Approximation", Transcript: "", Alt: "It's not my fault I haven't had a chance to measure the curvature of this particular universe.", Img: "https://imgs.xkcd.com/comics/types_of_approximation.png", Title: "Types of Approximation", Day: "20"},
}

var comic5 = []Comics{
	{Month: "9", Num: 2204, Link: "", Year: "2019", News: "", Safe_title: "Ksp 2", Transcript: "", Alt: "\"The committee appreciates that your 2020 launch is on track, but the 'human capital/personnel retention' budget includes a lot more unmarked cash payments than usual. What are th--\" \"Public outreach.\"", Img: "https://imgs.xkcd.com/comics/ksp_2.png", Title: "Ksp 2", Day: "18"},
}

func main() {

	var n int
	fmt.Print("Hello, for the latest comic press 1. For last 5 comics enter numbers 2-5 : ")
	fmt.Scanf("%d", &n)
	switch n {
	case 1:
		fmt.Println("Latest comic")
		data, err := json.MarshalIndent(comic, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	case 2:
		fmt.Println("Second to latest comic")
		data, err := json.MarshalIndent(comic2, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

	case 3:
		fmt.Println("Third to latest comic")
		data, err := json.MarshalIndent(comic3, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	case 4:
		fmt.Println("Third to latest comic")
		data, err := json.MarshalIndent(comic4, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	case 5:
		fmt.Println("Third to latest comic")
		data, err := json.MarshalIndent(comic5, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	default:
		fmt.Println("Please enter numbers 1-5")
	}

}
