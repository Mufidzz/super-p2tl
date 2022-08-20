package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetDataDIL(filter *presentation.FilterParamDIL, pagination presentation.Pagination) (res []presentation.GetDataDILResponse, err error) {
	return uc.repositories.GetDataDIL(filter, pagination)
}

func (uc *Usecase) GetDataDILCount(filter *presentation.FilterParamDIL) (totalData int64, err error) {
	return uc.repositories.GetDataDILCount(filter)
}

func (uc *Usecase) CreateSingleDataDIL(in presentation.DataDIL) (insertedID []int, err error) {
	return uc.repositories.CreateBulkDataDIL([]presentation.DataDIL{in})
}

func (uc *Usecase) UpdateSingleDataDIL(in presentation.DataDIL) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkDataDIL([]presentation.DataDIL{in})
}

func (uc *Usecase) UpdateBulkDataDIL(in []presentation.DataDIL) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkDataDIL(in)
}

func (uc *Usecase) DeleteBulkDataDIL(dataIDs []int) error {
	return uc.repositories.DeleteBulkDataDIL(dataIDs)
}
