package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	DataIndukLayananUC
	DataTOSOUC
	ReportUC
	TemuanUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine, dataToSoUsecase DataTOSOUC, dilUsecase DataIndukLayananUC, reportUC ReportUC, temuanUC TemuanUC) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			DataTOSOUC:         dataToSoUsecase,
			DataIndukLayananUC: dilUsecase,
			ReportUC:           reportUC,
			TemuanUC:           temuanUC,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	data := router.Group("/data")
	{
		// DIL
		data.GET("/dil", handler.HandlerGetDIL)
		data.POST("/dil", handler.HandlerCreateDIL)
		data.PUT("/dil", handler.HandlerUpdateDIL)
		data.DELETE("/dil", handler.HandlerDeleteDIL)

		// Bank Data
		data.GET("/bank", handler.HandlerGetFromBank)

		// TO/SO
		data.GET("/to-so", handler.HandlerGetTOSO)
		data.POST("/to-so", handler.HandlerCreateSingleDataTOSO)
		data.PUT("/to-so", handler.HandlerUpdateSingleDataTOSO)
		data.DELETE("/to-so", handler.HandlerDeleteBulkDataTOSO)

	}

	dataReport := router.Group("/data/report")
	{
		dataReport.GET("/temuan", handler.HandleGetTemuanReport)
		dataReport.PUT("/temuan", handler.HandleUpdateSingleDataTemuanReport)
		dataReport.PUT("/temuan/jenis-temuan", handler.HandleUpdateBulkJenisTemuanOnTemuanReport)
		dataReport.GET("/temuan/count", handler.HandleGetTemuanReportCount)

		dataReport.GET("/penormalan", handler.HandleGetPenormalanReport)
	}
}
