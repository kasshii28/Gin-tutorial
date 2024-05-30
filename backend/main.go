package main

import (
	"backend/tests"
	//"backend/app/models/controller"
)

func main() {
	// controller.CorsSettings()
	//controller.StartMainServer()

	tests.CreateUserTest()
	tests.GetUserTest()
	tests.UpdateUserTest()
	tests.DeleteUserTest()
}