package models

import (
	"errors"
	"fmt"
	"github.com/Teslenk0/users-management-jr/src/database/mysql"
	"github.com/Teslenk0/utils-go/logger"
	"github.com/Teslenk0/utils-go/rest_errors"
	"strings"
)

//Data access Object
//Interacts with DB

const (
	//QUERIES
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, password, image, date_created) VALUES (?,?,?,?,?,?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, image, date_created FROM users WHERE id=?;"
	queryGetAllUsers            = "SELECT id, first_name, last_name, email, image, date_created FROM users;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, image=? WHERE ID=?;"
	queryDeleteUser             = "DELETE FROM users WHERE ID=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, image, date_created FROM users WHERE email=? AND password=?;"
)

//GetAll - Get all users from database
func (user *User) GetAll() ([]User, *rest_errors.RestError) {
	//Prepares the query
	stmt, err := mysql.Client.Prepare(queryGetAllUsers)

	if err != nil {
		logger.Error("error when trying to prepare the get users statement", err)
		return nil, rest_errors.NewInternalServerError("error when trying to prepare the get users statement", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		logger.Error("error when trying to get all users", err)
		return nil, rest_errors.NewInternalServerError("error when trying to get all user", err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Image, &user.DateCreated); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when trying to get all users", err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError("there are not users created")
	}
	return results, nil

}

//Get - Gets user by ID from DB - act like method
func (user *User) Get() *rest_errors.RestError {

	//Prepares the query
	stmt, err := mysql.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare the get user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare the get user statement", err)
	}

	//Close the stametent when the function returns
	defer stmt.Close()

	//Make a select and looks for only one result
	result := stmt.QueryRow(user.Id)

	//Populates the user given with the data from DB
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Image, &user.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}

		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when trying to get user by id", getErr)

	}

	return nil
}

//Save - saves object to DB - act like method
func (user *User) Save() *rest_errors.RestError {

	//Prepares the statement
	stmt, err := mysql.Client.Prepare(queryInsertUser)
	//Ask if there was errors when attempting for preparing the stmt
	if err != nil {
		logger.Error("error when trying to prepare the save user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare insert user statement", err)
	}
	//Close the connection when the functions returns
	defer stmt.Close()

	//Exec the statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Image, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when saving the user", saveErr)
		return rest_errors.NewInternalServerError("error when trying to save the user", saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
	}
	user.Id = userId

	return nil
}

//Update - updates data from the database with the given one
func (user *User) Update() *rest_errors.RestError {

	stmt, err := mysql.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare the update user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare update user statement", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Image, user.Id)
	if err != nil {
		logger.Error("error when trying to update the user", err)
		return rest_errors.NewInternalServerError("error when trying to update the user", err)
	}
	return nil

}

//Delete - deletes a given user
func (user *User) Delete() *rest_errors.RestError {

	stmt, err := mysql.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare the delete user statement", err)
		return rest_errors.NewInternalServerError("error when trying to prepare the delete user statement", err)
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(user.Id)
	if delErr != nil {
		logger.Error("error when trying to delete the user", delErr)
		return rest_errors.NewInternalServerError("error when trying to delete the user", delErr)
	}
	return nil
}

//FindByEmailAndPassword - function that finds a user by email and password
func (user *User) FindByEmailAndPassword() *rest_errors.RestError {
	stmt, err := mysql.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("error when tying to find user", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Image, &user.DateCreated); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when tying to find user", getErr)
	}
	return nil
}
