package rest

import (
	"encoding/json"
	"fmt"
	"go-microservces_users_api/utils/errors"
	"go-microservices_oath_api/src/domain/users"
	"github.com/golang-restclient/rest"
	"time"

	//"github.com/merc
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL:        "https://api.bookstore.com",
		Timeout:        100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(request)
	fmt.Println(bytes)

	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}