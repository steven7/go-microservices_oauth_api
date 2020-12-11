package app

import (
	"github.com/gin-gonic/gin"
	"go-microservices_oath_api/src/client/cassandra"
	"go-microservices_oath_api/src/domain/access_token"
	http "go-microservices_oath_api/src/http/access_token"
	"go-microservices_oath_api/src/repository/db"
	"go-microservices_oath_api/src/repository/rest"
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