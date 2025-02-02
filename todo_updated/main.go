package main

import (
	"encoding/json" // for converting variable to and from json format
	"fmt"
	"log"
	"net/http" // web request and response
	"strconv" 
	"github.com/gorilla/mux" // for getting the path variable
)
// struct to represent each item as json format
type Todo struct{
	Id string `json:"id"`
	Item string `json:"item"`
	Checked bool `json:"checked"`
}

var todos []Todo

func main(){
	route := mux.NewRouter()
	// the routes for performing the crud operations
	route.HandleFunc("/todos", getTodos).Methods("GET")
	route.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	route.HandleFunc("/add_todo", addTodo).Methods("POST")
	route.HandleFunc("/update/{id}", updateTodo).Methods("PUT")
	route.HandleFunc("/delete/{id}", deleteTodo).Methods("DELETE")
	route.HandleFunc("/checkall", checkAll).Methods("GET")

	fmt.Printf("Starting server at port 8080\n ")
	log.Fatal(http.ListenAndServe(":8080", route))// starting a server
}

// function for getting all items in the list
func getTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_Type", "application/json")
	json.NewEncoder(w).Encode(todos) // encodes the todo slice into a json
}
// function to get a single item in the slice based on index
func getTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) //Get the path variable eg /id
	
	for _, item := range todos{
		
		if item.Id == param["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
// function to add item to the slice
func addTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_Type", "application/json")// set the header
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo) // converts the json into a  variable
	todo.Id = strconv.Itoa(len(todos))
	todos = append(todos, todo)// add the new item to the slice
	json.NewEncoder(w).Encode(todo) //return the slice as json
	

}
// Remove an item from the slice
func deleteTodo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // get the path variable
	for index, item := range todos{
		if param["id"] == item.Id{
			todos = append(todos[:index], todos[index+1:]...)// using append to omit an item in the slice based on index
			json.NewEncoder(w).Encode(todos) // encode it back as json
			break
		}
	}


}
// update an item in the slice base on index
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
// this function marks the checked field of all items in the slice as true
func checkAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	for index, item := range todos{
		item.Checked = true
		todos = append(todos[:index], todos[index+1:]...)
		
		todos = append(todos, item)
		
	}
	json.NewEncoder(w).Encode(todos)
}

