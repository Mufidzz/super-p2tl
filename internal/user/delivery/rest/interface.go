package rest

import "github.com/SuperP2TL/Backend/presentation"

type UserUC interface {
	GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error)
	AssignUserTOSOWorkload(userID int64, tosoIDs []int64) (id int64, err error)
	AssignUserTemuanWorkload(userID int64, temuanIds []int64) (id int64, err error)

	Login(username, password string) (user *presentation.GetUserPasswordResponse, err error)
}

type RoomEventUC interface {
	CreateSingleLogRoomEvent(input presentation.CreateLogRoomEventRequest) (lastInsertedID int, err error)
}

type ToSoUC interface {
	GetDataUserTOSOWorkload(userId int) (res []presentation.GetDataUserTOSOResponse, err error)
}

type TemuanUC interface {
	GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error)
}
