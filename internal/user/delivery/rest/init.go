package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	UserUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine, userUsecase UserUC) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			UserUC: userUsecase,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	uaa := router.Group("/user")
	{
		uaa.POST("/login")
		uaa.GET("/petugas", handler.HandlerGetDataPetugas)
	}

	userWork := router.Group("/user-work")
	{
		userWork.POST("/to-so", handler.HandleAssignUserTOSOWorkload)
	}
}
