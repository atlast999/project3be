package api

import (
	"log"
	"net/http"

	"github.com/atlast999/project3be/db/transaction"
	"github.com/atlast999/project3be/helper"
	"github.com/atlast999/project3be/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	txInstance *transaction.TxInstance
	route      *gin.Engine
	config     helper.Config
	tokenMaker token.TokenMaker
}

func NewServer(config helper.Config, txInstance *transaction.TxInstance) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSecretKey)
	if err != nil {
		log.Fatal("Cannot create token maker!")
		return nil, err
	}
	server := &Server{
		txInstance: txInstance,
		config:     config,
		tokenMaker: tokenMaker,
	}
	server.setUpRouter()
	return server, nil
}

func (server *Server) setUpRouter() {
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	route.POST("/signup", server.createUser)
	route.POST("/login", server.loginUser)

	authRoutes := route.Group("/").Use(authorizationMiddleware(server.tokenMaker))
	authRoutes.POST("/web_apps", server.createWebApp)
	authRoutes.GET("/web_apps", server.getMyWebAppList)
	authRoutes.POST("/share_list", server.shareMyWebAppList)
	authRoutes.GET("/collections", server.getCollectionList)
	authRoutes.GET("/collections/:id", server.getWebAppListByCollection)
	authRoutes.PUT("/collections/:id/take", server.takeCollection)
	server.route = route
}

func (server *Server) StartServer() error {
	return server.route.Run()
}

func (server *Server) getAuthenticatedPayload(ctx *gin.Context) *token.Payload {
	return ctx.MustGet(authorizationPayloadKey).(*token.Payload)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func dataResponse(data any) gin.H {
	return gin.H{
		"message": "OK",
		"data":    data,
	}
}

func successResponse() gin.H {
	return gin.H{
		"message": "OK",
	}
}
