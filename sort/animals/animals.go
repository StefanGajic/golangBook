package animals

import "fmt"

type Animals struct {
	Species string
	Age     int
}

func (a Animals) String() string {
	return fmt.Sprintf("%s: %d", a.Species, a.Age)
}

type SortByAge []*Animals

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
