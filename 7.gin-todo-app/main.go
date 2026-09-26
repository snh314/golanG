package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Books", Completed: false},
	{ID: "3", Item: "Completed Video", Completed: false},
}

func getTodos(cont *gin.Context) {
	cont.IndentedJSON(http.StatusOK, todos)
}

func addTodo(cont *gin.Context) {
	var newTodo todo

	if err := cont.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)

	cont.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoByID(id string) (*todo, error) {
	for index, value := range todos {
		if value.ID == id {
			return &todos[index], nil
		}
	}
	return nil, errors.New("not found")

}
func getTodo(cont *gin.Context) {
	id := cont.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		cont.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO not found"})
		return
	}
	cont.IndentedJSON(http.StatusOK, todo)
}
func toggleTodoStatus(cont *gin.Context) {
	id := cont.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		cont.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	todo.Completed = !todo.Completed
	cont.IndentedJSON(http.StatusOK, todo)

}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
