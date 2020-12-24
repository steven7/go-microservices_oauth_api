package access_token

import (
	"github.com/steven7/bookstore_utils-go/rest_errors"
	"github.com/steven7/go-microservices_oauth_api/repository/db"
	"github.com/steven7/go-microservices_oauth_api/repository/rest"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
	//repository Repository
}

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

//func NewService(repo ...Repository) Service {
//	return &service{
//		repository: repo,
//	}
//}

func (s *service) GetById(accessTokenId string) (*AccessToken, *rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}