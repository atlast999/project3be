package api

import "github.com/gin-gonic/gin"

type PagingRequest struct {
	Page int32 `form:"page" binding:"required,min=1"`
	Size int32 `form:"size" binding:"required,min=0"`
}

type PagingResponse struct {
	Items any        `json:"items"`
	Meta  PagingMeta `json:"meta"`
}

type PagingMeta struct {
	Page  int32 `json:"page"`
	Total int32 `json:"total"`
}

func pagingResponse[T any](items []T, page, total int32) gin.H {
	return dataResponse(PagingResponse{
		Items: items,
		Meta: PagingMeta{
			Page:  page,
			Total: total,
		},
	})
}
