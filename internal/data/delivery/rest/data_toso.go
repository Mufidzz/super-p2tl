package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/pkg/urlp"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *HTTPHandler) HandlerGetTOSO(ctx *gin.Context) {
	var pagination presentation.Pagination

	// Data Filtration & Pagination Parsing
	filterEncodedQP := ctx.Query("filter")
	paginationQP := ctx.Query("pagination")
	coreOnly := ctx.Query("core-only")

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

	var filterMap *presentation.FilterParamTOSOData
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
	today, all, err := handler.usecases.GetDataTOSOCount(filterMap)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	// Get Data With Pagination
	if coreOnly != "" {
		if filterMap == nil {
			filterMap = &presentation.FilterParamTOSOData{}
		}

		filterMap.NotAssignedOnly = true
		toSoDatas, err := handler.usecases.GetDataTOSOCore(filterMap, pagination)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Success: false,
				Message: "Ada Kesalahan di Server Kami, Coba Lagi",
				Type:    0,
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, response.SuccessResponse{
			Success: true,
			Message: "",
			Data: gin.H{
				"count": gin.H{
					"today": today,
					"all":   all,
				},
				"data": toSoDatas,
			},
		})

		return
	}

	toSoDatas, err := handler.usecases.GetDataTOSO(filterMap, pagination)
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
		Data: gin.H{
			"count": gin.H{
				"today": today,
				"all":   all,
			},
			"data": toSoDatas,
		},
	})
	return

}
