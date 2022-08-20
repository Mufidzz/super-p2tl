package usecase

import "github.com/SuperP2TL/Backend/presentation"

type ReportRepo interface {
	CreateFindingReports(input []presentation.CreateFindingReportsRequest) (lastInsertedID int, err error)
	CreatePenormalanReports(input []presentation.CreatePenormalanReportsRequest) (insertedID []int, err error)

	GetPerformanceKwhReport() (res []presentation.GetPerformanceKwhReportResponse, err error)
	GetListTemuanMangkrak(pagination presentation.Pagination) (res []presentation.GetListTemuanMangkrakResponse, err error)
	GetListTemuanMangkrakCount() (res int, err error)
}

type UserRepo interface {
	UpdateUserTOSO(newUserTOSO presentation.UpdateUserTOSORequest) error
}
