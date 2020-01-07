package models

import (
	"github.com/Teslenk0/utils-go/rest_errors"
	"strings"
)

//Defines the model
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Image       string `json:"image"`
	DateCreated string `json:"date_created"`
}

//Defines an alias for a user's array
type Users []User

//Validates the passed database
func (user *User) Validate() *rest_errors.RestError {

	//removes spaces from
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.FirstName == "" {
		return rest_errors.NewBadRequestError("user first_name must not be blank")
	}

	if user.LastName == "" {
		return rest_errors.NewBadRequestError("user last_name must not be blank")
	}

	if user.Email == "" {
		return rest_errors.NewBadRequestError("user email must not be blank")
	}

	if user.Password == "" {
		return rest_errors.NewBadRequestError("user password must not be blank")
	}
	return nil
}
