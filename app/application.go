package app

import (
	"github.com/gin-gonic/gin"
	"github.com/steven7/go-microservices_oauth_api/client/cassandra"
	"github.com/steven7/go-microservices_oauth_api/domain/access_token"
	http "github.com/steven7/go-microservices_oauth_api/http/access_token"
	"github.com/steven7/go-microservices_oauth_api/repository/db"
	"github.com/steven7/go-microservices_oauth_api/repository/rest"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//session, dbErr := cassandra.GetSession()
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	session := cassandra.GetSession()
	session.Close()

	atHandler := http.NewHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.GET("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}