package rest

import "github.com/gin-gonic/gin"

type Usecases struct {
	BankDataUC
	ReportUC
}

type HTTPHandler struct {
	router   *gin.Engine
	usecases Usecases
}

func NewHTTP(router *gin.Engine, bankDataUC BankDataUC, reportUC ReportUC) *HTTPHandler {
	return &HTTPHandler{
		router: router,
		usecases: Usecases{
			BankDataUC: bankDataUC,
			ReportUC:   reportUC,
		},
	}
}

func (handler *HTTPHandler) SetRoutes() {
	router := handler.router
	fs := router.Group("/fs")

	bankData := fs.Group("/bank-data")
	{
		bankData.POST("/", handler.HandleUploadFileBankData)
		bankData.GET("/folder-data", handler.HandleGetBankDataList)
	}

	penormalan := fs.Group("/penormalan")
	{
		penormalan.POST("/", handler.HandleUploadFilePenormalan)
		//penormalan.GET("/stream/:id/:filename", handler.HandleUploadFilePenormalan)
		//penormalan.GET("/download/:id/:filename", handler.HandleUploadFilePenormalan)
	}

	temuan := fs.Group("/temuan")
	{
		temuan.POST("/", handler.HandleUploadFileTemuan)
		temuan.GET("/stream/:id/:filename", handler.HandleStreamFileTemuan)
		//temuan.GET("/download/:id/:filename", handler.HandleUploadFilePenormalan)
	}

}
