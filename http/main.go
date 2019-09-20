package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"./utils"
)

type PriceDB struct {
	sync.Mutex
	db map[string]int
}

func (p *PriceDB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "No integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; ok {
		http.Error(w, fmt.Sprintf("%s already exists", item), http.StatusBadRequest)
		return
	}

	p.Lock()
	if p.db == nil {
		p.db = make(map[string]int, 0)
	}
	p.db[item] = price
	p.Unlock()
}

func (p *PriceDB) Update(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	priceStr := r.FormValue("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "No integer price given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	p.db[item] = price
	p.Unlock()
}

func (p *PriceDB) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	delete(p.db, item)
	p.Unlock()
}

func (p *PriceDB) Read(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	}

	if _, ok := p.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%s doesn't exist", item), http.StatusNotFound)
		return
	}

	p.Lock()
	fmt.Fprintf(w, "%s: %d\n", item, p.db[item])
	p.Unlock()
}

// func (p *PriceDB) List(w http.ResponseWriter, r *http.Request) {
// 	p.Lock()
// 	utils.LoadTemplates("templates/*.html")
// 	p.Unlock()
// }

func (p *PriceDB) htmlList(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "index.html", p.db)
}

func main() {
	db := &PriceDB{}
	db.db = make(map[string]int, 0)
	db.db["shoe"] = 100
	utils.LoadTemplates("templates/*.html")
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	//http.HandleFunc("/list", db.List)
	http.HandleFunc("/list", db.htmlList)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
