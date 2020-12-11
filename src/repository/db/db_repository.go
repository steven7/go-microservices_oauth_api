package db

import (
	"github.com/gocql/gocql"
	"github.com/steven7/bookstore_utils-go/rest_errors"
	"github.com/steven7/go-microservices_oauth_api/src/client/cassandra"
	"github.com/steven7/go-microservices_oauth_api/src/domain/access_token"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	updateQueryExpires 	   = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {

}

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (dbr *dbRepository) Create(token access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(updateQueryExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error())
	}
	return nil
}