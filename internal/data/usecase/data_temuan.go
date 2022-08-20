package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error) {
	return uc.repositories.GetDataUserTemuan(userId)
}

func (uc *Usecase) UpdateSingleDataTemuanReport(in presentation.DataTemuan) (updatedID []int, err error) {
	return uc.repositories.UpdateDataTemuanReport([]presentation.DataTemuan{in})
}

func (uc *Usecase) UpdateBulkJenisTemuanOnTemuanReport(in presentation.UpdateBulkJenisTemuanOnTemuanReportRequest) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkJenisTemuanOnTemuanReport(in)
}
