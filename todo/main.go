package main

import (
	"fmt"
	"strings"
)
var todoMap = [] string{}
func main(){
	var action string
	
	var input string
	for {
		fmt.Printf("Hi! Welcome to the New Todo App")
	fmt.Printf(`
	To view all list items Enter "V"
	To add items Enter "A"
	To delete an item Enter "D"
	To update an item Enter "U"\n
	`)
	fmt.Scan(&action)
	switch strings.ToUpper(action){
	case "V":
		getAll(todoMap)
	
	case "A":
		
		fmt.Printf("Enter todo item\n")
		fmt.Scan(&input)
		add(input)

	case "D":
		var index int
		fmt.Printf("Enter the index of item to delete\n")
		fmt.Scan(&index)
		delete(index)
		
	case "U":
		var index int
		var element string
		fmt.Printf("Enter the index of item to update\n")
		fmt.Scan(&index)
		fmt.Printf("Enter the item to update\n")
		fmt.Scan(element)
		update(index, element)

	default:
		fmt.Printf("You did not specify a valid input\n")
		getAll(todoMap)
	}
	fmt.Printf("Enter C to repeat\n")
	var op string
	fmt.Scan(&op)
	if op != "c" && op != "C"{
		break
	}
	}
	

}

func getAll(items []string) {
	if len(items) ==0 {
		fmt.Printf("Ther are no item in tod list\n")
	}else{
		fmt.Printf("Current items in todo list\nNo\tItem\n")
		
		for index, v := range items{
			fmt.Printf("%v\t%v\n", index+1, v)
			index += 1
		}
		
	}
	 
}
func getOne(index int, items []string){
	if items ==nil{
		fmt.Printf("Ther are no item in todo list")
	}else{
		fmt.Printf("No\tItem\n%v\t%v", index, items[index-1])
		
	}
}

func add( element string, ){
	todoMap = append(todoMap, element)
	getAll(todoMap)
}
func delete(index int){
	var newItem = []string{}
	for i, v := range todoMap {
		if index == i+1{
			continue
		}
		newItem = append(newItem, v)
	}
	todoMap = newItem
	getAll(todoMap)
}
func update(index int, element string){
	todoMap[index] = element
	getAll(todoMap)
}