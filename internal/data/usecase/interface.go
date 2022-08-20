package usecase

import "github.com/SuperP2TL/Backend/presentation"

type TOSORepo interface {
	GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error)
	GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error)
	//GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error)

	CreateBulkDataTOSO(in []presentation.DataTOSO) (insertedID []int, err error)
	UpdateBulkDataTOSO(in []presentation.DataTOSO) (updatedID []int, err error)
	DeleteBulkDataTOSO(dataIDs []int) error
}

type DILRepo interface {
	GetDataDIL(filter *presentation.FilterParamDIL, pagination presentation.Pagination) (res []presentation.GetDataDILResponse, err error)
	GetDataDILCount(filter *presentation.FilterParamDIL) (totalData int64, err error)

	CreateBulkDataDIL(in []presentation.DataDIL) (insertedID []int, err error)
	UpdateBulkDataDIL(in []presentation.DataDIL) (updatedID []int, err error)
	DeleteBulkDataDIL(dataIDs []int) error
}

type ReportRepo interface {
	GetTemuanReport(filter *presentation.FilterTemuanReport, pagination presentation.Pagination) (res []presentation.GetTemuanReportResponse, err error)
	GetTemuanReportCount(filter *presentation.FilterTemuanReport) (totalData int64, err error)

	GetPenormalanReport(filter *presentation.FilterPenormalanReport, pagination presentation.Pagination) (res []presentation.GetPenormalanReportResponse, err error)
	GetPenormalanReportCount(filter *presentation.FilterPenormalanReport) (totalData int64, err error)
}

type TemuanRepo interface {
	GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error)

	UpdateDataTemuanReport(in []presentation.DataTemuan) (updatedID []int, err error)
	UpdateBulkJenisTemuanOnTemuanReport(in presentation.UpdateBulkJenisTemuanOnTemuanReportRequest) (updatedID []int, err error)
}
