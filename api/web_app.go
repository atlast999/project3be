package api

import (
	"net/http"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/helper"
	"github.com/atlast999/project3be/repository"
	"github.com/gin-gonic/gin"
)

type CreateWebAppRequest struct {
	Name  string `json:"name" binding:"required"`
	Url   string `json:"url" binding:"required"`
	Image string `json:"image"`
}

type WebAppResponse struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Image        string `json:"image"`
	CollectionID *int32 `json:"collection_id"`
}

func newWebAppResponse(webApp db.WebApp) WebAppResponse {
	var collectionId *int32
	if webApp.CollectionID.Valid {
		collectionId = &webApp.CollectionID.Int32
	}
	return WebAppResponse{
		ID:           webApp.ID,
		Name:         webApp.Name,
		Url:          webApp.Url,
		Image:        webApp.Image,
		CollectionID: collectionId,
	}
}

func (server *Server) createWebApp(ctx *gin.Context) {
	var request CreateWebAppRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := server.getAuthenticatedPayload(ctx)
	webApp, err := repository.CreateWebApp(payload.UserID, db.CreateWebAppParams{
		Name:  request.Name,
		Url:   request.Url,
		Image: request.Image,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dataResponse(newWebAppResponse(webApp)))
}

func (server *Server) getMyWebAppList(ctx *gin.Context) {
	var request PagingRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := server.getAuthenticatedPayload(ctx)

	webApps, err := repository.GetMyList(db.GetMyListParams{
		UserID: payload.UserID,
		Offset: (request.Page - 1) * request.Size,
		Limit:  request.Size,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	webAppResponses := helper.Map(webApps, func(webapp db.WebApp) WebAppResponse {
		return newWebAppResponse(webapp)
	})
	ctx.JSON(http.StatusOK, pagingResponse(webAppResponses, request.Page, 1))
}

type ShareMyListRequest struct {
	CollectionName string `json:"collection_name" binding:"required"`
}

func (server *Server) shareMyWebAppList(ctx *gin.Context) {
	var request ShareMyListRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := server.getAuthenticatedPayload(ctx)

	collection, err := repository.ShareMyList(payload.UserID, request.CollectionName, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dataResponse(collection))
}

func (server *Server) getCollectionList(ctx *gin.Context) {
	var request PagingRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	collections, err := repository.GetCollections(db.GetCollectionsParams{
		Offset: (request.Page - 1) * request.Size,
		Limit:  request.Size,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, pagingResponse(collections, request.Page, 1))
}

type CollectionRequest struct {
	CollectionID int32 `uri:"id" binding:"required"`
}

func (server *Server) getWebAppListByCollection(ctx *gin.Context) {
	var pagingRequest PagingRequest
	if err := ctx.ShouldBindQuery(&pagingRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var collectionRequest CollectionRequest
	if err := ctx.ShouldBindUri(&collectionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	webApps, err := repository.GetCollectionDetail(db.GetByCollectionParams{
		Offset:       (pagingRequest.Page - 1) * pagingRequest.Size,
		Limit:        pagingRequest.Size,
		CollectionID: collectionRequest.CollectionID,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	webAppResponses := helper.Map(webApps, func(webapp db.WebApp) WebAppResponse {
		return newWebAppResponse(webapp)
	})
	ctx.JSON(http.StatusOK, pagingResponse(webAppResponses, pagingRequest.Page, 1))
}

func (server *Server) takeCollection(ctx *gin.Context) {
	var collectionRequest CollectionRequest
	if err := ctx.ShouldBindUri(&collectionRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	payload := server.getAuthenticatedPayload(ctx)

	err := repository.TakeCollection(payload.UserID, collectionRequest.CollectionID, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, successResponse())
}
