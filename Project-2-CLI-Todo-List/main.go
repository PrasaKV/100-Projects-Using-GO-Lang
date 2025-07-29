package main

import (
	"fmt"
)

type Task struct {
	title       string
	description string
}

func main() {

	var tasks []Task

	var option int

loop:
	for {

		fmt.Println("\nWelcome !")
		fmt.Println(`
      1 - Add A Task
      2 - List Tasks
      3 - Delete Task
      4 - Exit
      `)
		fmt.Println("Enter Your Option")
		fmt.Scan(&option)

		switch option {
		case 1:
			add(&tasks)
		case 2:
			list(&tasks)
		case 3:
			remove(&tasks)
		case 4:
			fmt.Println("See You Soon")
			break loop
		}

	}

}

func add(tasks *[]Task) {

	var title string
	var description string

	fmt.Println("Task Title :")
	fmt.Scan(&title)

	fmt.Println("Task Description :")
	fmt.Scan(&description)

	*tasks = append(*tasks, Task{title: title, description: description})

}

func list(tasks *[]Task) {

	fmt.Println(" Tasks Are - ")
	for _, value := range *tasks {
		fmt.Printf("- %s  -- %s", value.title, value.description)
	}
}
func remove(tasks *[]Task) {
	var id int

	fmt.Println(" Tasks Are - ")
	for i, value := range *tasks {
		fmt.Printf(" %d - %s  -- %s", i, value.title, value.description)
	}

	fmt.Println("\nEnter Id of The Task to Remove -")
	fmt.Scan(&id)

	*tasks = append((*tasks)[:id], (*tasks)[id+1:]...)

	fmt.Println(" Tasks Are - ")
	for i, value := range *tasks {
		fmt.Printf(" %d - %s  -- %s", i, value.title, value.description)
	}
}
