package usecase

import "github.com/SuperP2TL/Backend/presentation"

type UserRepo interface {
	GetDataPetugas(filter *presentation.FilterParamUserData, pagination presentation.Pagination) (res []presentation.GetDataPetugasResponse, err error)
	AssignToSoToUser(userID int64, toSoIDs []int64) (id int64, err error)
	AssignTemuanToUser(userID int64, temuanIDs []int64) (id int64, err error)

	GetUserPassword(username string) (res *presentation.GetUserPasswordResponse, err error)
}

type RoomEventRepo interface {
	CreateLogRoomEvent(input []presentation.CreateLogRoomEventRequest) (lastInsertedID int, err error)
}

type TOSORepo interface {
	GetDataUserTOSO(userId int) (res []presentation.GetDataUserTOSOResponse, err error)
}

type TemuanRepo interface {
	GetDataUserTemuan(userId int) (res []presentation.GetDataUserTemuanResponse, err error)
}
