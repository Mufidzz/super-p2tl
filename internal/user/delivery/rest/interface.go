package rest

import "github.com/SuperP2TL/Backend/presentation"

type UserUC interface {
	GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error)
	AssignUserTOSOWorkload(userID int64, tosoIDs []int64) (id int64, err error)
}
