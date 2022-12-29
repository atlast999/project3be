package api

import (
	"net/http"

	"github.com/atlast999/project3be/db/transaction"
	"github.com/gin-gonic/gin"
)

type Server struct {
	txInstance *transaction.TxInstance
	route      *gin.Engine
}

func NewServer(txInstance *transaction.TxInstance) *Server {
	server := &Server{
		txInstance: txInstance,
	}
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	route.POST("/signup", server.createUser)
	route.POST("/login", server.loginUser)
	route.POST("/web_app", server.createWebApp)
	server.route = route
	return server
}

func (server *Server) StartServer() error {
	return server.route.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}

func dataResponse(data any) gin.H {
	return gin.H{
		"data": data,
	}
}
