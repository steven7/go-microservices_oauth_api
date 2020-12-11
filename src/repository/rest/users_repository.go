package rest

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/steven7/bookstore_utils-go/rest_errors"
	"github.com/steven7/go-microservices_oauth_api/src/domain/users"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL:        "https://api.bookstore.com",
		Timeout:        100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct {}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(request)
	fmt.Println(bytes)

	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		fmt.Println(response.String())
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}