package rest

import (
	"encoding/json"
	"fmt"
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func (handler *HTTPHandler) HandleStreamFileTemuan(ctx *gin.Context) {
	id, fileName := ctx.Param("id"), ctx.Param("filename")

	dir := path.Join(handler.usecases.GetTemuanBaseDirectory(), id)

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, "Not Found")
		return
	}

	for _, v := range files {
		if strings.Contains(v.Name(), fileName) {
			fixDir := path.Join(dir, v.Name())
			ctx.File(fixDir)
			fmt.Println(fixDir)

			return
		}
	}

	ctx.JSON(http.StatusNotFound, "Not Found")
}

func (handler *HTTPHandler) HandleUploadFileTemuan(ctx *gin.Context) {

	data := ctx.PostForm("data")

	fpBA, err := ctx.FormFile("photo_ba")
	if err != nil {
		log.Println("photo_ba", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	fpLokasi, err := ctx.FormFile("photo_lokasi")
	if err != nil {
		log.Println("photo_lokasi", err)
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada pembacaan file, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	if fpBA == nil || fpLokasi == nil || data == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Ada Kesalahan pada inputmu, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	var in presentation.CreateFindingReportsRequest

	err = json.Unmarshal([]byte(data), &in)
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

	id, err := handler.usecases.CreateSingleFindingReports(in)
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

	err = handler.usecases.StoreTemuanPhotoFiles(fpBA, fpLokasi, fmt.Sprint(id))
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
