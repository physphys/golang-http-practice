package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

func new(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	switch r.Method {
	case "GET":
		t, _ := template.ParseFiles("./views/new.html")
		t.Execute(w, nil)
	case "POST":
		r.ParseForm()
		num, _ := strconv.Atoi(r.Form["num"][0])
		item := item{Name: r.Form["name"][0], Num: num}

		db, err := sql.Open("postgres", "user=todo_owner dbname=todo sslmode=disable")
		checkErr(err)

		stmt, err := db.Prepare("INSERT INTO todo(name,num) VALUES($1,$2)")
		checkErr(err)

		_, err = stmt.Exec(item.Name, item.Num)
		checkErr(err)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/new", new)
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
