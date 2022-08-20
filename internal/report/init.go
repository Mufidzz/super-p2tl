package report

import (
	"fmt"
	"github.com/SuperP2TL/Backend/internal/report/delivery/rest"
	"github.com/SuperP2TL/Backend/internal/report/usecase"
	"github.com/gin-gonic/gin"
)

type Domain struct {
	Usecase *usecase.Usecase
}

func StartHTTP(router *gin.Engine, reportRepo usecase.ReportRepo, userRepo usecase.UserRepo) *Domain {
	uc := usecase.New(&usecase.Repositories{
		ReportRepo: reportRepo,
		UserRepo:   userRepo,
	})

	res, err := uc.GetPerformanceKwhReport()
	fmt.Println(res, err)

	httpHandler := rest.NewHTTP(router, uc)
	httpHandler.SetRoutes()

	return &Domain{
		Usecase: uc,
	}
}
