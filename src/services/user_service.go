package services

import (
	"github.com/Teslenk0/users-management-jr/src/models"
	"github.com/Teslenk0/utils-go/crypto_utils"
	"github.com/Teslenk0/utils-go/date"
	"github.com/Teslenk0/utils-go/rest_errors"
)

//Interface with methods
type usersServiceInterface interface {
	GetAllUsers() (models.Users, *rest_errors.RestError)
	GetUser(userID int64) (*models.User, *rest_errors.RestError)
	CreateUser(user models.User) (*models.User, *rest_errors.RestError)
	UpdateUser(isPartial bool, user models.User) (*models.User, *rest_errors.RestError)
	DeleteUser(userID int64) *rest_errors.RestError
	LoginUser(request models.LoginRequest) (*models.User, *rest_errors.RestError)
}

//Struct
type usersService struct {
}

//Implementing the interface
var (
	UsersService usersServiceInterface = &usersService{}
)

//GetUsers - this function ask for all users in db
func (s *usersService) GetAllUsers() (models.Users, *rest_errors.RestError) {
	dao := &models.User{}
	return dao.GetAll()
}

//GetUser - this function interacts with DB and gets a user
func (s *usersService) GetUser(userID int64) (*models.User, *rest_errors.RestError) {

	if userID <= 0 {
		return nil, rest_errors.NewBadRequestError("user identifier must be greater than 0")
	}

	var result = &models.User{Id: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

//CreateUser - this function interacts with DAO and creates a user, the error must be at the final of the return statement
func (s *usersService) CreateUser(user models.User) (*models.User, *rest_errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowDBString()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//UpdateUser - this function updates the user with the data given
func (s *usersService) UpdateUser(isPartial bool, user models.User) (*models.User, *rest_errors.RestError) {
	current := &models.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}

		if user.Image != "" {
			current.Image = user.Image
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Image = user.Image
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser - this functions looks for a user and deletes it
func (s *usersService) DeleteUser(userID int64) *rest_errors.RestError {
	dao := &models.User{Id: userID}
	return dao.Delete()
}

//LoginUser - function that validates user and password
func (s *usersService) LoginUser(request models.LoginRequest) (*models.User, *rest_errors.RestError) {
	dao := &models.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}