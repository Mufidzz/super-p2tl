package usecase

import (
	"github.com/SuperP2TL/Backend/presentation"
)

func (uc *Usecase) GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error) {
	filter.Role = presentation.USER_ROLE_PETUGAS
	return uc.repositories.GetDataPetugas(filter, pagination)
}
