package main

import "fmt"

func main() {
	todos := Todos{}
	todos.addTask("Buy notebook")
	todos.addTask("Get visa")
	todos.addTask("Watch TI")

	fmt.Printf("%+v\n\n", todos)

	todos.removeTask(2)

	fmt.Printf("%+v\n\n", todos)
}
