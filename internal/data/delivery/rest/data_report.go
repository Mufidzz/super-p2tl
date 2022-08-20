package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/pkg/urlp"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (handler *HTTPHandler) HandleGetTemuanReport(ctx *gin.Context) {

	var pagination presentation.Pagination

	// Data Filtration & Pagination Parsing
	filterEncodedQP := ctx.Query("filter")
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

	var filterMap *presentation.FilterTemuanReport
	if filterEncodedQP != "" {
		err = urlp.DecodeEncodedString(filterEncodedQP, &filterMap)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Ada Kesalahan di Parameter Filter, Coba Lagi",
				Type:    0,
				Data:    nil,
			})
			return
		}
	}

	// Get Data Count
	all, err := handler.usecases.GetTemuanReportCount(filterMap)
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

	toSoDatas, err := handler.usecases.GetTemuanReport(filterMap, pagination)
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

	// Return Data
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data: gin.H{
			"count": gin.H{
				"all": all,
			},
			"data": toSoDatas,
		},
	})
	return
}

func (handler *HTTPHandler) HandleGetTemuanReportCount(ctx *gin.Context) {

	// Data Filtration & Pagination Parsing
	filterEncodedQP := ctx.Query("filter")

	var filterMap *presentation.FilterTemuanReport
	if filterEncodedQP != "" {
		err := urlp.DecodeEncodedString(filterEncodedQP, &filterMap)
		if err != nil {
			log.Println(err)

			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Ada Kesalahan di Parameter Filter, Coba Lagi",
				Type:    0,
				Data:    nil,
			})
			return
		}
	}

	// Get Data Count
	all, err := handler.usecases.GetTemuanReportCount(filterMap)
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

	// Return Data
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    all,
	})
	return
}

func (handler *HTTPHandler) HandleGetPenormalanReport(ctx *gin.Context) {

	var pagination presentation.Pagination

	// Data Filtration & Pagination Parsing
	filterEncodedQP := ctx.Query("filter")
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

	var filterMap *presentation.FilterPenormalanReport
	if filterEncodedQP != "" {
		err = urlp.DecodeEncodedString(filterEncodedQP, &filterMap)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Ada Kesalahan di Parameter Filter, Coba Lagi",
				Type:    0,
				Data:    nil,
			})
			return
		}
	}

	// Get Data Count
	all, err := handler.usecases.GetPenormalanReportCount(filterMap)
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

	toSoDatas, err := handler.usecases.GetPenormalanReport(filterMap, pagination)
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

	// Return Data
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data: gin.H{
			"count": gin.H{
				"all": all,
			},
			"data": toSoDatas,
		},
	})
	return
}

func (handler *HTTPHandler) HandleUpdateBulkJenisTemuanOnTemuanReport(ctx *gin.Context) {
	var in presentation.UpdateBulkJenisTemuanOnTemuanReportRequest

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

	_, err = handler.usecases.UpdateBulkJenisTemuanOnTemuanReport(in)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan saat memproses request, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
	return
}

func (handler *HTTPHandler) HandleUpdateSingleDataTemuanReport(ctx *gin.Context) {
	var in presentation.DataTemuan

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

	_, err = handler.usecases.UpdateSingleDataTemuanReport(in)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan saat memproses request, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, "")
	return
}
