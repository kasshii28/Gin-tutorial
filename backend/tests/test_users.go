package tests

import (
	"fmt"
	"backend/app/models"
)

func CreateUserTest() {
	fmt.Println("\n ******CreateUser****** \n")
	u := &models.User{}
	
	u.Name = "testUser"
	u.Email = "test@example.com"
	u.PassWord= "test1"
	fmt.Println(u)

	u.CreateUser()
}

func GetUserTest() {
	fmt.Println("\n ******GetUser****** \n")
	u, _ := models.GetUser(1)
	fmt.Println(u)
}

func UpdateUserTest() {
	fmt.Println("\n ******UpdateUser****** \n")
	u, _ := models.GetUser(1)

	u.Name = "test2"
	u.Email  = "test2@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}

func DeleteUserTest() {
	fmt.Println("\n ******DeleteUser****** \n")
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}