package app

import (
	"github.com/gin-gonic/gin"
	"github.com/steven7/go-microservices_oauth_api/services/access_token"
	http "github.com/steven7/go-microservices_oauth_api/http/access_token"
	"github.com/steven7/go-microservices_oauth_api/repository/db"
	"github.com/steven7/go-microservices_oauth_api/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.GET("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}