package api

import (
	"net/http"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/repository"
	"github.com/gin-gonic/gin"
)

type CreateWebAppRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Url    string `json:"url" binding:"required"`
	Image  string `json:"image" binding:"required"`
}

func (server *Server) createWebApp(ctx *gin.Context) {
	var request CreateWebAppRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	webApp, err := repository.CreateWebApp(request.UserID, db.CreateWebAppParams{
		Name:  request.Name,
		Url:   request.Url,
		Image: request.Image,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dataResponse(webApp))
}

