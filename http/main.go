package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"./utils"
)

type DB struct {
	sync.Mutex
	db []string
}

func (p *DB) Create(w http.ResponseWriter, r *http.Request) {
	//id := r.FormValue("id")
	item := r.FormValue("item")
	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	//case "1":
	default:
		p.Lock()
		p.db = append(p.db, item)
		p.Unlock()
	}

}

func (p *DB) Update(w http.ResponseWriter, r *http.Request) {

	//id := r.FormValue("id")
	item := r.FormValue("item")
	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	//case "1":
	default:
		p.Lock()
		//update
		for item, _ := range p.db {
			p.db = append(p.db[:item], p.db[item+1:]...)

		}
		p.Unlock()
	}

}

// 	item := r.FormValue("item")
// 	if item == "" {
// 		http.Error(w, "No item given", http.StatusBadRequest)
// 		return
// 	}

// 	priceStr := r.FormValue("price")
// 	price, err := strconv.Atoi(priceStr)
// 	if err != nil {
// 		http.Error(w, "No integer price given", http.StatusBadRequest)
// 		return
// 	}

// 	if _, ok := p.db[item]; !ok {
// 		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
// 		return
// 	}

// 	p.Lock()
// 	p.db[item] = price
// 	p.Unlock()
// }

func (p *DB) Delete(w http.ResponseWriter, r *http.Request) {

	item := r.FormValue("item")
	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	case item:
		p.Lock()
		//delete(p.db, item)
		//p.db[item] = p.db[len(p.db)-1]
		p.db[len(p.db)-1] = ""
		p.db = p.db[:len(p.db)-1]

		//p.db[len(p.db)-1] = ""
		//p.db = p.db[:len(p.db)-1] // brise poslednji

		//a = append(p.db[:item], p.db[item+1:]...)

		//p.db = append(p.db[:item[0]], p.db[item[0]+1:]...)
		//p.db = append(p.db[:])
		// func remove(db []int, s int) []int {
		// 	return append(db[:s], db[s+1:]...)
		// }

		// func remove(p.db []string, i string) []string {
		// 	copy(slice[i:], slice[i+1:])
		// 	return slice[:len(slice)-1]
		//   }

		// func remove(p.db []int, s int) []int {
		// 	return append(p.db[:s], p.db[s+1:]...)
		// }

		//copy(p.db[i:], p.db[i+1:])
		//p.db[len(p.db)-1] = nil

		//p.db = append(p.db, item)
		p.Unlock()
	}
}

// 	item := r.FormValue("item")
// 	if item == "" {
// 		http.Error(w, "No item given", http.StatusBadRequest)
// 		return
// 	}

// 	if _, ok := p.db[item]; !ok {
// 		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
// 		return
// 	}

// 	p.Lock()
// 	delete(p.db, item)
// 	p.Unlock()
// }

func (p *DB) Read(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	// case read:
	// 	http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
	// 	return
	case item:
		p.Lock()
		fmt.Fprintf(w, "%s \n", item)
		//p.db = append(p.db, item)
		p.Unlock()
	}

}

// 	item := r.FormValue("item")
// 	if item == "" {
// 		http.Error(w, "No item given", http.StatusBadRequest)
// 		return
// 	}

// 	if _, ok := p.db[item]; !ok {
// 		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
// 		return
// 	}

// 	p.Lock()
// 	fmt.Fprintf(w, "%s: %d\n", item, p.db[item])
// 	p.Unlock()
// }

func (p *DB) htmlList(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", p.db)
}

// func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	w.WriteHeader(http.StatusOK)
// 	cat := vars["id"]

// 	fmt.Fprintf(w, "Name: %v\n", cat)
// }

func main() {
	db := &DB{}
	//r := mux.NewRouter()
	//db.db = make(map[string]int, 0)
	//db.db["shoe"] = 100
	utils.LoadTemplates("templates/*.html")
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	//http.HandleFunc("/list", db.List)
	http.HandleFunc("/list", db.htmlList)
	//http.HandleFunc("/list/{id}", ArticlesCategoryHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
