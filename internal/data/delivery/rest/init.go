package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	DataIndukLayananUC
	DataTOSOUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine, dataToSoUsecase DataTOSOUC) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			DataTOSOUC: dataToSoUsecase,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	data := router.Group("/data")
	{
		data.GET("/dil", handler.HandlerGetDIL)
		data.GET("/bank", handler.HandlerGetFromBank)
		data.GET("/to-so", handler.HandlerGetTOSO)
	}

}
