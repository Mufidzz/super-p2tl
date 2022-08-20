package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	ReportUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine, reportUC ReportUC) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			ReportUC: reportUC,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	report := router.Group("/report")
	{
		report.GET("/finish/toso/:user-toso-id", handler.HandlerFinishTOSOCheck)
		report.GET("/temuan/kwh-report", handler.HandleGetPerformanceKwhReport)
		report.POST("/temuan", handler.HandlerCreateFindingReports)
		report.GET("/temuan/mangkrak", handler.HandleGetListTemuanMangkrak)
		report.POST("/penormalan", handler.HandleCreateSinglePenormalanReports)
	}
}
