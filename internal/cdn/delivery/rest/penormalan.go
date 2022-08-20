package rest

import (
	"encoding/json"
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (handler *HTTPHandler) HandleUploadFilePenormalan(ctx *gin.Context) {
	dataPenormalan := ctx.PostForm("data_penormalan")

	fpSebelum, err := ctx.FormFile("photo_sebelum")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	fpSesudah, err := ctx.FormFile("photo_sesudah")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	if fpSebelum == nil || fpSesudah == nil || dataPenormalan == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	var in presentation.CreatePenormalanReportsRequest

	err = json.Unmarshal([]byte(dataPenormalan), &in)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	ids, err := handler.usecases.CreateSinglePenormalanReports(in)
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

	err = handler.usecases.StorePenormalanPhotoFiles(fpSebelum, fpSesudah, fmt.Sprint(ids[0]))
	if err != nil {
		log.Println(err)
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
