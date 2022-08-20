package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/pkg/urlp"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (handler *HTTPHandler) HandlerCreateFindingReports(ctx *gin.Context) {
	var in presentation.CreateFindingReportsRequest

	err := ctx.BindJSON(&in)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    in,
		})
		return
	}

	_, err = handler.usecases.CreateSingleFindingReport(in)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	err = handler.usecases.FinishTOSOCheck(in.TOSOId)
	if err != nil {

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandlerFinishTOSOCheck(ctx *gin.Context) {
	id := ctx.Param("user-toso-id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	err = handler.usecases.FinishTOSOCheck(intID)
	if err != nil {

		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleCreateSinglePenormalanReports(ctx *gin.Context) {
	var in presentation.CreatePenormalanReportsRequest

	err := ctx.BindJSON(&in)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    in,
		})
		return
	}

	_, err = handler.usecases.CreateSinglePenormalanReports(in)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleGetPerformanceKwhReport(ctx *gin.Context) {
	res, err := handler.usecases.GetPerformanceKwhReport()
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Server Kami, Coba Lagi",
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

func (handler *HTTPHandler) HandleGetListTemuanMangkrak(ctx *gin.Context) {
	var pagination presentation.Pagination

	paginationQP := ctx.Query("pagination")
	if paginationQP == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Parameter Pagination, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	err := urlp.DecodeEncodedString(paginationQP, &pagination)
	if paginationQP == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Parameter Pagination, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	res, err := handler.usecases.GetListTemuanMangkrak(pagination)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	dataCount, err := handler.usecases.GetListTemuanMangkrakCount()

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data: gin.H{
			"count":  dataCount,
			"temuan": res,
		},
	})
}
