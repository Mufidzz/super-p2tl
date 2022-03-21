package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (handler *HTTPHandler) HandleAssignUserTOSOWorkload(ctx *gin.Context) {
	var requestData presentation.AssignUserTOSOWorkloadRequest

	err := ctx.BindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    requestData,
		})
		return
	}

	// Verify Input
	if requestData.UserID == 0 || len(requestData.ToSoID) <= 0 {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	_, err = handler.usecases.AssignUserTOSOWorkload(requestData.UserID, requestData.ToSoID)
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
