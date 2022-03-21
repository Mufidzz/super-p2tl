package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error) {
	return uc.repositories.GetDataTOSOCount(filter)
}

func (uc *Usecase) GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOResponse, err error) {
	return uc.repositories.GetDataTOSO(filter, pagination)
}

func (uc *Usecase) GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error) {
	return uc.repositories.GetDataTOSOCore(filter, pagination)
}
