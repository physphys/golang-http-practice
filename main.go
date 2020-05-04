package main

import (
	"html/template"
	"log"
	"net/http"
)

type item struct {
	Name string
	Num  int
}

type todo []item

func index(w http.ResponseWriter, r *http.Request) {
	todo := []item{
		item{Name: "curry", Num: 1},
		item{Name: "milk", Num: 2},
		item{Name: "egg", Num: 20},
	}
	t, _ := template.ParseFiles("./views/index.html")
	t.Execute(w, todo)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
