package main

import "fmt"

func main() {
	todos := Todos{}
	
	todos.add("Add some features")
	todos.add("Finish the project")
	fmt.Printf("%+v\n\n", todos)

	todos.delete(0)
	fmt.Printf("%+v", todos)
}