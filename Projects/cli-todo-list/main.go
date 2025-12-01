package main

func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Write Code")
	todos.add("Wakeup early")
	// fmt.Println(todos)
	todos.print()
	todos.toggle(1)
	// fmt.Println(todos)

	// todos.delete(0)
	todos.print()
}