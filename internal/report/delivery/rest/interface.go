package rest

import "github.com/SuperP2TL/Backend/presentation"

type ReportUC interface {
	CreateSingleFindingReport(input presentation.CreateFindingReportsRequest) (lastInsertedID int, err error)
	FinishTOSOCheck(userToSoID int) error
	CreateSinglePenormalanReports(input presentation.CreatePenormalanReportsRequest) (insertedID []int, err error)

	GetPerformanceKwhReport() (res []int64, err error)
	GetListTemuanMangkrak(pagination presentation.Pagination) (res []presentation.GetListTemuanMangkrakResponse, err error)
	GetListTemuanMangkrakCount() (res int, err error)
}
