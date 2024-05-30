package tests

import (
	"backend/app/models"
	"time"
	"log"
	"fmt"
)

func CreateTodoTest() {
	u, _ := models.GetUser(1)
	date, err := time.Parse("2006-01-02", "2024-06-10")
	if err != nil {
		log.Fatalln("failed to parse date", err)
	}

	fmt.Println("\n ******CreateTodo****** \n")

	u.CreateTodo("first Todo", "University", false, date)
}

func GetTodoTest() {
	fmt.Println("\n ******GetTodo****** \n")

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}

func GetTodosTest() {
	fmt.Println("\n ******GetTodos****** \n")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}
}

func GetTodosByUserTest() {
	fmt.Println("\n ******GetTodosByUser****** \n")

	u1, _ := models.GetUser(1)
	todos, _ := u1.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

}

func UpdateTodoTest() {
	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
}

func DeleteTodoTest() {
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}