package cdn

import (
	"github.com/SuperP2TL/Backend/internal/cdn/delivery/rest"
	"github.com/SuperP2TL/Backend/internal/cdn/usecase"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, bankDataFilopsRepo usecase.BankDataFilopsRepo, penormalanFilopsRepo usecase.PenormalanFilopsRepo, reportRepo usecase.ReportRepo) *Domain {
	uc := usecase.New(&usecase.Repositories{
		BankDataFilopsRepo:   bankDataFilopsRepo,
		PenormalanFilopsRepo: penormalanFilopsRepo,
		ReportRepo:           reportRepo,
	})

	httpHandler := rest.NewHTTP(router, uc, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}
