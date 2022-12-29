package api

import (
	"github.com/atlast999/project3be/db/transaction"
	"github.com/gin-gonic/gin"
)

type Server struct {
	txInstance *transaction.TxInstance
	route *gin.Engine
}

func NewServer(txInstance *transaction.TxInstance) *Server {
	server := &Server{
		txInstance: txInstance,
	}
	route 
}

func Start