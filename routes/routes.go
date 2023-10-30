package routes

import (
	"fmt"
	"gohtmx/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)


func sendErrorEmpty(w http.ResponseWriter) {
	todos, _ := model.GetAllTodos()
	tml := template.Must(template.ParseFiles("templates/index.html"))
	data := model.TodosData{
		Info: "erorr",
		Todos: todos,
	}
	tml.Execute(w, data)
	
}
func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Couldn't get all todos from database", err)
		return
	}
	tml := template.Must(template.ParseFiles("templates/index.html"))
	dt := model.TodosData{
		Info: "loaded",
		Todos: todos,
	}
	err = tml.Execute(w,dt)
	if err != nil {
		fmt.Println("Couldn't execute template", err)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
    dt := model.TodosData{
		Info: "loaded",
		Todos: todos,
	}
	err = tmpl.Execute(w, dt)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Couldn't parse id", err)
	}
	err = model.MarkDone(id)
	if err != nil {
		fmt.Println("Couldn't not update todo", err)
	}
	sendTodos(w)
}
func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
	}
	if strings.Trim(r.FormValue("todo")," ") == "" {
     	sendErrorEmpty(w)
		fmt.Println("Could not create todo empty", "err")
		return
		
	}else {
	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("Could not create todo", err)
	}
	sendTodos(w)
}
	
}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id")
	}
	err = model.Delete(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}
	sendTodos(w)
}
func SetupAndRun() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", createTodo).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
