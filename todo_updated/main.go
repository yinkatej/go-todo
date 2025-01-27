package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Todo struct{
	Id string `json:"id"`
	Item string `json:"item"`
	Checked bool `json:"checked"`
}

var todos []Todo

func main(){
	route := mux.NewRouter()

	route.HandleFunc("/todos", getTodos).Methods("GET")
	route.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	route.HandleFunc("/add_todo", addTodo).Methods("POST")
	route.HandleFunc("/update/{id}", updateTodo).Methods("PUT")
	route.HandleFunc("/delete/{id}", deleteTodo).Methods("DELETE")
	route.HandleFunc("/checkall", checkAll).Methods("GET")

	fmt.Printf("Starting server at port 8080\n ")
	log.Fatal(http.ListenAndServe(":8080", route))
}


func getTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
func getTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range todos{
		
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
func addTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo) 
	todo.Id = strconv.Itoa(len(todos))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
	

}
func deleteTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range todos{
		if param["id"] == item.Id{
			todos = append(todos[:index], todos[index+1:]...)
			json.NewEncoder(w).Encode(todos)
			break
		}
	}


}
func updateTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	param := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&todo)
	for index, item := range todos{
		if param["id"] == item.Id{
			todos = append(todos[:index], todos[index+1:]...)
			todo.Id = param["id"]
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todos)
		}
		

	}
}
func checkAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for index, item := range todos{
		item.Checked = true
		todos = append(todos[:index], todos[index+1:]...)
		
		todos = append(todos, item)
		
	}
	json.NewEncoder(w).Encode(todos)
}

