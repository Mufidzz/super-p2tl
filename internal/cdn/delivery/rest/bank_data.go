package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *HTTPHandler) HandleUploadFileBankData(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	category := ctx.PostForm("category")
	fileName := ctx.PostForm("name")

	if file == nil || category == "" || fileName == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	err = handler.usecases.StoreBankDataFile(file, category, fileName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada server, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleGetBankDataList(ctx *gin.Context) {

	res, err := handler.usecases.GetBankDataList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada server, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    res,
	})

}
