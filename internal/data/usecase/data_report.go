package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetTemuanReport(filter *presentation.FilterTemuanReport, pagination presentation.Pagination) (res []presentation.GetTemuanReportResponse, err error) {
	return uc.repositories.GetTemuanReport(filter, pagination)
}

func (uc *Usecase) GetTemuanReportCount(filter *presentation.FilterTemuanReport) (totalData int64, err error) {
	return uc.repositories.GetTemuanReportCount(filter)
}

func (uc *Usecase) GetPenormalanReportCount(filter *presentation.FilterPenormalanReport) (totalData int64, err error) {
	return uc.repositories.GetPenormalanReportCount(filter)
}
func (uc *Usecase) GetPenormalanReport(filter *presentation.FilterPenormalanReport, pagination presentation.Pagination) (res []presentation.GetPenormalanReportResponse, err error) {
	return uc.repositories.GetPenormalanReport(filter, pagination)
}
