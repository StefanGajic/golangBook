package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"./utils"
)

type DB struct {
	sync.Mutex
	db []string
}

func (p *DB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	default:
		p.Lock()
		p.db = append(p.db, item)
		p.Unlock()
	}

}

func (p *DB) Update(w http.ResponseWriter, r *http.Request) {
	///update?id=2&item=zz
	idc := r.FormValue("id")
	id, _ := strconv.Atoi(idc)
	item := r.FormValue("item")

	if len(p.db)-1 < id {
		http.Error(w, "no item of that id", http.StatusBadRequest)
		return
	}

	if idc == "" {
		http.Error(w, "No id given", http.StatusBadRequest)
		return
	}

	switch item {
	case "":
		http.Error(w, "No item given", http.StatusBadRequest)
		return
	default:
		p.Lock()
		//update
		for item, _ := range p.db {
			p.db = append(p.db[:item], p.db[item:]...)

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

	idc := r.FormValue("id")
	id, _ := strconv.Atoi(idc)
	if len(p.db)-1 < id {
		http.Error(w, "no particular id", http.StatusBadRequest)
		return
	}

	switch idc {
	case "":
		http.Error(w, "No id given", http.StatusBadRequest)
		return

	default:
		p.Lock()
		id, _ := strconv.Atoi(idc)
		sl := p.db
		sl = append(sl[0:id], sl[id+1:]...)
		p.db = sl
		p.Unlock()
	}
}

func (p *DB) Read(w http.ResponseWriter, r *http.Request) {
	idc := r.FormValue("id")
	id, _ := strconv.Atoi(idc)
	if len(p.db)-1 < id {
		http.Error(w, "no item of that id", http.StatusBadRequest)
		return
	}

	if idc == "" {
		http.Error(w, "No id given", http.StatusBadRequest)
		return
	}

	p.Lock()
	var item = p.db[id]
	fmt.Fprintf(w, "%s \n", item)
	p.Unlock()

}

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
	log.Fatal(http.ListenAndServe(":8000", nil))
}
