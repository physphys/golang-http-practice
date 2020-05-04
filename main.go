package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type item struct {
	ID   int
	Name string
	Num  int
}

type todo []item

func index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=todo_owner dbname=todo sslmode=disable")
	checkErr(err)

	rows, err := db.Query("SELECT id, name, num FROM todo")
	checkErr(err)

	var todo todo
	for rows.Next() {
		var id int
		var name string
		var num int
		err = rows.Scan(&id, &name, &num)
		checkErr(err)

		todo = append(todo,
			item{ID: id, Name: name, Num: num})
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
