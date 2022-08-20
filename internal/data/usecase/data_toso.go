package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error) {
	return uc.repositories.GetDataTOSOCount(filter)
}

func (uc *Usecase) GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
	return uc.repositories.GetDataTOSO(filter, pagination)
}

//func (uc *Usecase) GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
//	return uc.repositories.GetDataTOSO(filter, pagination)
//}

func (uc *Usecase) CreateSingleDataTOSO(in presentation.DataTOSO) (insertedID []int, err error) {
	return uc.repositories.CreateBulkDataTOSO([]presentation.DataTOSO{in})
}
func (uc *Usecase) UpdateSingleDataTOSO(in presentation.DataTOSO) (updatedID []int, err error) {
	return uc.repositories.UpdateBulkDataTOSO([]presentation.DataTOSO{in})
}
func (uc *Usecase) DeleteBulkDataTOSO(dataIDs []int) error {
	return uc.repositories.DeleteBulkDataTOSO(dataIDs)
}
