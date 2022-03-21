package usecase

import "github.com/SuperP2TL/Backend/presentation"

type UserRepo interface {
	GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error)
	AssignToSoToUser(userID int64, toSoIDs []int64) (id int64, err error)
}
