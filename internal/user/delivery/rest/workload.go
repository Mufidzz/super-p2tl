package rest

import (
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

	tosoUserID, err := handler.usecases.AssignUserTOSOWorkload(requestData.UserID, requestData.ToSoID)
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

	//evtID, _ := handler.usecases.CreateSingleLogRoomEvent(presentation.CreateLogRoomEventRequest{
	//	Namespace: "/",
	//	Room:      fmt.Sprint(requestData.UserID),
	//	Args:      fmt.Sprintf("%d#%d",tosoUserID, evtID),
	//	Event:     "notify",
	//	Status:    1,
	//})

	handler.socketServer.BroadcastToRoom("/", fmt.Sprint(requestData.UserID), "notify", fmt.Sprintf("%d#%d", tosoUserID, 0))

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleAssignUserTemuanWorkload(ctx *gin.Context) {
	var requestData presentation.AssignUserTemuanWorkloadRequest

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
	if requestData.UserID == 0 || len(requestData.TemuanID) <= 0 {
		log.Println(err)

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	tosoUserID, err := handler.usecases.AssignUserTemuanWorkload(requestData.UserID, requestData.TemuanID)
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

	handler.socketServer.BroadcastToRoom("/", fmt.Sprint(requestData.UserID), "notify", fmt.Sprintf("%d#%d", tosoUserID, 0))

	ctx.JSON(http.StatusNoContent, "")
}

func (handler *HTTPHandler) HandleGetDataUserTOSOWorkload(ctx *gin.Context) {
	id := ctx.Param("user-id")

	intUserID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	data, err := handler.usecases.GetDataUserTOSOWorkload(intUserID)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada Server kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    data,
	})
}

func (handler *HTTPHandler) HandleGetDataUserTemuanWorkload(ctx *gin.Context) {
	id := ctx.Param("user-id")

	intUserID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada Inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	data, err := handler.usecases.GetDataUserTemuan(intUserID)
	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada Server kami, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "",
		Data:    data,
	})

}
