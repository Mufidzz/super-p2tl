package usecase

import "github.com/SuperP2TL/Backend/presentation"

func (uc *Usecase) AssignUserTOSOWorkload(userID int64, tosoIDs []int64) (id int64, err error) {
	return uc.repositories.AssignToSoToUser(userID, tosoIDs)
}

func (uc *Usecase) AssignUserTemuanWorkload(userID int64, temuanIds []int64) (id int64, err error) {
	return uc.repositories.AssignTemuanToUser(userID, temuanIds)
}

func (uc *Usecase) GetDataUserTOSOWorkload(userId int) (res []presentation.GetDataUserTOSOResponse, err error) {
	return uc.repositories.GetDataUserTOSO(userId)
}
