package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"./utils"
	"github.com/gorilla/mux"
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
		p.db[id] = item
		p.Unlock()
	}

}

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

func (p *DB) GetOneItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	w.WriteHeader(http.StatusOK)
	idd, _ := strconv.Atoi(vars)

	if len(p.db)-1 < idd {
		http.Error(w, "no item of that id", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", p.db[idd])
}

func main() {
	db := &DB{}
	r := mux.NewRouter()
	utils.LoadTemplates("templates/*.html")
	r.HandleFunc("/create", db.Create)
	r.HandleFunc("/read", db.Read)
	r.HandleFunc("/update", db.Update)
	r.HandleFunc("/delete", db.Delete)
	r.HandleFunc("/list", db.htmlList)
	r.HandleFunc("/items/{id}", db.GetOneItem).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
