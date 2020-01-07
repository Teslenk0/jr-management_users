package app

import (
	"github.com/Teslenk0/users-management-jr/src/controllers/ping"
	"github.com/Teslenk0/users-management-jr/src/controllers/users"
)

func routes(){

	//Endpoint to verify if API is listening
	router.GET("/ping", ping.Ping)

	//Get all the users from the database
	router.GET("/users", users.GetAll)

	//Get a specified user from the database
	router.GET("/users/:user_id", users.Get)
	//----------------------------------------------------------------

	//Creates a new user in the database
	router.POST("/users", users.Create)

	//Login request
	router.POST("/users/login", users.Login)

	//---------------------------------------------------------------------
	//Complete Update
	router.PUT("/users/:user_id", users.Update)

	//Partial Update
	router.PATCH("/users/:user_id", users.Update)

	//--------------------------------------------------------------------
	router.DELETE("/users/:user_id", users.Delete)
}