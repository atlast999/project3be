package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/atlast999/project3be/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authorizationMiddleware(tokenMaker token.TokenMaker) func(*gin.Context) {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authorizationHeaderKey)
		if len(header) == 0 {
			err := errors.New("authorization header is missing")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(header)
		if len(fields) < 2 {
			err := errors.New("authorization header is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorType := strings.ToLower(fields[0])
		if authorType != authorizationTypeBearer {
			err := fmt.Errorf("authorization type is not supported: %s", authorType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		fmt.Println(payload)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
