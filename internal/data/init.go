package data

import (
	"github.com/SuperP2TL/Backend/internal/data/delivery/rest"
	"github.com/SuperP2TL/Backend/internal/data/usecase"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, tosoRepo usecase.TOSORepo, dilRepo usecase.DILRepo, reportRepo usecase.ReportRepo, temuanRepo usecase.TemuanRepo) *Domain {
	uc := usecase.New(&usecase.Repositories{
		TOSORepo:   tosoRepo,
		DILRepo:    dilRepo,
		ReportRepo: reportRepo,
		TemuanRepo: temuanRepo,
	})

	httpHandler := rest.NewHTTP(router, uc, uc, uc, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}
