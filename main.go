package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	todos.add("Add some features")
	todos.add("Finish the project")
	todos.toggle(0)
	todos.print()

	storage.Save(todos)
}