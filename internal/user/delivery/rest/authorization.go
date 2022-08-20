package rest

import (
	"github.com/SuperP2TL/Backend/pkg/response"
	"github.com/SuperP2TL/Backend/presentation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (handler *HTTPHandler) HandleLogin(ctx *gin.Context) {
	var loginData presentation.GetUserPasswordResponse

	err := ctx.BindJSON(&loginData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Success: false,
			Message: "Input tidak sesuai, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	userData, err := handler.usecases.Login(loginData.Username, loginData.Password)
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

	if userData == nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Success: false,
			Message: "Username & Password Anda Salah, Coba Lagi",
			Type:    0,
			Data:    nil,
		})
		return
	}

	userData.Password = "- SECRET -"
	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Login Berhasil",
		Data:    userData,
	})
	return
}
