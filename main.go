package main

func main() {
	todos := Todos{}

	todos.add("Add some features")
	todos.add("Finish the project")
	todos.toggle(0)
	todos.print()

	todos.delete(0)
	todos.print()
}