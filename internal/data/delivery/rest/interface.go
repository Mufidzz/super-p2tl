package rest

import "github.com/SuperP2TL/Backend/presentation"

type DataIndukLayananUC interface {
}

type BankDataUC interface {
}

type DataTOSOUC interface {
	GetDataTOSOCount(filter *presentation.FilterParamTOSOData) (todayCreation, totalData int64, err error)
	GetDataTOSO(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOResponse, err error)

	GetDataTOSOCore(filter *presentation.FilterParamTOSOData, pagination presentation.Pagination) (res []presentation.GetDataTOSOCoreResponse, err error)
}
