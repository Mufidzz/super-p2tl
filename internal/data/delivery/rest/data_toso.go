package rest

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/pkg/urlp"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (handler *HTTPHandler) HandlerGetTOSO(ctx *gin.Context) {
	var pagination presentation.Pagination

	// Data Filtration & Pagination Parsing
	filterEncodedQP := ctx.Query("filter")
	paginationQP := ctx.Query("pagination")
	//coreOnly := ctx.Query("core-only")

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
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan di Server Kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	fmt.Println(filterMap, pagination)

	toSoDatas, err := handler.usecases.GetDataTOSO(filterMap, pagination)
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
				"today": today,
				"all":   all,
			},
			"data": toSoDatas,
		},
	})
	return

}

func (handler *HTTPHandler) HandlerCreateSingleDataTOSO(ctx *gin.Context) {
	var in presentation.DataTOSO

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

	_, err = handler.usecases.CreateSingleDataTOSO(in)
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
func (handler *HTTPHandler) HandlerUpdateSingleDataTOSO(ctx *gin.Context) {

	var in presentation.DataTOSO

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

	_, err = handler.usecases.UpdateSingleDataTOSO(in)
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
func (handler *HTTPHandler) HandlerDeleteBulkDataTOSO(ctx *gin.Context) {
	ids := ctx.QueryArray("id")
	var intIds []int

	if len(ids) <= 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "ID Must more than 1",
			Type:    0,
			Data:    nil,
		})
		return
	}

	for _, v := range ids {
		_t, err := strconv.Atoi(v)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Success: false,
				Message: "Failed Parsing Tag ID, Please check all Tag ID is valid Number",
				Type:    0,
				Data:    ids,
			})
			return
		}

		intIds = append(intIds, _t)
	}

	err := handler.usecases.DeleteBulkDataTOSO(intIds)
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
