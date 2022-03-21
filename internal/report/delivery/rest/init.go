package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine) *HTTPHandler {
	return &HTTPHandler{
		router: router,
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	uaa := router.Group("/report")
	{
		uaa.POST("/performance")
		uaa.POST("/temuan")
		uaa.POST("/penormalan")
	}
}
