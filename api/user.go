package api

import (
	"net/http"

	db "github.com/atlast999/project3be/db/gen"
	"github.com/atlast999/project3be/helper"
	"github.com/atlast999/project3be/repository"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       int32  `json:"id"`
	UserName string `json:"username"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		UserName: user.Username,
	}
}

type AuthenticationResponse struct {
	User  UserResponse `json:"user"`
	Token string       `josn:"token"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var request UserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := helper.GeneratePassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	user, err := repository.CreateUser(db.CreateUserParams{
		Username: request.UserName,
		Password: hashedPassword,
	}, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	token, err := server.tokenMaker.CreateToken(user.ID, request.UserName, server.config.TokenExpiredDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, dataResponse(AuthenticationResponse{
		User:  newUserResponse(user),
		Token: token,
	}))
}

func (server *Server) loginUser(ctx *gin.Context) {
	var request UserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := repository.GetUser(request.UserName, server.txInstance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = helper.CheckPassword(request.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}
	token, err := server.tokenMaker.CreateToken(user.ID, request.UserName, server.config.TokenExpiredDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, dataResponse(AuthenticationResponse{
		User:  newUserResponse(user),
		Token: token,
	}))
}
