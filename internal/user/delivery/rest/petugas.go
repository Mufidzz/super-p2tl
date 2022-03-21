package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/pkg/urlp"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *HTTPHandler) HandlerGetDataPetugas(ctx *gin.Context) {
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

	filterMap := &presentation.FilterParamUserData{
		Role: 3,
	}
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

	// Get Data With Pagination
	petugasData, err := handler.usecases.GetDataPetugas(filterMap, pagination)
	if err != nil {
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
		Data:    petugasData,
	})
	return

}
