package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/steven7/bookstore_utils-go/rest_errors"
	"github.com/steven7/go-microservices_oauth_api/services/access_token"
	atDomain "github.com/steven7/go-microservices_oauth_api/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
//
//type AccessTokenHandler interface {
//	GetById(*gin.Context)
//	Create(*gin.Context)
//}
//
//type accessTokenHandler struct {
//	service access_token.Service
//}
//
//func NewHandler(service access_token.Service) AccessTokenHandler {
//	return &accessTokenHandler{
//		service: service,
//	}
//}
//
//func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
//	return &accessTokenHandler{
//		service: service,
//	}
//}
//
//func (handler *accessTokenHandler) GetById(c *gin.Context) {
//	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
//	if err != nil {
//		c.JSON(err.Status, err)
//		return
//	}
//	c.JSON(http.StatusOK, accessToken)
//}
//
////func (handler *accessTokenHandler) Create(c *gin.Context) {
////	var at access_token.AccessToken
////	if err := c.BindJSON(&at); err != nil {
////		restErr := rest_errors.NewBadRequestError("invalid json body")
////		c.JSON(restErr.Status, restErr)
////		return
////	}
////
////	if err := handler.service.Create(at); err != nil {
////		c.JSON(err.Status, err)
////		return
////	}
////	c.JSON(http.StatusCreated, at)
////}
//
//func (handler *accessTokenHandler) Create(c *gin.Context) {
//	var request atDomain.AccessTokenRequest
//	if err := c.ShouldBindJSON(&request); err != nil {
//		restErr := rest_errors.NewBadRequestError("invalid json body")
//		c.JSON(restErr.Status, restErr)
//		return
//	}
//
//	accessToken, err := handler.service.Create(request)
//	if err != nil {
//		c.JSON(err.Status, err)
//		return
//	}
//	c.JSON(http.StatusCreated, accessToken)
//}