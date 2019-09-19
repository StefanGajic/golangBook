package main

import (
	"fmt"
	"sort"
	"time"

	a "./animals"
)

func main() {

	start := time.Now()

	animalGroup := []*a.Animals{
		{"Kangaroo", 3},
		{"Wombat", 5},
		{"Koala", 1},
		{"Tasmanian devil", 7},
	}

	fmt.Println(animalGroup)
	fmt.Println(animalGroup[2])
	sort.Sort(a.SortByAge(animalGroup))
	fmt.Println(animalGroup)

	fmt.Println(time.Since(start))
}
