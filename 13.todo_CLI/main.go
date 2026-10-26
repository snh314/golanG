package main

import "fmt"

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos []Todo

func displayMenu() {
	fmt.Println("\n=== My Todo List ===")
	fmt.Println("1. Add - Add new Todo in the list")
	fmt.Println("2. View - see full Todo list")
	fmt.Println("3. Complete - complete Todo list")
	fmt.Println("4. Delete - Erase Todo")
	fmt.Println("5. Exit - Close program")

}
func addTodo() {
	fmt.Print("Add Todo Title: ")

	var title string
	fmt.Scanln(&title)

	newTodo := Todo{len(todos) + 1, title, false}

	todos = append(todos, newTodo)

	fmt.Printf("✓ '%s' is added! \n", title)
}
func viewTodos() {
	fmt.Println("\n=== Your Todos ===")
	if len(todos) == 0 {
		fmt.Println("📭 No todos yet!")
		return
	}

	for _, todo := range todos {
		status := "[ ]" // ← Default: সম্পন্ন হয়নি

		if todo.Completed == true {
			status = "[✓]" // সম্পন্ন হয়েছে
		}

		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Title)
	}
}

func completeTodo() {
	viewTodos()
	var id int
	fmt.Print("Which ID you want to Complete ? ")
	fmt.Scanln(&id)

	for index, todo := range todos {
		if todo.ID == id {
			todos[index].Completed = true
			fmt.Printf("✓ '%s' Task Completed!\n", todo.Title)
			return // function শেষ করো
		}
	}
	fmt.Println("❌ Your inputed ID is not available")

}

func deleteTodo() {
	viewTodos()
	var id int
	fmt.Print("Which ID you want to delete? ")
	fmt.Scanln(&id)

	for index, todo := range todos {
		if todo.ID == id {
			deletedTitle := todo.Title
			todos = append(todos[:index], todos[index+1:]...)
			fmt.Printf("✓ '%s' মুছা হয়েছে!\n", deletedTitle)
			return
		}
	}
	fmt.Println("❌ The list is deleted")
}

func main() {

	for {
		displayMenu()
		fmt.Print("\n Select an Option (1-5): ")

		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			addTodo()
		} else if choice == 2 {
			viewTodos()
		} else if choice == 3 {
			completeTodo()
		} else if choice == 4 {
			deleteTodo()
		} else if choice == 5 {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid choice!")
		}
	}

}
