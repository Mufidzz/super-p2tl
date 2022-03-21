package data

import (
	"github.com/SuperP2TL/Backend/internal/user/delivery/rest"
	"github.com/SuperP2TL/Backend/internal/user/usecase"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, userRepo usecase.UserRepo) *Domain {
	uc := usecase.New(&usecase.Repositories{
		UserRepo: userRepo,
	})

	httpHandler := rest.NewHTTP(router, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}
