package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cf := NewCmdFlags()
	cf.Execute(&todos)
	storage.Save(todos)
}